package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/nullsploit01/iosync/ent"
	"github.com/nullsploit01/iosync/internal/models"
	"github.com/nullsploit01/iosync/internal/mqtt_broker"
	"github.com/nullsploit01/iosync/internal/repository"
)

type NodeService struct {
	logger     *slog.Logger
	repo       repository.NodeRepository
	mqttBroker *mqtt_broker.MqttBroker
}

type CreateNodeRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type GenerateNodeAPIKeyRequest struct {
	Description string `json:"description" validate:"required"`
}

type AddNodeValueRequest struct {
	ApiKey string `json:"api_key" validate:"required"`
	Value  string `json:"value" validate:"required"`
}

func NewNodeService(db *ent.Client, mqttBroker *mqtt_broker.MqttBroker, logger *slog.Logger) NodeService {
	return NodeService{
		logger:     logger,
		repo:       repository.NewNodeRepository(db),
		mqttBroker: mqttBroker,
	}
}

func (n NodeService) InitNodeService(ctx context.Context) error {
	err := n.MonitorNodeOnlineStatus(ctx)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to monitor node online status: %v", err))
		return err
	}

	err = n.ListenForNodeValueUpdates(ctx)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to listen for node value updates: %v", err))
		return err
	}

	return nil
}

func (n NodeService) GetNodes(ctx context.Context) ([]models.Node, error) {
	nodes, err := n.repo.GetNodes(ctx)
	if err != nil {
		return nil, err
	}

	var nodeModels []models.Node
	for _, node := range nodes {
		nodeModels = append(nodeModels, models.Node{
			ID:           node.ID,
			Name:         node.Name,
			Description:  node.Description,
			Identifier:   node.Identifier,
			IsActive:     node.IsActive,
			IsOnline:     node.IsOnline,
			LastOnlineAt: node.LastOnlineAt,
			CreatedAt:    node.CreatedAt,
			UpdatedAt:    node.UpdatedAt,
		})
	}

	return nodeModels, nil
}

func (n NodeService) GetNode(ctx context.Context, id int) (models.NodeWithAPIKeys, error) {
	node, err := n.repo.GetNode(ctx, id)
	if err != nil {
		return models.NodeWithAPIKeys{}, err
	}

	nodeModel := models.NodeWithAPIKeys{
		ID:           node.ID,
		Name:         node.Name,
		Description:  node.Description,
		Identifier:   node.Identifier,
		IsActive:     node.IsActive,
		IsOnline:     node.IsOnline,
		LastOnlineAt: node.LastOnlineAt,
		CreatedAt:    node.CreatedAt,
		UpdatedAt:    node.UpdatedAt,
	}

	for _, apiKey := range node.Edges.APIKeys {
		nodeModel.APIKeys = append(nodeModel.APIKeys, models.NodeAPIKey{
			ID:          apiKey.ID,
			ApiKey:      apiKey.APIKey,
			Description: apiKey.Description,
			CreatedAt:   apiKey.CreatedAt,
			UpdatedAt:   apiKey.UpdatedAt,
		})
	}

	return nodeModel, nil

}

func (n NodeService) GetNodeValuesByAPIKey(ctx context.Context, apiKey string) ([]models.NodeApiKeyValue, error) {
	nodeApiKey, err := n.repo.GetNodeAPIByAPIKey(ctx, apiKey)
	if err != nil {
		return nil, err
	}

	nodeApiKeyValues, err := n.repo.GetNodeValuesByAPIKey(ctx, nodeApiKey)
	if err != nil {
		return nil, err
	}

	var nodeValues []models.NodeApiKeyValue
	for _, nodeValue := range nodeApiKeyValues {
		nodeValues = append(nodeValues, models.NodeApiKeyValue{
			ID:        nodeValue.ID,
			Value:     nodeValue.Value,
			CreatedAt: nodeValue.CreatedAt,
			UpdateAt:  nodeValue.UpdatedAt,
		})
	}

	return nodeValues, nil
}

func (n NodeService) CreateNode(ctx context.Context, request CreateNodeRequest) (models.Node, error) {
	node, err := n.repo.CreateNode(ctx, request.Name, request.Description)
	if err != nil {
		return models.Node{}, err
	}

	return models.Node{
		ID:           node.ID,
		Name:         node.Name,
		Description:  node.Description,
		Identifier:   node.Identifier,
		IsActive:     node.IsActive,
		LastOnlineAt: node.LastOnlineAt,
		CreatedAt:    node.CreatedAt,
		UpdatedAt:    node.UpdatedAt,
	}, nil

}

func (n NodeService) GenerateNodeAPIKey(ctx context.Context, nodeId int, request GenerateNodeAPIKeyRequest) (models.NodeAPIKey, error) {
	nodeApiKey, err := n.repo.GenerateNodeAPIKey(ctx, nodeId, request.Description)
	if err != nil {
		return models.NodeAPIKey{}, err
	}

	return models.NodeAPIKey{
		ID:          nodeApiKey.ID,
		ApiKey:      nodeApiKey.APIKey,
		Description: nodeApiKey.Description,
		CreatedAt:   nodeApiKey.CreatedAt,
		UpdatedAt:   nodeApiKey.UpdatedAt,
	}, nil
}

func (n NodeService) AddNodeValue(ctx context.Context, request AddNodeValueRequest) (models.NodeApiKeyValue, error) {
	node, err := n.repo.GetNodeAPIByAPIKey(ctx, request.ApiKey)
	if err != nil {
		return models.NodeApiKeyValue{}, err
	}

	apiKeyValue, err := n.repo.AddNodeValue(ctx, node, request.Value)
	if err != nil {
		return models.NodeApiKeyValue{}, err
	}

	return models.NodeApiKeyValue{
		ID:        apiKeyValue.ID,
		Value:     apiKeyValue.Value,
		CreatedAt: apiKeyValue.CreatedAt,
		UpdateAt:  apiKeyValue.UpdatedAt,
	}, nil
}

func (n NodeService) ListenForNodeValueUpdates(ctx context.Context) error {
	err := n.mqttBroker.Subscribe("nodes/+/value", 1, func(client mqtt.Client, msg mqtt.Message) {
		go func() {
			n.logger.Debug("Received node value update", "topic", msg.Topic(), "payload", string(msg.Payload()))
			nodeValue := string(msg.Payload())
			apiKey := parseTopicWildcard(msg.Topic())
			_, err := n.AddNodeValueToApiKey(ctx, apiKey, nodeValue)
			if err != nil {
				var notFound *ent.NotFoundError
				if errors.As(err, &notFound) {
					n.logger.Error("node not found for value update")
				} else {
					n.logger.Error("failed to add node value", "error", err)
				}
			}
		}()
	})

	if err != nil {
		return fmt.Errorf("failed to subscribe to nodes/+/value: %w", err)
	}

	return nil
}

func (n NodeService) AddNodeValueToApiKey(ctx context.Context, apiKey string, value string) (models.NodeApiKeyValue, error) {
	node, err := n.repo.GetNodeAPIByAPIKey(ctx, apiKey)
	if err != nil {
		return models.NodeApiKeyValue{}, err
	}

	apiKeyValue, err := n.repo.AddNodeValue(ctx, node, value)
	if err != nil {
		return models.NodeApiKeyValue{}, err
	}

	return models.NodeApiKeyValue{
		ID:        apiKeyValue.ID,
		Value:     apiKeyValue.Value,
		CreatedAt: apiKeyValue.CreatedAt,
		UpdateAt:  apiKeyValue.UpdatedAt,
	}, nil
}

func (n NodeService) MonitorNodeOnlineStatus(ctx context.Context) error {
	err := n.mqttBroker.Subscribe("nodes/+/status", 1, func(client mqtt.Client, msg mqtt.Message) {
		go func() {
			n.logger.Debug("Received node status update", "topic", msg.Topic(), "payload", string(msg.Payload()))
			nodeIdentifier := parseTopicWildcard(msg.Topic())
			payload := string(msg.Payload())
			err := n.UpdateNodeOnlineStatus(context.Background(), nodeIdentifier, payload == "online")
			if err != nil {
				var notFound *ent.NotFoundError
				if errors.As(err, &notFound) {
					n.logger.Error("node not found for status update")
				} else {
					n.logger.Error("failed to update node status", "error", err)
				}
			}
		}()
	})

	if err != nil {
		return fmt.Errorf("failed to subscribe to nodes/+/status: %w", err)
	}

	err = n.mqttBroker.Subscribe("nodes/+/lwt", 1, func(client mqtt.Client, msg mqtt.Message) {
		go func() {
			nodeIdentifier := parseTopicWildcard(msg.Topic())
			n.UpdateNodeOnlineStatus(context.Background(), nodeIdentifier, false)
		}()
	})

	if err != nil {
		return fmt.Errorf("failed to subscribe to nodes/+/lwt: %w", err)
	}

	go n.CheckNodeTimeouts(ctx)

	return nil
}

func (n NodeService) CheckNodeTimeouts(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	timeout := 60 * 5 * time.Second // 5 minutes

	for range ticker.C {
		nodes, err := n.repo.GetNodes(ctx)
		if err != nil {
			slog.Error(fmt.Sprintf("failed to get nodes: %v", err))
			continue
		}

		now := time.Now()

		for _, node := range nodes {
			lastSeen := node.LastOnlineAt
			if now.Sub(lastSeen) > timeout && node.IsOnline {
				n.logger.Warn("marking node as offline", slog.Group("node", "id", node.ID, "identifier", node.Identifier, "name", node.Name))
				n.UpdateNodeOnlineStatus(context.Background(), node.Identifier, false)
			}
		}
	}

}

func (n NodeService) UpdateNodeOnlineStatus(ctx context.Context, nodeIdentifier string, isOnline bool) error {
	return n.repo.UpdateNodeOnlineStatus(ctx, nodeIdentifier, isOnline)
}

func parseTopicWildcard(topic string) string {
	parts := strings.Split(topic, "/")
	return parts[1]
}

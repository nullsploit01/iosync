package service

import (
	"context"

	"github.com/nullsploit01/iosync/ent"
	"github.com/nullsploit01/iosync/internal/repository"
)

type NodeService struct {
	repo repository.NodeRepository
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

func NewNodeService(db *ent.Client) NodeService {
	return NodeService{
		repo: repository.NewNodeRepository(db),
	}
}

func (n NodeService) GetNodes(ctx context.Context) ([]*ent.Node, error) {
	return n.repo.GetNodes(ctx)
}

func (n NodeService) GetNode(ctx context.Context, id int) (*ent.Node, error) {
	return n.repo.GetNode(ctx, id)
}

func (n NodeService) CreateNode(ctx context.Context, request CreateNodeRequest) (*ent.Node, error) {
	return n.repo.CreateNode(ctx, request.Name, request.Description)
}

func (n NodeService) GenerateNodeAPIKey(ctx context.Context, nodeId int, request GenerateNodeAPIKeyRequest) (*ent.NodeApiKey, error) {
	return n.repo.GenerateNodeAPIKey(ctx, nodeId, request.Description)
}

func (n NodeService) AddNodeValue(ctx context.Context, request AddNodeValueRequest) (*ent.NodeValues, error) {
	node, err := n.repo.GetNodeByAPIKey(ctx, request.ApiKey)
	if err != nil {
		return nil, err
	}

	return n.repo.AddNodeValue(ctx, node, request.Value)
}

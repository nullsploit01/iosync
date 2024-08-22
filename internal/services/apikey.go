package services

import (
	"context"
	"errors"
	"iosync/ent"
	"iosync/internal/repositories"
	"iosync/pkg/utils"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type ApiKeyService struct {
	apiKeyRepository *repositories.ApiKeyRepository
	deviceRepository *repositories.DeviceRepository
	mqttService      *MQTTService
}

type CreateApiKeyRequest struct {
	Description string `json:"description"`
}

func NewApiKeyService(mqttClient mqtt.Client, dbClient *ent.Client) *ApiKeyService {
	apiKeyRepository := repositories.NewApiKeyRepository(dbClient)
	deviceRepository := repositories.NewDeviceRepository(dbClient)
	mqttService := NewMQTTService(mqttClient)

	return &ApiKeyService{
		apiKeyRepository: apiKeyRepository,
		deviceRepository: deviceRepository,
		mqttService:      mqttService,
	}
}

func (a *ApiKeyService) CreateApiKey(ctx context.Context, deviceId int, request *CreateApiKeyRequest) (*ent.ApiKey, error) {
	apiKey, err := utils.GenerateRandomString(32) // Fixed size of 32 digits for API KEYS
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	device, err := a.deviceRepository.Get(ctx, deviceId)
	if err != nil {
		return nil, err
	}

	expiryDate := time.Now().AddDate(10, 0, 0)

	payload := repositories.CreateApiKeyPayload{
		ApiKey:      apiKey,
		ExpiresAt:   expiryDate,
		Description: request.Description,
	}

	return a.apiKeyRepository.Create(ctx, &payload, device)
}

func (a *ApiKeyService) GetApiKey(ctx context.Context, deviceId int, key string) (*ent.ApiKey, error) {
	return a.apiKeyRepository.GetDeviceKey(ctx, deviceId, key)
}

func (a *ApiKeyService) ActivateApiKey(ctx context.Context, deviceId int, key string) error {
	apiKey, err := a.apiKeyRepository.GetDeviceKey(ctx, deviceId, key)
	if err != nil {
		return err
	}

	apiKey.IsActive = true
	apiKey.RevokedAt = &time.Time{}

	return a.apiKeyRepository.Update(ctx, apiKey)
}

func (a *ApiKeyService) RevokeApiKey(ctx context.Context, deviceId int, key string) error {
	apiKey, err := a.apiKeyRepository.GetDeviceKey(ctx, deviceId, key)
	if err != nil {
		return err
	}

	revokedAt := time.Now()
	apiKey.IsActive = false
	apiKey.RevokedAt = &revokedAt

	return a.apiKeyRepository.Update(ctx, apiKey)
}

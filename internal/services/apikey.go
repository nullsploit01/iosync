package services

import (
	"context"
	"errors"
	"iosync/ent"
	"iosync/internal/repositories"
	"iosync/pkg/utils"
	"time"
)

type ApiKeyService struct {
	apiKeyRepository *repositories.ApiKeyRepository
	deviceRepository *repositories.DeviceRepository
}

type CreateApiKeyRequest struct {
	Description string `json:"description"`
}

func NewApiKeyService(dbClient *ent.Client) *ApiKeyService {
	apiKeyRepository := repositories.NewApiKeyRepository(dbClient)
	deviceRepository := repositories.NewDeviceRepository(dbClient)

	return &ApiKeyService{
		apiKeyRepository: apiKeyRepository,
		deviceRepository: deviceRepository,
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

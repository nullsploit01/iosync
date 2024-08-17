package services

import (
	"context"
	"errors"
	"iosync/ent"
	"iosync/internal/repositories"
	"iosync/pkg/utils"
)

type ApiKeyService struct {
	apiKeyRepository *repositories.ApiKeyRepository
	deviceRepository *repositories.DeviceRepository
}

func NewApiKeyService(dbClient *ent.Client) *ApiKeyService {
	apiKeyRepository := repositories.NewApiKeyRepository(dbClient)
	deviceRepository := repositories.NewDeviceRepository(dbClient)

	return &ApiKeyService{
		apiKeyRepository: apiKeyRepository,
		deviceRepository: deviceRepository,
	}
}

func (a *ApiKeyService) CreateApiKey(ctx context.Context, deviceId int) (*ent.ApiKey, error) {
	apiKey, err := utils.GenerateRandomString(32) // Fixed size of 32 digits for API KEYS
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	device, err := a.deviceRepository.Get(ctx, deviceId)
	if err != nil {
		return nil, err
	}

	return a.apiKeyRepository.Create(ctx, apiKey, device)
}

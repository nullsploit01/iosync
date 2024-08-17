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
}

func NewApiKeyService(dbClient *ent.Client) *ApiKeyService {
	apiKeyRepository := repositories.NewApiKeyRepository(dbClient)

	return &ApiKeyService{
		apiKeyRepository: apiKeyRepository,
	}
}

func (a *ApiKeyService) CreateApiKey(ctx context.Context, deviceId int) (*ent.ApiKey, error) {
	apiKey, err := utils.GenerateRandomString(32) // Fixed size of 32 digits for API KEYS
	if err != nil {
		return nil, errors.New("something went wrong")
	}

	payload := repositories.CreateApiKeyPayload{
		ApiKey:   apiKey,
		DeviceId: deviceId,
	}

	return a.apiKeyRepository.Create(ctx, &payload)
}

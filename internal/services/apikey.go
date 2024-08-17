package services

import (
	"iosync/ent"
	"iosync/internal/repositories"
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

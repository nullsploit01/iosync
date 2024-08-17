package repositories

import (
	"context"
	"iosync/ent"
)

type ApiKeyRepository struct {
	dbClient *ent.Client
}

func NewApiKeyRepository(dbClient *ent.Client) *ApiKeyRepository {
	return &ApiKeyRepository{
		dbClient: dbClient,
	}
}

func (a *ApiKeyRepository) Create(ctx context.Context, apiKey string, device *ent.Device) (*ent.ApiKey, error) {
	return a.dbClient.ApiKey.Create().
		SetKey(apiKey).
		SetDevice(device).
		Save(ctx)
}

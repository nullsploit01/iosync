package repositories

import (
	"context"
	"iosync/ent"
)

type ApiKeyRepository struct {
	dbClient *ent.Client
}

type CreateApiKeyPayload struct {
	DeviceId int
	ApiKey   string
}

func NewApiKeyRepository(dbClient *ent.Client) *ApiKeyRepository {
	return &ApiKeyRepository{
		dbClient: dbClient,
	}
}

func (a *ApiKeyRepository) Create(ctx context.Context, payload *CreateApiKeyPayload) (*ent.ApiKey, error) {
	return a.dbClient.ApiKey.Create().
		SetDeviceID(payload.DeviceId).
		SetKey(payload.ApiKey).
		Save(ctx)
}

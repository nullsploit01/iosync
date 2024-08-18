package repositories

import (
	"context"
	"iosync/ent"
	"time"
)

type ApiKeyRepository struct {
	dbClient *ent.Client
}

type CreateApiKeyPayload struct {
	ApiKey      string
	Description string
	ExpiresAt   time.Time
}

func NewApiKeyRepository(dbClient *ent.Client) *ApiKeyRepository {
	return &ApiKeyRepository{
		dbClient: dbClient,
	}
}

func (a *ApiKeyRepository) Create(ctx context.Context, payload *CreateApiKeyPayload, device *ent.Device) (*ent.ApiKey, error) {
	return a.dbClient.ApiKey.Create().
		SetKey(payload.ApiKey).
		SetDescription(payload.Description).
		SetExpiresAt(payload.ExpiresAt).
		SetDevice(device).
		Save(ctx)
}

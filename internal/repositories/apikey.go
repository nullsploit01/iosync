package repositories

import (
	"context"
	"iosync/ent"
	"iosync/ent/apikey"
	"iosync/ent/device"
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

func (a *ApiKeyRepository) Update(ctx context.Context, apiKeyToUpdate *ent.ApiKey) error {
	updateQuery := a.dbClient.ApiKey.Update().Where(apikey.ID(apiKeyToUpdate.ID))

	if apiKeyToUpdate.Description != nil {
		updateQuery.SetDescription(*apiKeyToUpdate.Description)
	}
	if apiKeyToUpdate.ExpiresAt != nil {
		updateQuery.SetExpiresAt(*apiKeyToUpdate.ExpiresAt)
	}
	if apiKeyToUpdate.RevokedAt != nil {
		updateQuery.SetRevokedAt(*apiKeyToUpdate.RevokedAt)
	}
	if apiKeyToUpdate.LastUsed != (time.Time{}) {
		updateQuery.SetLastUsed(apiKeyToUpdate.LastUsed)
	}
	updateQuery.SetIsActive(apiKeyToUpdate.IsActive)

	_, err := updateQuery.Save(ctx)
	return err
}

func (a *ApiKeyRepository) GetDeviceKey(ctx context.Context, deviceId int, apiKey string) (*ent.ApiKey, error) {
	return a.dbClient.ApiKey.
		Query().
		Where(
			apikey.Key(apiKey),
			apikey.HasDeviceWith(device.ID(deviceId)),
		).
		First(ctx)
}

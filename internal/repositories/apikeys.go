package repositories

import "iosync/ent"

type ApiKeysRepository struct {
	dbClient *ent.Client
}

func NewApiKeysRepository(dbClient *ent.Client) *ApiKeysRepository {
	return &ApiKeysRepository{
		dbClient: dbClient,
	}
}

package repositories

import "iosync/ent"

type UserRepository struct {
	dbClient *ent.Client
}

func NewUserRepository(dbClient *ent.Client) *UserRepository {
	return &UserRepository{dbClient: dbClient}
}

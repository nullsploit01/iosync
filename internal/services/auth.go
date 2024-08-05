package services

import "iosync/ent"

type AuthService struct {
	dbClient *ent.Client
}

func NewAuthService(dbClient *ent.Client) *AuthService {
	return &AuthService{dbClient: dbClient}
}

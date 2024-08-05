package services

import (
	"iosync/ent"
)

type RegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type AuthService struct {
	dbClient *ent.Client
}

func NewAuthService(dbClient *ent.Client) *AuthService {
	return &AuthService{dbClient: dbClient}
}

func (a *AuthService) AuthenticateUser(request LoginRequest) error {
	return nil
}

func (a *AuthService) AddUser(request RegisterRequest) error {
	return nil
}

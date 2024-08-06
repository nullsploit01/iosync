package services

import (
	"context"
	"iosync/ent"
	"iosync/internal/repositories"
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
	userRepository *repositories.UserRepository
}

func NewAuthService(dbClient *ent.Client) *AuthService {
	userRepository := repositories.NewUserRepository(dbClient)

	return &AuthService{
		userRepository: userRepository,
	}
}

func (a *AuthService) AuthenticateUser(request LoginRequest) error {
	return nil
}

func (a *AuthService) AddUser(ctx context.Context, request RegisterRequest) (*ent.User, error) {
	addUserPayload := repositories.AddUserPayload{
		Username: request.Username,
		Password: request.Password,
		Name:     request.Name,
	}

	return a.userRepository.AddUser(ctx, &addUserPayload)
}

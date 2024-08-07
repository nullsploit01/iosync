package services

import (
	"context"
	"iosync/ent"
	"iosync/internal/repositories"

	"golang.org/x/crypto/bcrypt"
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
	hashedPassword, err := hashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	addUserPayload := repositories.AddUserPayload{
		Username: request.Username,
		Password: hashedPassword,
		Name:     request.Name,
	}

	return a.userRepository.AddUser(ctx, &addUserPayload)
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword compares a plaintext password with a hashed password and returns an error if they don't match.
func checkPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

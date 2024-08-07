package repositories

import (
	"context"
	"fmt"
	"iosync/ent"
	"iosync/ent/user"
	"time"
)

type UserRepository struct {
	dbClient *ent.Client
}

type AddUserPayload struct {
	Name     string `json:"name" validate:"required"`
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type GetUserPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func NewUserRepository(dbClient *ent.Client) *UserRepository {
	return &UserRepository{dbClient: dbClient}
}

func (u *UserRepository) AddUser(ctx context.Context, payload *AddUserPayload) (*ent.User, error) {
	fmt.Println("register", payload)

	user, err := u.dbClient.User.Create().
		SetName(payload.Name).
		SetUsername(payload.Username).
		SetPassword(payload.Password).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	return user, err
}

func (u *UserRepository) FindUserByUsername(ctx context.Context, username string) (*ent.User, error) {
	return u.dbClient.User.Query().
		Where(user.Username(username)).
		First(ctx)
}

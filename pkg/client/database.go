package client

import (
	"context"
	"fmt"
	"iosync/ent"
	"iosync/pkg/config"

	_ "github.com/lib/pq"
)

func NewDbClient(c context.Context, environment *config.Environment) (*ent.Client, error) {
	client, err := ent.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		environment.DB_HOST, environment.DB_PORT, environment.DB_USER, environment.DB_NAME, environment.DB_PASSWORD,
	))

	if err != nil {
		return nil, err
	}

	err = client.Schema.Create(c)
	return client, err
}

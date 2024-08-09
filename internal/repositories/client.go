package repositories

import (
	"context"
	"fmt"
	"iosync/ent"
	"iosync/pkg/config"
	"log"

	_ "github.com/lib/pq"
)

func GetDbClient(c context.Context) (*ent.Client, error) {
	env, err := config.LoadEnvironment()
	if err != nil {
		return nil, err
	}

	client, err := ent.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		env.DB_HOST, env.DB_PORT, env.DB_USER, env.DB_NAME, env.DB_PASSWORD,
	))

	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	err = client.Schema.Create(c)
	return client, err
}

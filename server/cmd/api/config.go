package main

import (
	"context"

	"github.com/nullsploit01/iosync/ent"
	"github.com/nullsploit01/iosync/internal/database"
	"github.com/nullsploit01/iosync/internal/env"
)

type appConfig struct {
	baseURL  string
	httpPort int
	dbClient *ent.Client
}

func GetAppConfig(ctx context.Context) (appConfig, error) {
	dbClient, err := database.NewDbClient(ctx)

	if err != nil {
		return appConfig{}, err
	}

	return appConfig{
		httpPort: env.GetInt("PORT", 8080),
		baseURL:  env.GetString("BASE_URL", "http://localhost:4444"),
		dbClient: dbClient,
	}, nil
}

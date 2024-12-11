package main

import (
	"github.com/nullsploit01/iosync/internal/env"
)

type appConfig struct {
	baseURL          string
	httpPort         int
	databaseUser     string
	databaseHost     string
	databasePort     string
	databaseName     string
	databasePassword string
}

func GetAppConfig() (appConfig, error) {
	return appConfig{
		httpPort:         env.GetInt("PORT", 8080),
		baseURL:          env.GetString("BASE_URL", "http://localhost:8080"),
		databaseUser:     env.GetString("DATABASE_USER", "postgres"),
		databasePassword: env.GetString("DATABASE_PASSWORD", "password"),
		databaseHost:     env.GetString("DATABASE_HOST", "localhost"),
		databasePort:     env.GetString("DATABASE_PORT", "5432"),
		databaseName:     env.GetString("DATABASE_NAME", "iosync"),
	}, nil
}

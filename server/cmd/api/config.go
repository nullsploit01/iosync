package main

import "github.com/nullsploit01/iosync/internal/env"

type appConfig struct {
	baseURL  string
	httpPort int
}

func GetAppConfig() appConfig {
	return appConfig{
		httpPort: env.GetInt("PORT", 8080),
		baseURL:  env.GetString("BASE_URL", "http://localhost:4444"),
	}
}

package config

import "github.com/nullsploit01/iosync/internal/env"

type config struct {
	Port int
}

func Get() config {
	return config{
		Port: env.GetInt("PORT", 8080),
	}
}

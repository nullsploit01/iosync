package config

import (
	"errors"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Environment struct {
	PORT        string
	DB_HOST     string
	DB_PORT     string
	DB_NAME     string
	DB_USER     string
	DB_PASSWORD string
}

func LoadEnvironment() (*Environment, error) {
	env := &Environment{
		PORT:        os.Getenv("PORT"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
	}

	// Optional: Validate that required environment variables are set
	if env.DB_HOST == "" || env.DB_PORT == "" || env.DB_NAME == "" || env.DB_USER == "" || env.DB_PASSWORD == "" {
		return nil, errors.New("missing required environment variables")
	}

	return env, nil
}

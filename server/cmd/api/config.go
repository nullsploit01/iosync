package main

import (
	"github.com/nullsploit01/iosync/internal/env"
)

type appConfig struct {
	baseURL            string
	httpPort           int
	databaseUser       string
	databaseHost       string
	databasePort       string
	databaseName       string
	databasePassword   string
	mqttBrokerHost     string
	mqttBrokerPort     int
	mqttBrokerUser     string
	mqttBrokerPassword string
	mqttBrokerClientID string
}

func GetAppConfig() (appConfig, error) {
	return appConfig{
		httpPort:           env.GetInt("PORT", 8080),
		baseURL:            env.GetString("BASE_URL", "http://localhost:8080"),
		databaseUser:       env.GetString("DATABASE_USER", "postgres"),
		databasePassword:   env.GetString("DATABASE_PASSWORD", "password"),
		databaseHost:       env.GetString("DATABASE_HOST", "localhost"),
		databasePort:       env.GetString("DATABASE_PORT", "5432"),
		databaseName:       env.GetString("DATABASE_NAME", "iosync"),
		mqttBrokerHost:     env.GetString("MQTT_BROKER_HOST", "localhost"),
		mqttBrokerPort:     env.GetInt("MQTT_BROKER_PORT", 1883),
		mqttBrokerUser:     env.GetString("MQTT_BROKER_USER", "mosquitto"),
		mqttBrokerPassword: env.GetString("MQTT_BROKER_PASSWORD", "killmenow"),
		mqttBrokerClientID: env.GetString("MQTT_BROKER_CLIENT_ID", "go_mqtt_client"),
	}, nil
}

package services

import mqtt "github.com/eclipse/paho.mqtt.golang"

type MQTTService struct {
	mqttClient *mqtt.Client
}

func NewMQTTService(mqttClient *mqtt.Client) *MQTTService {
	mqttService := &MQTTService{
		mqttClient: mqttClient,
	}

	return mqttService
}

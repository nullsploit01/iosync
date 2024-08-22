package services

import (
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTService struct {
	mqttClient mqtt.Client
}

func NewMQTTService(mqttClient mqtt.Client) *MQTTService {
	mqttService := &MQTTService{
		mqttClient: mqttClient,
	}

	return mqttService
}

func (m *MQTTService) PublishMessage(topic string, message any) error {
	if token := m.mqttClient.Publish(topic, 1, false, message); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	log.Printf("Published message: %s to topic: %s\n", message, topic)

	return nil
}

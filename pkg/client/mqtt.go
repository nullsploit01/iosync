package client

import (
	"fmt"
	"iosync/pkg/config"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func NewMQTTClient(environment *config.Environment) (mqtt.Client, error) {

	mqttClientOptions := mqtt.NewClientOptions()
	mqttClientOptions.AddBroker(environment.RABBITMQ_BROKER)
	mqttClientOptions.SetClientID("IOSYNC_MQTT_CLIENT")
	mqttClientOptions.SetUsername(environment.RABBITMQ_USERNAME)
	mqttClientOptions.SetPassword(environment.RABBITMQ_PASSWORD)

	mqttClientOptions.OnConnect = func(c mqtt.Client) {
		fmt.Println("Connected to RabbitMQ MQTT broker")
	}

	mqttClientOptions.OnConnectionLost = func(c mqtt.Client, err error) {
		fmt.Printf("RabbitMQ MQTT broker Connection lost: %v\n", err)
	}

	client := mqtt.NewClient(mqttClientOptions)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return client, nil
}

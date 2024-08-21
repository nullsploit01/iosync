package services

import (
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func InitMQTTClient() {
	mqttClientOptions := mqtt.NewClientOptions()
	mqttClientOptions.AddBroker("tcp://localhost:1883")
	mqttClientOptions.SetClientID("IOSYNC_MQTT_CLIENT")
	mqttClientOptions.SetUsername("admin")
	mqttClientOptions.SetPassword("instar")

	mqttClientOptions.OnConnect = func(c mqtt.Client) {
		fmt.Println("Connected to RabbitMQ MQTT broker")
	}

	mqttClientOptions.OnConnectionLost = func(c mqtt.Client, err error) {
		fmt.Printf("RabbitMQ MQTT broker Connection lost: %v\n", err)
	}

	client := mqtt.NewClient(mqttClientOptions)
	defer client.Disconnect(250)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("Error connecting: %v\n", token.Error())
	}
}

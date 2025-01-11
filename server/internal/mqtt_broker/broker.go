package mqtt_broker

import (
	"fmt"
	"sync"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttBroker struct {
	client       mqtt.Client
	handlers     map[string]mqtt.MessageHandler
	handlerMutex sync.RWMutex
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	fmt.Println("Connected to MQTT broker")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("MQTT Broker Connect lost: %v", err)
}

func NewMqttBroker(host string, port int, clientId, username, password string) (*MqttBroker, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", host, port))
	opts.SetClientID(clientId)
	opts.SetUsername(username)
	opts.SetPassword(password)

	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, token.Error()
	}

	return &MqttBroker{
		client:   client,
		handlers: make(map[string]mqtt.MessageHandler),
	}, nil
}

func (m *MqttBroker) Publish(topic string, qos byte, retained bool, payload interface{}) error {
	token := m.client.Publish(topic, qos, retained, payload)
	token.Wait()
	return token.Error()
}

func (m *MqttBroker) Subscribe(topic string, qos byte, handler mqtt.MessageHandler) error {
	m.handlerMutex.Lock()
	defer m.handlerMutex.Unlock()

	m.handlers[topic] = handler

	token := m.client.Subscribe(topic, qos, func(c mqtt.Client, msg mqtt.Message) {
		m.handlerMutex.RLock()
		defer m.handlerMutex.RUnlock()

		if handler, ok := m.handlers[msg.Topic()]; ok {
			handler(c, msg)
		} else {
			fmt.Println("No handler for topic: ", msg.Topic())
		}
	})
	token.Wait()
	return token.Error()
}

func (m *MqttBroker) Unsubscribe(topic string) error {
	m.handlerMutex.Lock()
	defer m.handlerMutex.Unlock()

	delete(m.handlers, topic)

	token := m.client.Unsubscribe(topic)
	token.Wait()
	return token.Error()
}

func (m *MqttBroker) Disconnect() {
	m.client.Disconnect(250)
}

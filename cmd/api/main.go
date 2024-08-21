package main

import (
	"context"
	"iosync/internal/server"
	"iosync/pkg/client"
	"iosync/pkg/config"
	"log"
)

func main() {
	ctx := context.Background()

	environment, err := config.LoadEnvironment()
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}

	dbClient, err := client.NewDbClient(ctx, environment)
	if err != nil {
		log.Fatal("error occured connecting database", err)
	}

	mqttClient, err := client.NewMQTTClient(environment)
	if err != nil {
		log.Fatal("error occured connecting mqtt broker", err)
	}

	server := server.InitServer(mqttClient, dbClient)

	log.Println("Started Server on address", server.Addr)

	err = server.ListenAndServeTLS("./certs/localhost.pem", "./certs/localhost-key.pem")
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}

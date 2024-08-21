package main

import (
	"context"
	"iosync/internal/repositories"
	"iosync/internal/server"
	"iosync/internal/services"
	"log"
)

func main() {
	c := context.Background()
	services.InitMQTTClient()

	dbClient, err := repositories.GetDbClient(c)

	if err != nil {
		log.Fatal("error occured connecting database, ", err)
	}

	server := server.InitServer(dbClient)

	log.Println("Started Server on address", server.Addr)

	err = server.ListenAndServeTLS("./certs/localhost.pem", "./certs/localhost-key.pem")
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}

package main

import (
	"context"
	"iosync/internal/data"
	"iosync/internal/server"
	"log"
)

func main() {
	c := context.Background()

	_, err := data.Init(c)

	if err != nil {
		log.Fatal("error occured connecting database, ", err)
	}

	server := server.NewServer()

	log.Println("Started Server on address", server.Addr)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}

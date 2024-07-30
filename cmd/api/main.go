package main

import (
	"iosync/internal/server"
	"log"
)

func main() {

	server := server.NewServer()

	log.Println("Started Server on address", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}

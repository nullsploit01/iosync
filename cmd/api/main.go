package main

import (
	"fmt"
	"iosync/internal/server"
	"log"
)

func main() {

	server := server.NewServer()

	log.Println("Started Server on address", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}

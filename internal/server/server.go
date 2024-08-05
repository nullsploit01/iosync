package server

import (
	"fmt"
	"iosync/ent"
	"iosync/internal/services"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port        int
	dbClient    *ent.Client
	authService *services.AuthService
}

func InitServer(client *ent.Client) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	authService := services.NewAuthService(client)

	NewServer := &Server{
		port:        port,
		dbClient:    client,
		authService: authService,
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}

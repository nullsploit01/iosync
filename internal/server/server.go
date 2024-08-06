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
	authService *services.AuthService
}

func InitServer(dbClient *ent.Client) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	authService := services.NewAuthService(dbClient)

	server := &Server{
		port:        port,
		authService: authService,
	}

	httpServer := &http.Server{
		Addr:         fmt.Sprintf(":%d", server.port),
		Handler:      server.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return httpServer
}

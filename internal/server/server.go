package server

import (
	"fmt"
	"iosync/ent"
	"iosync/internal/repositories"
	"iosync/internal/services"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port           int
	authService    *services.AuthService
	userRepository *repositories.UserRepository
}

func InitServer(client *ent.Client) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	authService := services.NewAuthService()

	userRepository := repositories.NewUserRepository(client)

	server := &Server{
		port:           port,
		authService:    authService,
		userRepository: userRepository,
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

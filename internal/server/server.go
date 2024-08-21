package server

import (
	"fmt"
	"iosync/ent"
	"iosync/internal/services"
	"net/http"
	"os"
	"strconv"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Server struct {
	port           int
	authService    *services.AuthService
	deviceService  *services.DeviceService
	apiKeyService  *services.ApiKeyService
	sessionService *services.SessionService
}

func InitServer(mqttClient *mqtt.Client, dbClient *ent.Client) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	authService := services.NewAuthService(dbClient)
	deviceService := services.NewDeviceService(dbClient)
	apiKeyService := services.NewApiKeyService(dbClient)
	sessionService := services.NewSessionService(dbClient)

	server := &Server{
		port:           port,
		authService:    authService,
		deviceService:  deviceService,
		apiKeyService:  apiKeyService,
		sessionService: sessionService,
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

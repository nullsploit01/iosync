package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
	"sync"

	_ "github.com/joho/godotenv/autoload"
	"github.com/lmittmann/tint"
	"github.com/nullsploit01/iosync/internal/database"
	"github.com/nullsploit01/iosync/internal/mqtt_broker"
	"github.com/nullsploit01/iosync/internal/service"
	"github.com/nullsploit01/iosync/internal/version"
)

type application struct {
	config   appConfig
	logger   *slog.Logger
	wg       sync.WaitGroup
	services appServices
}

type appServices struct {
	nodeService service.NodeService
}

func main() {
	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: slog.LevelDebug}))

	err := run(logger)
	if err != nil {
		trace := string(debug.Stack())
		logger.Error(err.Error(), "trace", trace)
		os.Exit(1)
	}
}

func run(logger *slog.Logger) error {
	config, err := GetAppConfig()
	if err != nil {
		return err
	}

	dbClient, err := database.NewDbClient(context.Background(), config.databaseUser, config.databasePassword, config.databaseHost, config.databasePort, config.databaseName)
	if err != nil {
		return err
	}

	defer dbClient.Close()

	mqttBroker, err := mqtt_broker.NewMqttBroker(config.mqttBrokerHost, config.mqttBrokerPort, config.mqttBrokerClientID, config.mqttBrokerUser, config.mqttBrokerPassword)
	if err != nil {
		return err
	}

	defer mqttBroker.Disconnect()

	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	nodeService := service.NewNodeService(dbClient)

	appServices := appServices{
		nodeService: nodeService,
	}

	app := &application{
		config:   config,
		logger:   logger,
		services: appServices,
	}

	app.logger.Info("connected to database", slog.Group("database", "host", config.databaseHost, "port", config.databasePort, "name", config.databaseName))
	return app.serveHTTP()
}

package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"runtime/debug"
	"sync"

	"github.com/lmittmann/tint"
	"github.com/nullsploit01/iosync/internal/version"
)

type application struct {
	config appConfig
	logger *slog.Logger
	wg     sync.WaitGroup
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
	config, err := GetAppConfig(context.Background())
	if err != nil {
		return err
	}

	showVersion := flag.Bool("version", false, "display version and exit")

	flag.Parse()

	if *showVersion {
		fmt.Printf("version: %s\n", version.Get())
		return nil
	}

	app := &application{
		config: config,
		logger: logger,
	}

	return app.serveHTTP()
}

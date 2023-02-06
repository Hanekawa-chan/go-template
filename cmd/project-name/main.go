package main

import (
	"context"
	"fmt"
	"github.com/rs/zerolog"
	"log"
	"os"
	"os/signal"
	"project-name/internal/app"
	"project-name/internal/app/config"
	"project-name/internal/httpserver"
	"project-name/internal/version"
	"syscall"
	"time"
)

func main() {
	//Print version and commit sha
	log.Println("Loading Mailing - v", version.Version, "| Commit:", version.Commit)

	// Parse all configs form env
	cfg, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)

	level, err := zerolog.ParseLevel(cfg.Logger.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	logger := zerolog.New(os.Stdout).Level(level)

	zl := &logger

	service := app.NewService(zl, cfg)
	httpServerAdapter := httpserver.NewAdapter(zl, cfg.HTTPServer, service)

	// Channels for errors and os signals
	stop := make(chan error, 1)
	osSig := make(chan os.Signal, 1)
	signal.Notify(osSig, syscall.SIGINT, syscall.SIGTERM)

	// Receive errors form start bot func into error channel
	go func(stop chan<- error) {
		stop <- httpServerAdapter.ListenAndServe()
	}(stop)

	// Blocking select
	select {
	case sig := <-osSig:
		zl.Info().Msgf("Received os syscall signal %v", sig)
	case err := <-stop:
		zl.Error().Err(err).Msg("Received Error signal")
	}

	// Shutdown code
	zl.Info().Msg("Shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	if err := httpServerAdapter.Shutdown(ctx); err != nil {
		zl.Error().Err(err).Msg("Error shutting down the HTTP server!")
	}

	time.Sleep(time.Second * 2)

	zl.Info().Msg("Shutdown - success")
}

package main

import (
	"context"
	"fmt"
	"github.com/Ira11111/ProductService/internal/app"
	"github.com/Ira11111/ProductService/internal/config"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)

	logger := initLogger(cfg.Env)
	if logger == nil {
		panic("logger is nil")
	}

	logger.Info("init application")
	app := app.NewApp(logger, &cfg.DB, cfg.Server.Port, cfg.Server.ReadTimeout, cfg.Server.WriteTimeout)

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	logger.Info("start application")
	go func() {
		logger.Info("start server on port", slog.String("port", cfg.Server.Port))
		err := app.HttpServer.Start()
		if err != nil {
			panic(err)
		}
	}()

	<-stopChan
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	app.HttpServer.Stop(ctx)
	logger.Info("Application stopped")
}

const (
	envlocal       = "local"
	envdevelopment = "develop"
	envproduction  = "prod"
)

func initLogger(envType string) *slog.Logger {
	var logger *slog.Logger

	switch envType {
	case envlocal:
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envdevelopment:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case envproduction:
		logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return logger
}

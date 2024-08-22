package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"service-template/config"
	v1 "service-template/internal/controller/v1"
	"service-template/internal/repo"
	"service-template/internal/service"

	"service-template/pkg/httpserver"
	"service-template/pkg/postgres"
)

func Run(pathConfig string) {
	// Configuration
	cfg, err := config.NewConfig(pathConfig)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Logger
	logrusLevel, err := log.ParseLevel(cfg.Log.Level)
	if err != nil {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(logrusLevel)
	}
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.SetOutput(os.Stdout)

	// Database connection
	log.Info("Initializing postgres...")
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.MaxPoolSize))
	if err != nil {
		log.Fatal(fmt.Errorf("initializing postgres error: %w", err))
	}
	defer pg.Close()

	// Repositories
	log.Info("Initializing repositories...")
	repositories := repo.NewRepositories(pg)

	// Services
	log.Info("Initializing services...")
	services := service.NewServices(service.ServiceDependencies{
		Repos: repositories,
	})

	// Router handler
	log.Info("Initializing handlers and routes...")
	handler := echo.New()
	// setup handler validator as lib validator
	// handler.Validator = validator.NewCustomValidator()
	v1.NewRouter(handler, services)

	// HTTP server
	log.Info("Starting http server...")
	log.Debugf("Server port: %s", cfg.HTTP.Port)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	log.Info("Configuring graceful shutdown...")
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Graceful shutdown
	log.Info("Shutting down...")
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}

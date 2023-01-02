package app

import (
	"context"
	"deep-ocean-api/internal/config"
	"deep-ocean-api/internal/domain/service"
	"deep-ocean-api/internal/repository"
	"deep-ocean-api/internal/server"
	h "deep-ocean-api/internal/transport/http"
	"deep-ocean-api/pkg/database/postgres"
	"deep-ocean-api/pkg/hash"
	"deep-ocean-api/pkg/logger"

	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	cfg, err := config.Init()
	if err != nil {
		logger.Error("failed to parse environment variables", err)

		return
	}

	db, err := postgres.NewClient(cfg)
	if err != nil {
		logger.Error("failed to create a database connection", err)

		return
	}

	repositories := repository.NewRepositories(db)

	hasher := hash.NewSHA1Hasher(cfg.HashSalt)

	serviceDeps := &service.ServicesDeps{
		Repositories: repositories,
		Hasher:       hasher,
	}
	services := service.NewServices(serviceDeps)

	handler := h.NewHandler(services)

	srv := server.NewServer(cfg, handler.Init())

	go func() {
		if err := srv.Run(); err != nil {
			logger.Error("failed to run server", err)
			os.Exit(1)
			return
		}
	}()

	logger.Info("Server is started")

	// Graceful Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	_, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()
}

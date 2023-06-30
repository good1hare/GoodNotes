// Package app configures and runs application.
package app

import (
	"MateMind/configs"
	v1 "MateMind/internal/controller/http"
	"MateMind/internal/usecase"
	"MateMind/internal/usecase/repo"
	"MateMind/pkg/httpserver"
	"MateMind/pkg/logger"
	"MateMind/pkg/postgres"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

// Run creates objects via constructors.
func Run(cfg *configs.Config) {
	log := logger.New(cfg.Logger.Level)

	////Repository
	pg, err := postgres.New(cfg.Postgres.Url, postgres.MaxPoolSize(cfg.Postgres.PoolMax))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}

	// Use case
	userUseCase := usecase.New(
		repo.New(pg),
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, log, userUseCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.Http.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		log.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		log.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}

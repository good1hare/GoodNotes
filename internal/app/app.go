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
	"github.com/gin-gonic/gin"

	handler "MateMind/internal/controller/telegramHandler"
	telegram "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"os/signal"
	"syscall"
)

// Run creates objects via constructors.
func Run(cfg *configs.Config) {
	log := logger.New(cfg.Logger.Level)

	//Repository
	pg, err := postgres.New(cfg.Postgres.Url, postgres.MaxPoolSize(cfg.Postgres.PoolMax))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.NewUserUseCase: %w", err))
	}

	//Use case
	userUseCase := usecase.NewUserUseCase(
		repo.New(pg),
	)

	//Telegram API
	bot, err := telegram.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = cfg.Telegram.Debug
	log.Info("Authorized on account %s", bot.Self.UserName)
	u := telegram.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		th := handler.NewTelegramHandler(bot, update, log, userUseCase)
		th.Handle()
	}
	//HTTP Server
	ginHandler := gin.New()
	v1.NewRouter(ginHandler)
	httpServer := httpserver.New(ginHandler, httpserver.Port(cfg.Http.Port))

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

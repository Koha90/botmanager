// Package app composes and write all application dependecies.
//
// It initializes ifrastructure, services, transport layer
// and start the application serer.
package app

import (
	"log/slog"
	"net/http"

	"botmanager/internal/config"
	transporthttp "botmanager/internal/transport/http"
	"botmanager/internal/transport/http/handler"
	"botmanager/pkg/logger"
)

type App struct {
	server *http.Server
}

func NewApp(cfg *config.Config) *App {
	logger, err := logger.Setup(cfg.Env)
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	slog.Info("application started")

	orderService := buildOrderService(logger.Logger)

	handler := handler.NewOrderHandler(orderService)
	router := transporthttp.NewRouter(handler)

	server := &http.Server{
		Addr:    ":" + cfg.HTTP.Port,
		Handler: router,
	}

	return &App{server: server}
}

func (a *App) Run() error {
	return a.server.ListenAndServe()
}

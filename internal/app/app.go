// Package app composes and write all application dependecies.
//
// It initializes ifrastructure, services, transport layer
// and start the application serer.
package app

import (
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
	logger := logger.SetupLogger(cfg.Env)

	orderService := buildOrderService(logger)

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

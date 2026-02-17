// Package app provide general ...
package app

import (
	"database/sql"

	"botmanager/internal/service"
	"botmanager/internal/storage/postgres"
)

type App struct {
	OrderService   *service.OrderService
	ProductService *service.ProductService
}

func NewApp(db *sql.DB) *App {
	productRepo := postgres.NewProductRepo()
	orderRepo := postgres.NewOrderRepo()

	orderService := service.NewOrderService(productRepo, orderRepo)

	return &App{
		OrderService: orderService,
	}
}

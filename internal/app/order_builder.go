package app

import (
	"log/slog"
	"sync"

	"botmanager/internal/infrastructure/eventbus"
	"botmanager/internal/service"
	"botmanager/internal/storage/memory"
)

func buildOrderService(logger *slog.Logger) *service.OrderService {
	mu := &sync.Mutex{}
	// repositories
	orderRepo := memory.NewOrderRepository()
	productRepo := memory.NewProductRepository()

	// event bus
	bus := eventbus.New(logger)

	// transaction manager
	txManager := memory.NewTxManager(mu)

	return service.NewOrderService(productRepo, orderRepo, orderRepo, txManager, bus, logger)
}

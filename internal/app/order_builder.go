package app

import (
	"context"
	"log/slog"
	"sync"

	"botmanager/internal/domain"
	"botmanager/internal/infrastructure/eventbus"
	"botmanager/internal/service"
	"botmanager/internal/storage/memory"
)

func buildOrderService(logger *slog.Logger) *service.OrderService {
	mu := &sync.Mutex{}
	// repositories
	orderRepo := memory.NewOrderRepository()
	productRepo := memory.NewProductRepository(mu)

	// Seed data
	{
		p, err := domain.NewProduct("Test product", 1, "Desc Test", "")
		if err != nil {
			panic(err)
		}
		_ = productRepo.Create(context.Background(), p)
	}

	// event bus
	bus := eventbus.New(logger)

	// transaction manager
	txManager := memory.NewTxManager(mu)

	return service.NewOrderService(
		productRepo,
		orderRepo,
		orderRepo,
		txManager,
		bus,
		logger,
	)
}

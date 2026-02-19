package service

import (
	"context"

	"botmanager/internal/domain"
)

type OrderRepository interface {
	Create(ctx context.Context, order *domain.Order) error
	ByID(ctx context.Context, id int) (*domain.Order, error)
	ListByCustomer(ctx context.Context, customerID int) ([]domain.Order, error)
	Update(ctx context.Context, order *domain.Order) error
}

type orderCreator interface {
	Create(ctx context.Context, order *domain.Order) error
}

type orderUpdater interface {
	ByID(ctx context.Context, id int) (*domain.Order, error)
	Update(ctx context.Context, order *domain.Order) error
}

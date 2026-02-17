package postgres

import (
	"context"

	"botmanager/internal/domain"
	"botmanager/internal/service"
)

type OrderRepo struct{}

func NewOrderRepo() *OrderRepo {
	return &OrderRepo{}
}

// ByID implements [service.OrderRepository].
func (o *OrderRepo) ByID(ctx context.Context, id int) (*domain.Order, error) {
	panic("unimplemented")
}

// Create implements [service.OrderRepository].
func (o *OrderRepo) Create(ctx context.Context, order *domain.Order) error {
	panic("unimplemented")
}

// ListByCustomer implements [service.OrderRepository].
func (o *OrderRepo) ListByCustomer(ctx context.Context, customerID int) ([]domain.Order, error) {
	panic("unimplemented")
}

// Update implements [service.OrderRepository].
func (o *OrderRepo) Update(ctx context.Context, order *domain.Order) error {
	panic("unimplemented")
}

var _ service.OrderRepository = (*OrderRepo)(nil)

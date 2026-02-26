package memory

import (
	"context"

	"botmanager/internal/domain"
)

type OrderRepository struct {
	orders map[int]*domain.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[int]*domain.Order),
	}
}

func (r *OrderRepository) Create(ctx context.Context, order *domain.Order) error {
	r.orders[order.ID()] = order

	return nil
}

func (r *OrderRepository) ByID(ctx context.Context, id int) (*domain.Order, error) {
	order, ok := r.orders[id]
	if !ok {
		return nil, domain.ErrOrderNotFound
	}
	return order, nil
}

func (r *OrderRepository) Update(ctx context.Context, order *domain.Order) error {
	if _, ok := r.orders[order.ID()]; !ok {
		return domain.ErrOrderNotFound
	}

	r.orders[order.ID()] = order
	return nil
}

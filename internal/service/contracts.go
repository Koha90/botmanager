package service

import (
	"context"
	"time"

	"botmanager/internal/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, p *domain.Product) error
	ByID(ctx context.Context, id int) (*domain.Product, error)
}

type ProductReader interface {
	ByID(ctx context.Context, id int) (*domain.Product, error)
}

// OrderRepository defines persistence contain for Order aggregate.
type OrderRepository interface {
	Save(ctx context.Context, order *domain.Order) error
	ByID(ctx context.Context, id int) (*domain.Order, error)
	Cancel(now time.Time) error
}

// UserRepository defines persistence operations
// required by UserService.
type UserRepository interface {
	Save(ctx context.Context, u *domain.User) error
	ByID(ctx context.Context, id int) (*domain.User, error)
}

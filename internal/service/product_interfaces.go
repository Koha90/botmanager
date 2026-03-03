package service

import (
	"context"

	"botmanager/internal/domain"
)

type ProductRepository interface {
	Save(ctx context.Context, p *domain.Product) error
	ByID(ctx context.Context, id int) (*domain.Product, error)
}

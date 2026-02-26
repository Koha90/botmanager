package service

import (
	"context"

	"botmanager/internal/domain"
)

type ProductRepository interface {
	// List(ctx context.Context) ([]domain.Product, error)
	// ListByCity(ctx context.Context, cityID int) ([]domain.Product, error)
	// ListByDistrict(ctx context.Context, districtID int) ([]domain.Product, error)
	// ListByCategory(ctx context.Context, categoryID int) ([]domain.Product, error)
	ByID(ctx context.Context, id int) (*domain.Product, error)
	Create(ctx context.Context, product *domain.Product) error
	Update(ctx context.Context, product *domain.Product) error
	Delete(ctx context.Context, id int) error
}

type productCreator interface {
	Create(ctx context.Context, product *domain.Product)
}

type productReader interface {
	ByID(ctx context.Context, id int) (*domain.Product, error)
	// VariantByID(id int) (*domain.ProductVariant, error)
}

type productUpdater interface {
	ByID(ctx context.Context, id int) (*domain.Product, error)
	Update(ctx context.Context, product *domain.Product) error
}

package service

import (
	"context"

	"botmanager/internal/domain"
)

type ProductRepository interface {
	List(ctx context.Context) ([]domain.Product, error)
	ListByCity(ctx context.Context, cityID int) ([]domain.Product, error)
	ListByDistrict(ctx context.Context, districtID int) ([]domain.Product, error)
	ListByCategory(ctx context.Context, categoryID int) ([]domain.Product, error)
	ByID(ctx context.Context, id int) (*domain.Product, error)
	Create(ctx context.Context, product *domain.Product) error
	Update(ctx context.Context, product *domain.Product) error
	Delete(ctx context.Context, id int) error
}

type ProductService struct {
	products ProductRepository
}

func NewProductService(products ProductRepository) *ProductService {
	return &ProductService{
		products: products,
	}
}

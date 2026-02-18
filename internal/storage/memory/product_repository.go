package memory

import (
	"context"

	"botmanager/internal/domain"
)

type ProductRepository struct {
	products map[int]*domain.Product
}

func NewProductRepository() *ProductRepository {
	return &ProductRepository{
		products: map[int]*domain.Product{
			1: {ID: 1, Price: 1000},
			2: {ID: 2, Price: 2000},
		},
	}
}

func (r *ProductRepository) ByID(ctx context.Context, id int) (*domain.Product, error) {
	product, ok := r.products[id]
	if !ok {
		return nil, domain.ErrProductNotFound
	}
	return product, nil
}

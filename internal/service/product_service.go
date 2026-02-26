package service

import (
	"context"

	"botmanager/internal/domain"
)

type ProductService struct {
	repo ProductRepository
}

// NewProductService creates new ProductSevice.
func NewProductService(repo ProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

// Create validate input, create domain entity
// and persist it using repository.
//
// Returns created product.
func (s *ProductService) Create(
	id int,
	ctx context.Context,
	name string,
	categoryID int,
	description string,
	imagePath string,
) (*domain.Product, error) {
	// 1. Create domain entity (validation happens inside)
	product, err := domain.NewProduct(
		id,
		name,
		categoryID,
		description,
		imagePath,
	)
	if err != nil {
		return nil, err
	}

	// 2. Persist
	if err := s.repo.Create(ctx, product); err != nil {
		return nil, err
	}

	return product, nil
}

// ByID returns product by identifier.
func (s *ProductService) ByID(
	ctx context.Context,
	id int,
) (*domain.Product, error) {
	return s.repo.ByID(ctx, id)
}

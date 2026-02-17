// Package postgres provide methods for monipulate data postgresql.
package postgres

import (
	"context"

	"botmanager/internal/domain"
	"botmanager/internal/service"
)

type ProductRepo struct{}

var _ service.ProductRepository = (*ProductRepo)(nil)

func NewProductRepo() *ProductRepo {
	return &ProductRepo{}
}

// ByID implements [service.ProductRepository].
func (p *ProductRepo) ByID(ctx context.Context, id int) (*domain.Product, error) {
	panic("unimplemented")
}

// Create implements [service.ProductRepository].
func (p *ProductRepo) Create(ctx context.Context, product *domain.Product) error {
	panic("unimplemented")
}

// Delete implements [service.ProductRepository].
func (p *ProductRepo) Delete(ctx context.Context, id int) error {
	panic("unimplemented")
}

// List implements [service.ProductRepository].
func (p *ProductRepo) List(ctx context.Context) ([]domain.Product, error) {
	panic("unimplemented")
}

// ListByCategory implements [service.ProductRepository].
func (p *ProductRepo) ListByCategory(
	ctx context.Context,
	categoryID int,
) ([]domain.Product, error) {
	panic("unimplemented")
}

// ListByCity implements [service.ProductRepository].
func (p *ProductRepo) ListByCity(ctx context.Context, cityID int) ([]domain.Product, error) {
	panic("unimplemented")
}

// ListByDistrict implements [service.ProductRepository].
func (p *ProductRepo) ListByDistrict(
	ctx context.Context,
	districtID int,
) ([]domain.Product, error) {
	panic("unimplemented")
}

// Update implements [service.ProductRepository].
func (p *ProductRepo) Update(ctx context.Context, product *domain.Product) error {
	panic("unimplemented")
}

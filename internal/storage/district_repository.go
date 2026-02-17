package storage

import (
	"context"

	"botmanager/internal/domain"
)

type DistrictRepository interface {
	List(ctx context.Context) ([]domain.District, error)
	ListByCity(ctx context.Context, cityID int) ([]domain.District, error)
	ListByCategory(ctx context.Context, categoryID int) ([]domain.District, error)
	ListByProduct(ctx context.Context, productID int) ([]domain.District, error)
	ByID(ctx context.Context, id int) (*domain.District, error)
	Create(ctx context.Context, district *domain.District) error
	Update(ctx context.Context, district *domain.District) error
	DeleteByID(ctx context.Context, id int) error
}

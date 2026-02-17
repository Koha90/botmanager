package storage

import (
	"context"

	"botmanager/internal/domain"
)

type CategoryRepository interface {
	List(ctx context.Context) ([]domain.Category, error)
	ListByCity(ctx context.Context, cityID int) ([]domain.Category, error)
	ListByDistrict(ctx context.Context, districtID int) ([]domain.Category, error)
	ByID(ctx context.Context, id int) (*domain.Category, error)
	Create(ctx context.Context, category *domain.Category) error
	Update(ctx context.Context, category *domain.Category) error
	DeleteByID(ctx context.Context, id int) error
}

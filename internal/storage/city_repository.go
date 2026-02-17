// Package storage provide methods for manupulate of data.
package storage

import (
	"context"

	"botmanager/internal/domain"
)

type CityRepository interface {
	List(ctx context.Context) ([]domain.City, error)
	ByID(ctx context.Context, id int) (*domain.City, error)
	Create(ctx context.Context, city *domain.City) error
	Update(ctx context.Context, city *domain.City) error
	Delete(ctx context.Context, id int) error
}

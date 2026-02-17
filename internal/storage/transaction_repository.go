package storage

import (
	"context"

	"botmanager/internal/domain"
)

type TransactionRepository interface {
	Create(ctx context.Context, tx *domain.Transaction) error
	ListByCustomer(ctx context.Context, customerID int) ([]domain.Transaction, error)
}

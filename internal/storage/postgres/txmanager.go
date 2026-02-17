package postgres

import (
	"context"
	"database/sql"
)

type TxManager struct {
	db *sql.DB
}

func NewTxManager(db *sql.DB) *TxManager {
	return &TxManager{db: db}
}

func (m *TxManager) WithinTransaction(
	ctx context.Context,
	fn func(ctx context.Context) error,
) error {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, txKey{}, tx)

	if err := fn(ctx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}

type txKey struct{}

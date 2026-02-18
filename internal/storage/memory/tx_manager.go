// Package memory ...
package memory

import (
	"context"
	"sync"
)

// TxManager implements service.TxManager for in-memory storage.
//
// It simulates a transaction by locking shared storage state
// during execution of a function.
type TxManager struct {
	mu *sync.Mutex
}

// NewTxManager creates a new in-memory transaction manager.
func NewTxManager(mu *sync.Mutex) *TxManager {
	return &TxManager{mu: mu}
}

// WithinTransaction executes fn inside a critical section.
// It gurantees exclusive access to shared in-memory state.
func (t *TxManager) WithinTransaction(
	ctx context.Context,
	fn func(ctx context.Context) error,
) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	return fn(ctx)
}

// Package eventbus provide events in memory.
package eventbus

import (
	"context"
	"log/slog"
	"sync"

	"botmanager/internal/domain"
	"botmanager/internal/service"
)

type InMemoryBus struct {
	mu       sync.RWMutex
	handlers map[string][]service.EventHandler
	logger   *slog.Logger
}

func New(logger *slog.Logger) *InMemoryBus {
	return &InMemoryBus{
		handlers: make(map[string][]service.EventHandler),
		logger:   logger,
	}
}

func (b *InMemoryBus) Subscribe(name string, h service.EventHandler) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.handlers[name] = append(b.handlers[name], h)
}

func (b *InMemoryBus) Publish(ctx context.Context, events ...domain.DomainEvent) error {
	for _, e := range events {
		name := e.Name()

		b.mu.RLock()
		handlers := b.handlers[name]
		b.mu.RUnlock()

		for _, h := range handlers {
			if err := h(ctx, e); err != nil {
				b.logger.Error(
					"event handler failed",
					"event", name,
					"error", err,
				)
			}
		}
	}

	return nil
}

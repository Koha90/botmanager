// Package eventbus provide events in memory.
package eventbus

import (
	"log/slog"

	"botmanager/internal/domain"
)

type InMemoryBus struct {
	logger *slog.Logger
}

func New(logger *slog.Logger) *InMemoryBus {
	return &InMemoryBus{logger: logger}
}

func (b *InMemoryBus) Publish(events ...domain.DomainEvent) {
	for _, e := range events {
		b.logger.Info(
			"domain event published",
			"event", e.EventName(),
		)
	}
}

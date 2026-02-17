package service

import (
	"context"

	"botmanager/internal/domain"
)

type EventBus interface {
	Publish(ctx context.Context, events ...domain.DomainEvent)
	Subscribe(eventName string, handler EventHandler)
}

type EventHandler func(ctx context.Context, event domain.DomainEvent) error

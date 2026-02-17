package service

import "botmanager/internal/domain"

type EventBus interface {
	Publish(events ...domain.DomainEvent)
}

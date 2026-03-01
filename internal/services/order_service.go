// Package services contain application use cases.
//
// It coordinates domain logic, repositories and transactions.
// It does not contain business rules.
package services

import (
	"context"
	"log/slog"
	"time"

	"botmanager/internal/domain"
)

// OrderRepository defines persistence contain for Order aggregate.
type OrderRepository interface {
	Save(ctx context.Context, order *domain.Order) error
	FindByID(ctx context.Context, id int) (*domain.Order, error)
}

// EventPublisher publishes domain events.
type EventPublisher interface {
	Publish(ctx context.Context, events ...domain.DomainEvent) error
}

// OrderService orchestrates order use cases.
type OrderService struct {
	repo      OrderRepository
	publisher EventPublisher

	logger *slog.Logger
}

// NewOrderService creates OrderService instance.
func NewOrderService(
	repo OrderRepository,
	publisher EventPublisher,
) *OrderService {
	return &OrderService{
		repo:      repo,
		publisher: publisher,
	}
}

// ConfirmPayment marks order as paid and publishes domain events.
func (s *OrderService) ConfirmPayment(
	ctx context.Context,
	orderID int,
) error {
	order, err := s.repo.FindByID(ctx, orderID)
	if err != nil {
		s.logger.Error("failed to find order", "orderID", orderID, "err", err)
		return err
	}

	if err := order.MarkPaid(time.Now()); err != nil {
		s.logger.Error("failes to mark order as paid", "orderID", orderID, "err", err)
		return err
	}

	if err := s.repo.Save(ctx, order); err != nil {
		s.logger.Error("failed to save order", "err", err)
		return err
	}

	events := order.PullEvents()

	if err := s.publisher.Publish(ctx, events...); err != nil {
		s.logger.Error("failed to publishe order events",
			"orderID", orderID,
			"err", err,
		)
		return err
	}

	s.logger.Info(
		"order successfully paid",
		"orderID", orderID,
	)
	return nil
}

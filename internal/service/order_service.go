// Package service contain application use scases.
//
// It coordinates domain logic, repositories and transactions.
// It does not contain business rules.
package service

import (
	"context"
	"log/slog"

	"botmanager/internal/domain"
)

type OrderRepository interface {
	Create(ctx context.Context, order *domain.Order) error
	ByID(ctx context.Context, id int) (*domain.Order, error)
	ListByCustomer(ctx context.Context, customerID int) ([]domain.Order, error)
	Update(ctx context.Context, order *domain.Order) error
}

type orderCreator interface {
	Create(ctx context.Context, order *domain.Order) error
}

type productReader interface {
	ByID(ctx context.Context, id int) (*domain.Product, error)
}

type orderUpdater interface {
	ByID(ctx context.Context, id int) (*domain.Order, error)
	Update(ctx context.Context, order *domain.Order) error
}

type OrderService struct {
	products productReader
	creator  orderCreator
	updater  orderUpdater
	tx       TxManager
	bus      EventBus
	logger   *slog.Logger
}

func NewOrderService(
	products productReader,
	creator orderCreator,
	updater orderUpdater,
	tx TxManager,
	bus EventBus,
	logger *slog.Logger,
) *OrderService {
	return &OrderService{
		products: products,
		creator:  creator,
		updater:  updater,
		tx:       tx,
		bus:      bus,
		logger:   logger,
	}
}

func (s *OrderService) Create(
	ctx context.Context,
	customerID int,
	productID int,
) (*domain.Order, error) {
	product, err := s.products.ByID(ctx, productID)
	if err != nil {
		return nil, err
	}

	order := domain.NewOrder(customerID, productID, product.Price)

	if err := s.creator.Create(ctx, order); err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderService) Confirm(ctx context.Context, id int) error {
	return s.tx.WithinTransaction(ctx, func(ctx context.Context) error {
		s.logger.Info("confirming order", "order_id", id)

		order, err := s.updater.ByID(ctx, id)
		if err != nil {
			s.logger.Error("failed to load order", "error", err)
			return err
		}

		if err := order.Confirm(); err != nil {
			s.logger.Warn("order confirm rejected", "error", err)
			return err
		}

		if err := s.updater.Update(ctx, order); err != nil {
			s.logger.Error("failed to update order", "error", err)
			return err
		}

		events := order.PullEvents()
		s.bus.Publish(ctx, events...)

		s.logger.Info("order confirmed successfully", "order_id", id)

		return nil
	})
}

func (s *OrderService) Cancel(ctx context.Context, id int) error {
	return s.tx.WithinTransaction(ctx, func(ctx context.Context) error {
		s.logger.Info("cancelling order", "order_id", id)

		order, err := s.updater.ByID(ctx, id)
		if err != nil {
			s.logger.Error("failed to load order", "error", err)
			return err
		}

		if err := order.Cancel(); err != nil {
			s.logger.Warn("order cancel rejected", "error", err)
			return err
		}

		if err := s.updater.Update(ctx, order); err != nil {
			s.logger.Error("failed to update order", "error", err)
			return err
		}

		events := order.PullEvents()
		s.bus.Publish(ctx, events...)

		s.logger.Info("order canceled successfully", "order_id", id)

		return nil
	})
}

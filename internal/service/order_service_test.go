package service

import (
	"context"

	"botmanager/internal/domain"
)

type stubProductReader struct {
	product *domain.Product
	err     error
}

func (s stubProductReader) ByID(ctx context.Context, id int) (*domain.Product, error) {
	return s.product, s.err
}

type stubProductRepository struct {
	order   *domain.Order
	byIDErr error
	saveErr error
	saved   *domain.Order
}

func (s *stubProductRepository) ByID(ctx context.Context, id int) (*domain.Order, error) {
	return s.order, s.byIDErr
}

func (s *stubProductRepository) Save(ctx context.Context, order *domain.Order) error {
	s.saved = order
	return s.saveErr
}

type stubUserRepository struct {
	user    *domain.User
	byIDErr error
	saveErr error
	saved   *domain.User
}

func (s *stubUserRepository) ByID(ctx context.Context, id int) (*domain.User, error) {
	return s.user, s.byIDErr
}

func (s *stubUserRepository) Save(ctx context.Context, user *domain.User) error {
	s.saved = user
	return s.saveErr
}

type stubTxManager struct {
	err error
}

func (s stubTxManager) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	if s.err != nil {
		return s.err
	}
	return fn(ctx)
}

type stubEventBus struct {
	err error
}

func (s stubEventBus) Publish(ctx context.Context, events ...domain.Event) error {
	return s.err
}

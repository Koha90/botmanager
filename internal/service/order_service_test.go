package service

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"

	"botmanager/internal/domain"
)

type fakeProductRepo struct {
	product *domain.Product
}

func (f *fakeProductRepo) ByID(ctx context.Context, id int) (*domain.Product, error) {
	return f.product, nil
}

type fakeOrderRepo struct {
	created *domain.Order
	updated *domain.Order
	order   *domain.Order
}

func (f *fakeOrderRepo) Create(ctx context.Context, order *domain.Order) error {
	f.created = order
	f.order = order
	return nil
}

func (f *fakeOrderRepo) Update(ctx context.Context, order *domain.Order) error {
	f.updated = order
	f.order = order
	return nil
}

func (f *fakeOrderRepo) ByID(ctx context.Context, id int) (*domain.Order, error) {
	return f.order, nil
}

type fakeTx struct {
	called bool
}

func (f *fakeTx) WithinTransaction(ctx context.Context, fn func(ctx context.Context) error) error {
	f.called = true
	return fn(ctx)
}

type fakeBus struct {
	published []domain.DomainEvent
}

func (f *fakeBus) Publish(ctx context.Context, events ...domain.DomainEvent) {
	f.published = append(f.published, events...)
}

func (f *fakeBus) Subscribe(eventName string, handler EventHandler) {}

func TestCreateOrder_SetsCartStatusAndPrice(t *testing.T) {
	ctx := context.Background()

	productRepo := &fakeProductRepo{
		product: &domain.Product{
			ID:    10,
			Price: 1500,
		},
	}

	orderRepo := &fakeOrderRepo{}

	tx := &fakeTx{}
	bus := &fakeBus{}
	logger := slog.Default()

	service := NewOrderService(productRepo, orderRepo, orderRepo, tx, bus, logger)

	order, err := service.Create(ctx, 1, 10)
	require.NoError(t, err)

	require.Equal(t, domain.StatusCart, order.Status())
	require.Equal(t, int64(1500), order.Price())
	require.Equal(t, 1, order.CustomerID())
	require.Equal(t, 10, order.ProductID())
}

func TestConfirmOrder_ChangesStatusAndPublishesEvent(t *testing.T) {
	ctx := context.Background()

	order := domain.NewOrder(1, 10, 1500)

	repo := &fakeOrderRepo{order: order}
	tx := &fakeTx{}
	bus := &fakeBus{}
	logger := slog.Default()

	service := NewOrderService(nil, repo, repo, tx, bus, logger)

	err := service.Confirm(ctx, 1)
	require.NoError(t, err)

	require.True(t, tx.called)
	require.Equal(t, domain.StatusConfirmed, order.Status())
	require.NotNil(t, repo.updated)
	require.NotEmpty(t, bus.published)
}

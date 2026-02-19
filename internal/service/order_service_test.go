package service

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"

	"botmanager/internal/domain"
	"botmanager/internal/storage/memory"
)

type fakeProductRepo struct {
	product *domain.Product
}

func (f *fakeProductRepo) ByID(ctx context.Context, id int) (*domain.Product, error) {
	return f.product, nil
}

type fakeOrderRepo struct {
	created   *domain.Order
	updated   *domain.Order
	order     *domain.Order
	createErr error
	updateErr error
	byIDErr   error
}

func (f *fakeOrderRepo) Create(ctx context.Context, order *domain.Order) error {
	if f.createErr != nil {
		return f.createErr
	}
	f.created = order
	f.order = order
	return nil
}

func (f *fakeOrderRepo) Update(ctx context.Context, order *domain.Order) error {
	if f.updateErr != nil {
		return f.updateErr
	}
	f.updated = order
	f.order = order
	return nil
}

func (f *fakeOrderRepo) ByID(ctx context.Context, id int) (*domain.Order, error) {
	if f.byIDErr != nil {
		return nil, f.byIDErr
	}
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
	mu := &sync.Mutex{}

	productRepo := memory.NewProductRepository(mu)
	orderRepo := memory.NewOrderRepository()

	// seed
	product, _ := domain.NewProduct("Test", 1000)
	_ = productRepo.Create(ctx, product)

	bus := &fakeBus{}

	service := NewOrderService(
		productRepo,
		orderRepo,
		orderRepo,
		memory.NewTxManager(mu),
		bus,
		slog.Default(),
	)

	order, err := service.Create(ctx, 1, product.ID())

	require.NoError(t, err)
	require.Equal(t, domain.StatusCart, order.Status())
	require.Equal(t, int64(1000), order.Price())
	require.Equal(t, 1, order.CustomerID())
	require.Equal(t, product.ID(), order.ProductID())
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

func TestOrderService_Create_ProductNotFound(t *testing.T) {
	ctx := context.Background()
	mu := &sync.Mutex{}

	productRepo := memory.NewProductRepository(mu)
	orderRepo := memory.NewOrderRepository()

	bus := &fakeBus{}

	service := NewOrderService(
		productRepo,
		orderRepo,
		orderRepo,
		memory.NewTxManager(mu),
		bus,
		slog.Default(),
	)

	_, err := service.Create(ctx, 1, 999)

	require.ErrorIs(t, err, domain.ErrProductNotFound)
}

func TestOrderService_Confirm_Success(t *testing.T) {
	ctx := context.Background()
	mu := &sync.Mutex{}

	productRepo := memory.NewProductRepository(mu)
	orderRepo := memory.NewOrderRepository()

	product, _ := domain.NewProduct("Test", 1000)
	_ = productRepo.Create(ctx, product)

	bus := &fakeBus{}

	service := NewOrderService(
		productRepo,
		orderRepo,
		orderRepo,
		memory.NewTxManager(mu),
		bus,
		slog.Default(),
	)

	order, _ := service.Create(ctx, 1, product.ID())

	err := service.Confirm(ctx, order.ID())

	require.NoError(t, err)
}

func TestCreateOrder_SaveFails(t *testing.T) {
	ctx := context.Background()

	product, _ := domain.NewProduct("Test", 1000)

	productRepo := &fakeProductRepo{product: product}
	orderRepo := &fakeOrderRepo{createErr: errors.New("db error")}
	bus := &fakeBus{}

	service := NewOrderService(
		productRepo,
		orderRepo,
		orderRepo,
		&fakeTx{},
		bus,
		slog.Default(),
	)

	_, err := service.Create(ctx, 1, 1)

	require.Error(t, err)
}

func TestConfirmOrder_NotFound(t *testing.T) {
	ctx := context.Background()

	repo := &fakeOrderRepo{byIDErr: errors.New("not found")}
	tx := &fakeTx{}
	bus := &fakeBus{}

	service := newTestService(repo, tx, bus)

	err := service.Confirm(ctx, 1)

	require.Error(t, err)
	require.True(t, tx.called)
}

func TestConfirmOrder_UpdateFails(t *testing.T) {
	ctx := context.Background()

	order := domain.NewOrder(1, 1, 1000)

	repo := &fakeOrderRepo{
		order:     order,
		updateErr: domain.ErrOrderUpdate,
	}

	tx := &fakeTx{}
	bus := &fakeBus{}

	service := newTestService(repo, tx, bus)

	err := service.Confirm(ctx, 1)

	require.Error(t, err)
	require.True(t, tx.called)
}

func TestCancelOrder_Success(t *testing.T) {
	ctx := context.Background()

	order := domain.NewOrder(1, 1, 1000)

	repo := &fakeOrderRepo{order: order}
	tx := &fakeTx{}
	bus := &fakeBus{}

	service := newTestService(repo, tx, bus)

	err := service.Cancel(ctx, 1)

	require.NoError(t, err)
	require.Equal(t, domain.StatusCancelled, order.Status())
	require.NotEmpty(t, bus.published)
}

func TestOrderService_Confirm(t *testing.T) {
	ctx := context.Background()

	tests := []struct {
		name         string
		setupRepo    func() *fakeOrderRepo
		expectError  error
		expectStatus *domain.OrderStatus
		expectEvents bool
	}{
		{
			name: "success",
			setupRepo: func() *fakeOrderRepo {
				order := domain.NewOrder(1, 1, 1000)
				return &fakeOrderRepo{order: order}
			},
			expectError:  nil,
			expectStatus: ptr(domain.StatusConfirmed),
			expectEvents: true,
		},
		{
			name: "order not found",
			setupRepo: func() *fakeOrderRepo {
				return &fakeOrderRepo{byIDErr: domain.ErrOrderNotFound}
			},
			expectError: domain.ErrOrderNotFound,
		},
		{
			name: "already confirmed",
			setupRepo: func() *fakeOrderRepo {
				order := domain.NewOrder(1, 1, 1000)
				_ = order.Confirm()
				return &fakeOrderRepo{order: order}
			},
			expectError:  domain.ErrOrderAlreadyConfirmed,
			expectStatus: ptr(domain.StatusConfirmed),
		},
		{
			name: "update fails",
			setupRepo: func() *fakeOrderRepo {
				order := domain.NewOrder(1, 1, 1000)
				return &fakeOrderRepo{
					order:     order,
					updateErr: domain.ErrOrderUpdate,
				}
			},
			expectError: domain.ErrOrderUpdate,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.setupRepo()
			tx := &fakeTx{}
			bus := &fakeBus{}

			service := NewOrderService(
				nil,
				repo,
				repo,
				tx,
				bus,
				slog.Default(),
			)

			err := service.Confirm(ctx, 1)

			if tt.expectError != nil {
				require.Error(t, err)
				require.ErrorIs(t, err, tt.expectError)
				require.Empty(t, bus.published)
				return
			}
			require.NoError(t, err)

			require.True(t, tx.called)

			if tt.expectStatus != nil {
				require.NotNil(t, repo.order)
				require.Equal(t, *tt.expectStatus, repo.order.Status())
			}

			if tt.expectEvents {
				require.Len(t, bus.published, 1)
			} else {
				require.Empty(t, bus.published)
			}
		})
	}
}

// helpers
func ptr[T any](v T) *T {
	return &v
}

func newTestService(repo *fakeOrderRepo, tx *fakeTx, bus *fakeBus) *OrderService {
	return NewOrderService(nil, repo, repo, tx, bus, slog.Default())
}

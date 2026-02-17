package service

import (
	"context"
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
}

func (f *fakeOrderRepo) Create(ctx context.Context, order *domain.Order) error {
	f.created = order
	return nil
}

func TestCreateOrder_SetsCartStatusAndPrice(t *testing.T) {
	ctx := context.Background()

	productRepo := &fakeProductRepo{
		product: &domain.Product{
			ID:    10,
			Price: 1500,
		},
	}

	orderRepo := &fakeOrderRepo{}

	service := NewOrderService(productRepo, orderRepo)

	order, err := service.Create(ctx, 1, 10)
	require.NoError(t, err)

	require.Equal(t, domain.OrderStatusCart, order.Status)
	require.Equal(t, int64(1500), order.PriceAtPurchase)
	require.Equal(t, 1, order.CustomerID)
	require.Equal(t, 10, order.ProductID)
}

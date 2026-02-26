package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func RehydrateOrder(
	id int,
	customerID int,
	productID int,
	price int64,
	status OrderStatus,
	version int,
) *Order {
	return &Order{
		id:         id,
		customerID: customerID,
		productID:  productID,
		price:      price,
		status:     status,
		version:    version,
	}
}

func TestNewOrderStatus(t *testing.T) {
	tests := []struct {
		name      string
		setup     func() *Order
		expectErr error
	}{
		{
			name: "success",
			setup: func() *Order {
				o := NewOrder(1, 1, 1, 1, 200)
				return o
			},
		},
		{
			name: "already confirmed",
			setup: func() *Order {
				o := NewOrder(1, 1, 1, 1, 200)
				_ = o.Confirm()
				return o
			},
			expectErr: ErrOrderAlreadyConfirmed,
		},
		{
			name: "invalid state after cancel",
			setup: func() *Order {
				o := NewOrder(1, 1, 1, 1, 200)
				_ = o.Cancel()
				return o
			},
			expectErr: ErrOrderAlreadyCanceled,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := tt.setup()

			err := o.Confirm()

			if tt.expectErr != nil {
				require.ErrorIs(t, err, tt.expectErr)
				return
			}

			require.NoError(t, err)
			require.Equal(t, StatusConfirmed, o.Status())
			require.Equal(t, 2, o.Version())

			events := o.PullEvents()
			require.Len(t, events, 1)
			require.Equal(t, NameOrderConfirm, events[0].Name())
			require.Empty(t, o.PullEvents())
		})
	}
}

func TestOrder_Cancel_Table(t *testing.T) {
	tests := []struct {
		name      string
		setup     func() *Order
		expectErr error
	}{
		{
			name: "success",
			setup: func() *Order {
				o := NewOrder(1, 1, 1, 1, 200)
				return o
			},
		},
		{
			name: "already canceled",
			setup: func() *Order {
				o := NewOrder(1, 1, 1, 1, 200)
				_ = o.Cancel()
				return o
			},
			expectErr: ErrOrderAlreadyCanceled,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := tt.setup()

			err := o.Cancel()

			if tt.expectErr != nil {
				require.ErrorIs(t, err, tt.expectErr)
				return
			}

			require.NoError(t, err)
			require.Equal(t, StatusCancelled, o.Status())
			require.Equal(t, 2, o.Version())

			events := o.PullEvents()
			require.Len(t, events, 1)
			require.Equal(t, NameOrderCancel, events[0].Name())
			require.Empty(t, o.PullEvents())
		})
	}
}

func TestOrderStatus_Confirm_Table(t *testing.T) {
	tests := []struct {
		status    OrderStatus
		expectErr error
	}{
		{StatusCart, nil},
		{StatusConfirmed, ErrOrderAlreadyConfirmed},
		{StatusCancelled, ErrOrderAlreadyCanceled},
	}

	for _, tt := range tests {
		_, err := tt.status.Confirm()
		if tt.expectErr != nil {
			require.ErrorIs(t, err, tt.expectErr)
			continue
		}
		require.NoError(t, err)
	}
}

func TestOrderStatus_Cancel_Table(t *testing.T) {
	tests := []struct {
		status    OrderStatus
		expectErr error
	}{
		{StatusCart, nil},
		{StatusCancelled, ErrOrderAlreadyCanceled},
		{StatusConfirmed, ErrInvalidOrderState},
	}

	for _, tt := range tests {
		_, err := tt.status.Cancel()
		if tt.expectErr != nil {
			require.ErrorIs(t, err, tt.expectErr)
			continue
		}
		require.NoError(t, err)
	}
}

func TestOrderStatus_OtherMethods(t *testing.T) {
	require.True(t, StatusCancelled.IsFinal())
	require.False(t, StatusCart.IsFinal())

	require.True(t, StatusConfirmed.CanShip())
	require.False(t, StatusCart.CanShip())

	require.Equal(t, "cart", StatusCart.String())
}

func TestNewOrder(t *testing.T) {
	o := NewOrder(1, 1, 1, 1, 200)

	require.Equal(t, StatusCart, o.Status())
	require.Equal(t, 1, o.Version())
	require.Equal(t, int64(200), o.Price())
	require.Empty(t, o.PullEvents())
}

func TestNewOrderStatus_FromString(t *testing.T) {
	tests := []struct {
		value     string
		expectErr bool
	}{
		{"cart", false},
		{"confirmed", false},
		{"canceled", false},
		{"invalid", true},
	}

	for _, tt := range tests {
		_, err := NewOrderStatus(tt.value)
		if tt.expectErr {
			require.Error(t, err)
			continue
		}
		require.NoError(t, err)
	}
}

func TestOrder_Confirm_FromConfirmed(t *testing.T) {
	o := NewOrder(1, 1, 1, 1, 200)

	require.NoError(t, o.Confirm())

	err := o.Confirm()
	require.ErrorIs(t, err, ErrOrderAlreadyConfirmed)

	require.Equal(t, 2, o.Version())
}

func TestOrder_Cancel_FromConfirmed(t *testing.T) {
	o := NewOrder(1, 1, 1, 1, 200)

	require.NoError(t, o.Confirm())

	err := o.Cancel()
	require.ErrorIs(t, err, ErrInvalidOrderState)
}

func TestOrder_Getters(t *testing.T) {
	o := RehydrateOrder(99, 10, 20, 200, StatusCart, 1)

	require.Equal(t, 99, o.ID())
	require.Equal(t, 10, o.CustomerID())
	require.Equal(t, 20, o.ProductID())
}

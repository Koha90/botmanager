package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewOrderStatus(t *testing.T) {
	tests := []struct {
		name      string
		setup     func() *Order
		expectErr error
	}{
		{
			name: "success",
			setup: func() *Order {
				o := NewOrder(1, 1, 1000)
				o.SetID(42)
				return o
			},
		},
		{
			name: "already confirmed",
			setup: func() *Order {
				o := NewOrder(1, 1, 1000)
				o.SetID(42)
				_ = o.Confirm()
				return o
			},
			expectErr: ErrOrderAlreadyConfirmed,
		},
		{
			name: "invalid state after cancel",
			setup: func() *Order {
				o := NewOrder(1, 1, 1000)
				o.SetID(42)
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
			require.Equal(t, NameConfirm, events[0].Name())
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
				o := NewOrder(1, 1, 1000)
				o.SetID(42)
				return o
			},
		},
		{
			name: "already canceled",
			setup: func() *Order {
				o := NewOrder(1, 1, 1000)
				o.SetID(42)
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
			require.Equal(t, NameCancel, events[0].Name())
			require.Empty(t, o.PullEvents())
		})
	}
}

func TestOrderStatus_CanConfirm(t *testing.T) {
	require.True(t, StatusCart.CanConfirm())
	require.False(t, StatusConfirmed.CanConfirm())
	require.False(t, StatusCancelled.CanConfirm())
}

func TestOrderStatus_CanCancel(t *testing.T) {
	require.True(t, StatusCart.CanCancel())
	require.False(t, StatusConfirmed.CanCancel())
	require.False(t, StatusCancelled.CanCancel())
}

func TestOrderStatus_OtherMethods(t *testing.T) {
	require.True(t, StatusCancelled.IsFinal())
	require.False(t, StatusCart.IsFinal())

	require.True(t, StatusConfirmed.CanShip())
	require.False(t, StatusCart.CanShip())

	require.Equal(t, "cart", StatusCart.String())
}

func TestNewOrder(t *testing.T) {
	o := NewOrder(10, 20, 5000)

	require.Equal(t, StatusCart, o.Status())
	require.Equal(t, 1, o.Version())
	require.Equal(t, int64(5000), o.Price())
	require.Empty(t, o.PullEvents())
}

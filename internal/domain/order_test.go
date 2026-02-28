package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewOrder(t *testing.T) {
	_, err := NewOrder(0, nil, time.Now())
	require.ErrorIs(t, err, ErrInvalidOrderUserID)

	_, err = NewOrder(1, []OrderItem{}, time.Now())
	require.ErrorIs(t, err, ErrOrderEmpty)

	items := []OrderItem{
		{variantID: 1, quantity: 2, price: 100},
	}

	o, err := NewOrder(1, items, time.Now())
	require.NoError(t, err)
	require.Equal(t, int64(200), o.Total())
	require.Equal(t, OrderStatusPending, o.Status())
}

func TestOrder_MarkPaid(t *testing.T) {
	items := []OrderItem{
		{variantID: 1, quantity: 1, price: 100},
	}

	o, _ := NewOrder(1, items, time.Now())

	err := o.MarkPaid(time.Now())
	require.NoError(t, err)
	require.Equal(t, OrderStatusPaid, o.Status())

	err = o.MarkPaid(time.Now())
	require.ErrorIs(t, err, ErrOrderAlreadyPaid)
}

func TestOrder_Cancel(t *testing.T) {
	items := []OrderItem{
		{variantID: 1, quantity: 1, price: 100},
	}

	o, _ := NewOrder(1, items, time.Now())

	require.NoError(t, o.Cancel())
	require.Equal(t, OrderStatusCancelled, o.Status())

	err := o.MarkPaid(time.Now())
	require.ErrorIs(t, err, ErrOrderAlreadyCancelled)
}

func TestOrder_CannotCancelPaid(t *testing.T) {
	items := []OrderItem{
		{variantID: 1, quantity: 1, price: 100},
	}

	o, _ := NewOrder(1, items, time.Now())
	_ = o.MarkPaid(time.Now())

	err := o.Cancel()
	require.ErrorIs(t, err, ErrOrderAlreadyPaid)
}

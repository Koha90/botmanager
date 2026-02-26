package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrder_Confirm_Table(t *testing.T) {
	tests := []struct {
		name      string
		setup     func() *Order
		expectErr error
	}{
		{
			name:  "success",
			setup: func() *Order { return NewOrder(1, 1, 1, 200) },
		}, {
			name: "already confirmed",
			setup: func() *Order {
				o := NewOrder(1, 1, 1, 200)
				_ = o.Confirm()
				return o
			},
			expectErr: ErrOrderAlreadyConfirmed,
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

			require.Equal(t, int64(200), o.Price())

			require.Empty(t, o.PullEvents())
		})
	}
}

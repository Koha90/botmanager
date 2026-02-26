package domain

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewProductVariant(t *testing.T) {
	tests := []struct {
		name      string
		id        int
		packSize  string
		district  int
		price     int64
		expectErr error
	}{
		{"success", 1, "250g", 1, 100, nil},
		{"invalid pack", 1, "", 1, 100, ErrInvalidPackSize},
		{"invalid district", 1, "250g", 0, 100, ErrInvalidDisctrictID},
		{"invalid price", 1, "250g", 1, 0, ErrInvalidProductPrice},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewProductVariant(
				tt.packSize,
				tt.district,
				tt.price,
			)

			if tt.expectErr != nil {
				fmt.Println(err)
				require.ErrorIs(t, err, tt.expectErr)
				return
			}

			require.NoError(t, err)
		})
	}
}

func TestProductVariant_PriceInvalid(t *testing.T) {
	v, _ := NewProductVariant("250g", 1, 100)

	require.NoError(t, v.ChangePrice(200))
	require.Equal(t, int64(200), v.price)

	require.ErrorIs(t, v.ChangePrice(0), ErrInvalidProductPrice)
}

func TestProductVariant_ChangePackSize(t *testing.T) {
	v, _ := NewProductVariant("250g", 1, 100)

	require.NoError(t, v.ChangePackSize("500g"))
	require.Equal(t, "500g", v.packSize)

	require.ErrorIs(t, v.ChangePackSize(""), ErrInvalidPackSize)
}

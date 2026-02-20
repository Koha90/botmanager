package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProduct_Rename(t *testing.T) {
	p, _ := NewProduct("coffee", 100)

	require.NoError(t, p.Rename("latte"))
	require.Equal(t, "latte", p.Name())

	require.ErrorIs(t, p.Rename(""), ErrInvalidProductName)
}

func TestProduct_ChangePrice(t *testing.T) {
	p, _ := NewProduct("coffee", 100)

	require.NoError(t, p.ChangePrice(200))
	require.Equal(t, int64(200), p.Price())

	require.ErrorIs(t, p.ChangePrice(0), ErrInvalidProductPrice)
}

func TestNewProduct_Invalid(t *testing.T) {
	_, err := NewProduct("", 100)
	require.ErrorIs(t, err, ErrInvalidProductName)

	_, err = NewProduct("coffee", 0)
	require.ErrorIs(t, err, ErrInvalidProductPrice)
}

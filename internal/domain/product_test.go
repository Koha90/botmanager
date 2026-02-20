package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProduct_Rename(t *testing.T) {
	p, _ := NewProduct("coffee", 1, "", "")

	require.NoError(t, p.Rename("latte"))
	require.Equal(t, "latte", p.Name())

	require.ErrorIs(t, p.Rename(""), ErrInvalidProductName)
}

func TestNewProduct_Invalid(t *testing.T) {
	_, err := NewProduct("", 1, "", "")
	require.ErrorIs(t, err, ErrInvalidProductName)
}

package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProduct_AddVariant(t *testing.T) {
	p, _ := NewProduct("Test Product", 1, "Description", "")

	require.Equal(t, 1, p.Version())

	err := p.AddVariant(1, "1p", 1, 100)
	require.NoError(t, err)
	require.Equal(t, 2, p.Version())
}

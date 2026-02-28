package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewStock(t *testing.T) {
	_, err := NewStock(0, 1, 10)
	require.ErrorIs(t, err, ErrInvalidWarehouseID)

	_, err = NewStock(1, 0, 10)
	require.ErrorIs(t, err, ErrInvalidVariantID)

	_, err = NewStock(1, 1, -1)
	require.ErrorIs(t, err, ErrInvalidQuantity)

	s, err := NewStock(1, 1, 10)
	require.NoError(t, err)
	require.Equal(t, 10, s.Quantity())
	require.Equal(t, 0, s.Reserved())
	require.Equal(t, 10, s.Available())
	require.Equal(t, 1, s.Version())
}

func TestStock_Reserve(t *testing.T) {
	s, _ := NewStock(1, 1, 10)

	err := s.Reserve(5)
	require.NoError(t, err)
	require.Equal(t, 5, s.Reserved())
	require.Equal(t, 5, s.Available())
	require.Equal(t, 2, s.Version())

	err = s.Reserve(6)
	require.ErrorIs(t, err, ErrInsufficientStock)
}

func TestStock_Release(t *testing.T) {
	s, _ := NewStock(1, 1, 10)
	_ = s.Reserve(5)

	err := s.Release(3)
	require.NoError(t, err)
	require.Equal(t, 2, s.Reserved())
	require.Equal(t, 8, s.Available())

	err = s.Release(5)
	require.ErrorIs(t, err, ErrInvalidReleaseQuantity)
}

func TestStock_Decrease(t *testing.T) {
	s, _ := NewStock(1, 1, 10)
	_ = s.Reserve(5)

	err := s.Decrease(4)
	require.NoError(t, err)

	require.Equal(t, 6, s.Quantity())
	require.Equal(t, 1, s.Reserved())
	require.Equal(t, 5, s.Available())

	err = s.Decrease(2)
	require.ErrorIs(t, err, ErrInsufficientReservedStock)
}

func TestStock_AvailableCalculation(t *testing.T) {
	s, _ := NewStock(1, 1, 20)

	_ = s.Reserve(7)
	require.Equal(t, 13, s.Available())

	_ = s.Release(2)
	require.Equal(t, 15, s.Available())
}

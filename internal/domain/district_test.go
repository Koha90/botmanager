package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewDistrict(t *testing.T) {
	_, err := NewDistrict(0, "x")
	require.ErrorIs(t, err, ErrInvalidCityID)

	_, err = NewDistrict(1, "")
	require.ErrorIs(t, err, ErrInvalidDistrictName)

	d, err := NewDistrict(1, "Center")
	require.NoError(t, err)
	require.Equal(t, "Center", d.Name())
}

func TestDistrict_Rename(t *testing.T) {
	d, _ := NewDistrict(1, "Old")

	err := d.Rename("")
	require.ErrorIs(t, err, ErrInvalidDistrictName)

	require.NoError(t, d.Rename("New"))
	require.Equal(t, "New", d.Name())
}

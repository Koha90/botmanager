package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCity(t *testing.T) {
	_, err := NewCity("")
	require.ErrorIs(t, err, ErrInvalidCityName)

	c, err := NewCity("London")
	require.NoError(t, err)
	require.Equal(t, "London", c.Name())
}

func TestCity_Rename(t *testing.T) {
	c, _ := NewCity("Old")

	err := c.Rename("")
	require.ErrorIs(t, err, ErrInvalidCityName)

	require.NoError(t, c.Rename("New"))
	require.Equal(t, "New", c.Name())
}

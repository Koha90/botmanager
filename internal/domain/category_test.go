package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCategory(t *testing.T) {
	_, err := NewCategory("", "desc")
	require.ErrorIs(t, err, ErrInvalidCategoryName)

	c, err := NewCategory("Tea", "desc")
	require.NoError(t, err)
	require.Equal(t, "Tea", c.Name())
}

func TestCategory_Rename(t *testing.T) {
	c, _ := NewCategory("Old", "")

	err := c.Rename("")
	require.ErrorIs(t, err, ErrInvalidCategoryName)

	require.NoError(t, c.Rename("New"))
	require.Equal(t, "New", c.Name())
}

func TestCategory_UpdateDescription(t *testing.T) {
	c, _ := NewCategory("Tea", "old")
	c.UpdateDescription("new")

	require.Equal(t, "new", c.Desecription())
}

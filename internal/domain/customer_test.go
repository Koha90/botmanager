package domain

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCustomer(t *testing.T) {
	c := NewCustomer(123, "john")

	require.Equal(t, int64(123), c.TelegramID())
	require.Equal(t, "john", c.username)
}

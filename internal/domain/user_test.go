package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewUser_DefaultRole(t *testing.T) {
	u, err := NewUser(NewUserParams{
		Email:        "a@a.com",
		PasswordHash: "hash",
	})
	require.NoError(t, err)
	require.Equal(t, RoleCustomer, u.role)
	require.False(t, u.isEnabled)
}

func TestNewUser_InvalidRole(t *testing.T) {
	_, err := NewUser(NewUserParams{
		Email:        "a@a.com",
		PasswordHash: "hash",
		Role:         "superman",
	})
	require.ErrorIs(t, err, ErrInvalidRole)
}

func TestNewUser_InvalidCredentials(t *testing.T) {
	_, err := NewUser(NewUserParams{})
	require.ErrorIs(t, err, ErrInvalidCredentials)
}

func TestNewUser_AdminAutoEnabled(t *testing.T) {
	u, err := NewUser(NewUserParams{
		Email:        "admin@a.com",
		PasswordHash: "hash",
		Role:         RoleAdmin,
	})
	require.NoError(t, err)
	require.True(t, u.isEnabled)
}

func TestUser_CanUseAdminPanel(t *testing.T) {
	now := time.Now()
	exp := now.Add(time.Hour)

	u, _ := NewUser(NewUserParams{
		Email:        "admin@a.com",
		PasswordHash: "hash",
		Role:         RoleAdmin,
	})
	u.GrantAdminAccess(exp)

	require.True(t, u.CanUseAdminPanel(now))
	require.False(t, u.CanUseAdminPanel(exp.Add(time.Hour)))
}

func TestUser_EnableDisable(t *testing.T) {
	u, _ := NewUser(NewUserParams{
		Email:        "a@a.com",
		PasswordHash: "hash",
	})

	u.Enable()
	require.True(t, u.isEnabled)

	u.Disable()
	require.False(t, u.isEnabled)
}

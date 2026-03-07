package domain

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNewUser(t *testing.T) {
	t.Run("default role is customer", func(t *testing.T) {
		tgID := int64(123)

		user, err := NewUser(NewUserParams{
			TgID: &tgID,
		})
		if err != nil {
			t.Fatalf("expected no error,  got %v", err)
		}

		if user.role != RoleCustomer {
			t.Fatalf("expected role %q, got %q", RoleCustomer, user.role)
		}

		if user.isEnabled {
			t.Fatal("expected customer to be disabled by default")
		}
	})

	t.Run("invalid role returns error", func(t *testing.T) {
		_, err := NewUser(NewUserParams{
			Email:        "a@b.c",
			PasswordHash: "hash",
			Role:         Role("superman"),
		})
		if !errors.Is(err, ErrInvalidRole) {
			t.Fatalf("expected ErrInvalidRole, got %v", err)
		}
	})

	t.Run("missing credentials returns error", func(t *testing.T) {
		_, err := NewUser(NewUserParams{})
		if !errors.Is(err, ErrInvalidCredentials) {
			t.Fatalf("expected ErrInvalidCredentials, got %v", err)
		}
	})

	t.Run("admin is enabled by default", func(t *testing.T) {
		user, err := NewUser(NewUserParams{
			Email:        "admin@site.dev",
			PasswordHash: "hash",
			Role:         RoleAdmin,
		})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if !user.isEnabled {
			t.Fatal("expected admin to be enabled by default")
		}
	})
}

func TestUserAddBalance(t *testing.T) {
	tgID := int64(123)

	user, err := NewUser(NewUserParams{
		TgID: &tgID,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if err := user.AddBalance(500); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.balance != 500 {
		t.Fatalf("expected balance 500, got %d", user.balance)
	}
}

func TestUserDeductBalance(t *testing.T) {
	tgID := int64(123)

	user, err := NewUser(NewUserParams{
		TgID: &tgID,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if err := user.AddBalance(500); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if err := user.DeductBalance(200); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user.balance != 300 {
		t.Fatalf("expected balance 300, got %d", user.balance)
	}
}

func TestUserDeductBalance_Insufficient(t *testing.T) {
	tgID := int64(123)

	user, err := NewUser(NewUserParams{
		TgID: &tgID,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = user.DeductBalance(100)
	if !errors.Is(err, ErrInsufficientBalance) {
		t.Fatalf("expected ErrInsufficientBalance, got %v", err)
	}
}

func TestUserCanUseAdminPanel(t *testing.T) {
	now := time.Now()

	t.Run("admin with valid access can use panel", func(t *testing.T) {
		user, err := NewUser(NewUserParams{
			Email:        "admin@site.dev",
			PasswordHash: "hash",
			Role:         RoleAdmin,
		})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		user.GrantAdminAccess(now.Add(time.Hour))

		if !user.CanUseAdminPanel(now) {
			t.Fatal("expected amin to have access")
		}
	})

	t.Run("customer cannot use panel", func(t *testing.T) {
		tgID := int64(123)

		user, err := NewUser(NewUserParams{
			TgID: &tgID,
		})
		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if user.CanUseAdminPanel(now) {
			t.Fatal("expected customer to have no access")
		}
	})
}

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

func TestNewUser_TelegramAuth(t *testing.T) {
	tgID := int64(123456789)
	tgName := "koha"

	u, err := NewUser(NewUserParams{
		TgID:   &tgID,
		TgName: tgName,
	})
	require.NoError(t, err)

	gotID, ok := u.TelegramID()
	require.True(t, ok)
	require.Equal(t, tgID, gotID)

	gotName, ok := u.TelegramName()
	require.True(t, ok)
	require.Equal(t, tgName, gotName)

	require.Equal(t, RoleCustomer, u.role)
	require.False(t, u.isEnabled)
}

func TestNewUser_AddBalance(t *testing.T) {
	u, err := NewUser(NewUserParams{
		Email:        "a@a.com",
		PasswordHash: "hash",
	})
	require.NoError(t, err)

	err = u.AddBalance(500)
	require.NoError(t, err)
	require.EqualValues(t, 500, u.balance)
}

func TestUser_AddBalance_InvalidAmount(t *testing.T) {
	tests := []struct {
		name   string
		amount int64
	}{
		{name: "zero", amount: 0},
		{name: "negative", amount: -100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := NewUser(NewUserParams{
				Email:        "a@a.com",
				PasswordHash: "hash",
			})
			require.NoError(t, err)

			err = u.AddBalance(tt.amount)
			require.ErrorIs(t, err, ErrInvalidAmount)
			require.EqualValues(t, 0, u.balance)
		})
	}
}

func TestUser_DeductBalance(t *testing.T) {
	u, err := NewUser(NewUserParams{
		Email:        "a@a.com",
		PasswordHash: "hash",
	})
	require.NoError(t, err)

	err = u.AddBalance(500)
	require.NoError(t, err)

	err = u.DeductBalance(200)
	require.NoError(t, err)
	require.EqualValues(t, 300, u.balance)
}

func TestUser_DeductBalance_InvalidAmount(t *testing.T) {
	tests := []struct {
		name   string
		amount int64
	}{
		{name: "zero", amount: 0},
		{name: "negative", amount: -50},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := NewUser(NewUserParams{
				Email:        "a@a.com",
				PasswordHash: "hash",
			})
			require.NoError(t, err)

			err = u.DeductBalance(tt.amount)
			require.ErrorIs(t, err, ErrInvalidAmount)
		})
	}
}

func TestUser_DeductBalance_InsufficientBalance(t *testing.T) {
	u, err := NewUser(NewUserParams{
		Email:        "a@a.com",
		PasswordHash: "hash",
	})
	require.NoError(t, err)

	err = u.AddBalance(100)
	require.NoError(t, err)

	err = u.DeductBalance(200)
	require.ErrorIs(t, err, ErrInsufficientBalance)
	require.EqualValues(t, 100, u.balance)
}

func TestUser_GrantAdminAccess_CustomerIgnored(t *testing.T) {
	now := time.Now()

	u, err := NewUser(NewUserParams{
		Email:        "user@a.com",
		PasswordHash: "hash",
	})
	require.NoError(t, err)

	u.GrantAdminAccess(now.Add(time.Hour))

	require.Nil(t, u.adminAccessExpiresAt)
	require.False(t, u.CanUseAdminPanel(now))
}

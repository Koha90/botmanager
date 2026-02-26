package domain

import (
	"time"
)

type Role string

const (
	RoleCustomer Role = "customer"
	RoleAdmin    Role = "admin"
)

// User represent user of the application.
type User struct {
	id                   int
	tgID                 *int64
	tgName               string
	email                string
	passwordHash         string
	role                 Role
	balance              int64
	isEnabled            bool
	adminAccessExpiresAt *time.Time
	createdAt            time.Time
	updatedAt            time.Time
}

type NewUserParams struct {
	TgID         *int64
	TgName       string
	Email        string
	PasswordHash string
	Role         Role
}

// NewUser created a new user of the application.
func NewUser(p NewUserParams) (*User, error) {
	now := time.Now()

	if p.Role == "" {
		p.Role = RoleCustomer
	}

	if p.Role != RoleCustomer && p.Role != RoleAdmin {
		return nil, ErrInvalidRole
	}

	if p.TgID == nil && (p.Email == "" || p.PasswordHash == "") {
		return nil, ErrInvalidCredentials
	}

	user := &User{
		tgID:      p.TgID,
		tgName:    p.TgName,
		email:     p.Email,
		role:      p.Role,
		isEnabled: false,
		balance:   0,
		createdAt: now,
		updatedAt: now,
	}

	if user.role == RoleAdmin {
		user.isEnabled = true
	}

	return user, nil
}

// CanUseAdminPanel returns wheter the user can use the admin panel.
func (u *User) CanUseAdminPanel(now time.Time) bool {
	if u.role != RoleAdmin || !u.isEnabled {
		return false
	}

	if u.adminAccessExpiresAt == nil {
		return false
	}

	return now.Before(*u.adminAccessExpiresAt)
}

// ---- SETTERS ----

// Enable sets enabled and updates user.
func (u *User) Enable() {
	u.isEnabled = true
	u.updatedAt = time.Now()
}

// Disable sets disable user and updates user.
func (u *User) Disable() {
	u.isEnabled = false
	u.updatedAt = time.Now()
}

// GrantAdminAccess ...
func (u *User) GrantAdminAccess(until time.Time) {
	if u.role != RoleAdmin {
		return
	}
	u.adminAccessExpiresAt = &until
	u.updatedAt = time.Now()
}

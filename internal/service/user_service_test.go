package service

import (
	"context"
	"testing"

	"botmanager/internal/domain"
)

func TestUserService_Create(t *testing.T) {
	repo := &stubUserRepository{}
	svc := NewUserService(
		repo,
		stubTxManager{},
		stubEventBus{},
		nil,
	)

	tgID := int64(123)

	user, err := svc.Create(context.Background(), domain.NewUserParams{
		TgID: &tgID,
	})
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if user == nil {
		t.Fatal("expected user, got nil")
	}

	if repo.saved == nil {
		t.Fatal("expected user to be saved")
	}
}

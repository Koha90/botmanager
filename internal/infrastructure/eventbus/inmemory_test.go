package eventbus

import (
	"context"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/require"

	"botmanager/internal/domain"
)

func TestInMemoryBus_MultipleEvents(t *testing.T) {
	bus := New(slog.Default())

	count := 0

	bus.Subscribe(domain.OrderConfirm, func(ctx context.Context, event domain.DomainEvent) error {
		count++
		return nil
	})

	e1 := domain.NewOrderConfirmed(1)
	e2 := domain.NewOrderConfirmed(2)

	_ = bus.Publish(context.Background(), e1, e2)

	require.Equal(t, 2, count)
}

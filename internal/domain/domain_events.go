package domain

import "time"

// DomainEvent represent a fact that happened inside the domain.
//
// Events are generate by aggregates and later published
// by the application layer.
type DomainEvent interface {
	Name() string
	OccurredAt() time.Time
}

const (
	NameOrderConfirm string = "order_confirmed"
	NameOrderCancel  string = "order_cancelled"
)

type OrderConfirmed struct {
	OrderID int
	at      time.Time
}

type OrderCanceled struct {
	OrderID int
	at      time.Time
}

func NewOrderConfirmed(orderID int) OrderConfirmed {
	return OrderConfirmed{
		OrderID: orderID,
		at:      time.Now(),
	}
}

func (e OrderConfirmed) Name() string {
	return NameOrderConfirm
}

func (e OrderConfirmed) OccurredAt() time.Time {
	return e.at
}

func NewOrderCanceled(orderID int) OrderCanceled {
	return OrderCanceled{
		OrderID: orderID,
		at:      time.Now(),
	}
}

func (e OrderCanceled) Name() string {
	return NameOrderCancel
}

func (e OrderCanceled) OccurredAt() time.Time {
	return e.at
}

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
	// Orders
	NameOrderConfirmed string = "order_confirmed"
	NameOrderCanceled  string = "order_cancelled"

	// Products
	NameProductVariantAdded    string = "product variant added"
	NameProductVariantArchived string = "product variant archived"
)

type ProductVariantAdded struct {
	ProductVariantID int
	at               time.Time
}

type ProductVariantArchived struct {
	ProductVariantID int
	at               time.Time
}

type OrderConfirmed struct {
	OrderID int
	at      time.Time
}

type OrderCanceled struct {
	OrderID int
	at      time.Time
}

func NewProductVariantAdded(productVariantID int) ProductVariantAdded {
	return ProductVariantAdded{
		ProductVariantID: productVariantID,
		at:               time.Now(),
	}
}

func (v ProductVariantAdded) Name() string {
	return NameProductVariantAdded
}

func (v ProductVariantAdded) OccurredAt() time.Time {
	return v.at
}

func NewVariantArchived(productVariantID int) ProductVariantArchived {
	return ProductVariantArchived{
		ProductVariantID: productVariantID,
		at:               time.Now(),
	}
}

func (v ProductVariantArchived) Name() string {
	return NameProductVariantArchived
}

func (v ProductVariantArchived) OccurredAt() time.Time {
	return v.at
}

// NewOrderConfirmed set time where order was confirmed.
func NewOrderConfirmed(orderID int) OrderConfirmed {
	return OrderConfirmed{
		OrderID: orderID,
		at:      time.Now(),
	}
}

func (e OrderConfirmed) Name() string {
	return NameOrderConfirmed
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
	return NameOrderCanceled
}

func (e OrderCanceled) OccurredAt() time.Time {
	return e.at
}

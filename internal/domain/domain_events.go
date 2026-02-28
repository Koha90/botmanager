package domain

import "time"

// DomainEvent represent a business fact that occured inside the domain.
//
// Domain events are produced by aggregates to signal that something
// meaningful has happened. They are later collected and published
// by the application layer.
//
// The domain layer must not publish events directly.
// It only records them.
type DomainEvent interface {
	Name() string
	OccurredAt() time.Time
}

const (
	// Orders
	NameOrderConfirmed string = "order_paid"
	NameOrderCancelled string = "order_cancelled"

	// Products
	NameProductVariantAdded    string = "product_variant_added"
	NameProductVariantArchived string = "product_variant_archived"
)

// --------------------
// Product Events
// --------------------

// ProductVariantAdded is emitted when a new product variant
// is successfully added to a product aggregate.
type ProductVariantAdded struct {
	ProductVariantID int
	at               time.Time
}

// ProductVariantArchived is emitted when a product variant
// is archived and becomes unavailable for purchase.
type ProductVariantArchived struct {
	ProductVariantID int
	at               time.Time
}

// NewProductVariantAdded creates ProductVariantAdded event
// with current timestamp.
func NewProductVariantAdded(productVariantID int) ProductVariantAdded {
	return ProductVariantAdded{
		ProductVariantID: productVariantID,
		at:               time.Now(),
	}
}

// Name returns event type identifier.
func (v ProductVariantAdded) Name() string {
	return NameProductVariantAdded
}

// OccurredAt returns event timestamp.
func (v ProductVariantAdded) OccurredAt() time.Time {
	return v.at
}

// NewVariantArchived creates ProductVarianArchived event
// with current timestamp.
func NewVariantArchived(productVariantID int) ProductVariantArchived {
	return ProductVariantArchived{
		ProductVariantID: productVariantID,
		at:               time.Now(),
	}
}

// Name returns event type identifier.
func (v ProductVariantArchived) Name() string {
	return NameProductVariantArchived
}

// OccurredAt returns event timestamp.
func (v ProductVariantArchived) OccurredAt() time.Time {
	return v.at
}

// --------------------
// Order Events
// --------------------

// OrderConfirmed is emitted when an order
// transitions from pending to paid state.
type OrderPaid struct {
	OrderID int
	at      time.Time
}

// OrderCancelled is emitted when an order
// transitions from pending to cancelled state.
type OrderCancelled struct {
	OrderID int
	at      time.Time
}

// NewOrderPaid creates OrderPaid event
// with current timestamp.
func NewOrderPaid(orderID int) OrderPaid {
	return OrderPaid{
		OrderID: orderID,
		at:      time.Now(),
	}
}

// Name returns event type identifier.
func (e OrderPaid) Name() string {
	return NameOrderConfirmed
}

// OccurredAt returns event timestamp.
func (e OrderPaid) OccurredAt() time.Time {
	return e.at
}

// NewOrderCanceled creates OrderCancelled event
// with current timestamp.
func NewOrderCancelled(orderID int) OrderCancelled {
	return OrderCancelled{
		OrderID: orderID,
		at:      time.Now(),
	}
}

// Name returns event type identifier.
func (e OrderCancelled) Name() string {
	return NameOrderCancelled
}

// OccurredAt returns event timestamp.
func (e OrderCancelled) OccurredAt() time.Time {
	return e.at
}

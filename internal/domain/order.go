package domain

import (
	"errors"
	"time"
)

// OrderStatus represents lifecycle of order.
type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusPaid      OrderStatus = "paid"
	OrderStatusCancelled OrderStatus = "cancelled"
)

var (
	ErrInvalidOrderUserID    error = errors.New("invalid order user id")
	ErrOrderEmpty            error = errors.New("order must contain items")
	ErrOrderAlreadyPaid      error = errors.New("order already paid")
	ErrOrderAlreadyCancelled error = errors.New("order already cancelled")
	ErrOrderNotPending       error = errors.New("order is not pending")
)

// Order represents confirmed purchase intent.
//
// Business rules:
//   - Created only with items.
//   - Total is immutable after creation.
//   - Only pending order can be paid or cancelled.
//   - Paid order cannot be cancelled.
//   - Cancelled order cannot be paid.
type Order struct {
	id        int
	userID    int
	items     []OrderItem
	total     int64
	status    OrderStatus
	createdAt time.Time
	paidAt    *time.Time
	version   int

	events []DomainEvent
}

// OrderItem represents snapshot of product at purchase time.
type OrderItem struct {
	variantID int
	quantity  int
	price     int64
}

// NewOrder creates new pending order.
func NewOrder(userID int, items []OrderItem, createdAt time.Time) (*Order, error) {
	if userID <= 0 {
		return nil, ErrInvalidOrderUserID
	}

	if len(items) == 0 {
		return nil, ErrOrderEmpty
	}

	var total int64
	for _, item := range items {
		total += int64(item.quantity) * item.price
	}

	return &Order{
		userID:    userID,
		items:     items,
		total:     total,
		status:    OrderStatusPending,
		createdAt: createdAt,
		version:   1,
	}, nil
}

// ---- GETTERS ----

// ID returns order id.
func (o *Order) ID() int {
	return o.id
}

// Status returns order status.
func (o *Order) Status() OrderStatus {
	return o.status
}

// Total returns order total amount.
func (o *Order) Total() int64 {
	return o.total
}

// Items returns copy of order items.
func (o *Order) Items() []OrderItem {
	result := make([]OrderItem, len(o.items))
	copy(result, o.items)
	return result
}

// Version returns aggregate version.
func (o *Order) Version() int {
	return o.version
}

// MarkPaid marks order as paid.
//
// Fails if:
//   - already paid
//   - already cancelled
//   - not pending
func (o *Order) MarkPaid(now time.Time) error {
	if o.status == OrderStatusPaid {
		return ErrOrderAlreadyPaid
	}

	if o.status == OrderStatusCancelled {
		return ErrOrderAlreadyCancelled
	}

	if o.status != OrderStatusPending {
		return ErrOrderNotPending
	}

	o.status = OrderStatusPaid
	o.paidAt = &now
	o.version++

	o.addEvent(NewOrderPaid(o.id))

	return nil
}

// Cancel cancels order.
//
// Fails if:
//   - already cancelled
//   - already paid
func (o *Order) Cancel() error {
	if o.status == OrderStatusCancelled {
		return ErrOrderAlreadyCancelled
	}

	if o.status == OrderStatusPaid {
		return ErrOrderAlreadyPaid
	}

	if o.status != OrderStatusPending {
		return ErrOrderNotPending
	}

	o.status = OrderStatusCancelled
	o.version++

	o.addEvent(NewOrderCancelled(o.id))

	return nil
}

// PullEvents returns and clears accumulated domain events.
func (o *Order) PullEvents() []DomainEvent {
	result := make([]DomainEvent, len(o.events))
	copy(result, o.events)

	o.events = nil

	return result
}

// addEvent appends domain evet to aggregate buffer.
func (o *Order) addEvent(ev DomainEvent) {
	o.events = append(o.events, ev)
}

// ---- SETTERS ----

// SetID is intended for repository layer only.
func (o *Order) SetID(id int) {
	o.id = id
}

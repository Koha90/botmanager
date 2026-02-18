package domain

// Order represent the Order aggregate root.
//
// It encapsulate business invariants and control all state transitions.
// Order is responsible for generating domain events when its state changes.
type Order struct {
	id         int
	customerID int
	productID  int
	price      int64

	status  OrderStatus
	version int

	events []DomainEvent
}

// NewOrder creates a new Order in "cart" status.
// The returned order is not persisted yet.
// All invariants are enforced inside aggregate methods.
func NewOrder(customerID int, productID int, price int64) *Order {
	o := &Order{
		customerID: customerID,
		productID:  productID,
		price:      price,
		status:     StatusCart,
		version:    1,
	}

	return o
}

// Confirm moves the order from cart to confirmed status.
//
// It validates business rules and generates OrderConfimed event.
// Returns error if transition is not allowed.
func (o *Order) Confirm() error {
	if !o.status.CanConfirm() {
		return ErrInvalidState
	}

	o.status = StatusConfirmed
	o.version++

	o.addEvent(NewOrderConfirmed(o.id))

	return nil
}

// Cancel moves the order from cart to cancelled status.
//
// It validates business rules and generates OrderCancelled event.
// Returns error if transition is not allowed.
func (o *Order) Cancel() error {
	if !o.status.CanCancel() {
		return ErrInvalidState
	}

	o.status = StatusCancelled
	o.version++

	o.addEvent(NewOrderCanceled(o.id))

	return nil
}

// PullEvents returns and clear all accumulated domain events.
//
// Events are typically published by the application layer.
func (o *Order) PullEvents() []DomainEvent {
	events := o.events
	o.events = nil
	return events
}

// ---- GETTERS ----

func (o *Order) ID() int {
	return o.id
}

func (o *Order) CustomerID() int {
	return o.customerID
}

func (o *Order) ProductID() int {
	return o.productID
}

func (o *Order) PriceAtPurchase() int64 {
	return o.price
}

func (o *Order) Status() OrderStatus {
	return o.status
}

func (o *Order) Version() int {
	return o.version
}

func (o *Order) addEvent(event DomainEvent) {
	o.events = append(o.events, event)
}

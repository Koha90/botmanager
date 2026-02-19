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
	if o.status == StatusConfirmed {
		return ErrOrderAlreadyConfirmed
	}
	if !o.status.CanConfirm() {
		return ErrInvalidOrderState
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
	if o.status == StatusCancelled {
		return ErrOrderAlreadyCanceled
	}
	if !o.status.CanCancel() {
		return ErrInvalidOrderState
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

// ID returned id of order.
func (o *Order) ID() int {
	return o.id
}

// CustomerID returned customer's id of order.
func (o *Order) CustomerID() int {
	return o.customerID
}

// ProductID returned product's id of order.
func (o *Order) ProductID() int {
	return o.productID
}

// Price returned price of order.
func (o *Order) Price() int64 {
	return o.price
}

// Status returned status of order.
func (o *Order) Status() OrderStatus {
	return o.status
}

// Version returned version of order.
func (o *Order) Version() int {
	return o.version
}

// ---- SETTERS ----

// SetID setups id of order.
func (o *Order) SetID(id int) {
	o.id = id
}

func (o *Order) addEvent(event DomainEvent) {
	o.events = append(o.events, event)
}

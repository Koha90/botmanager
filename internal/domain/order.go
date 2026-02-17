package domain

import (
	"errors"
)

type Order struct {
	id              int
	customerID      int
	productID       int
	priceAtPurchase int64
	status          OrderStatus
	version         int

	events []DomainEvent
}

func NewOrder(customerID int, productID int, price int64) *Order {
	return &Order{
		customerID:      customerID,
		productID:       productID,
		priceAtPurchase: price,
		status:          StatusCart,
	}
}

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
	return o.priceAtPurchase
}

func (o *Order) Status() OrderStatus {
	return o.status
}

func (o *Order) Confirm() error {
	if !o.status.CanConfirm() {
		return errors.New("only cart order can be confirmed")
	}

	o.status = StatusConfirmed
	o.version++

	return nil
}

func (o *Order) Cancel() error {
	if !o.status.CanCancel() {
		return ErrInvalidState
	}

	o.status = StatusCancelled
	o.version++

	o.addEvent(NewOrderCanceled(o.id))

	return nil
}

func (o *Order) PullEvents() []DomainEvent {
	events := o.events
	o.events = nil
	return events
}

func (o *Order) addEvent(event DomainEvent) {
	o.events = append(o.events, event)
}

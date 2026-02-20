package domain

import "errors"

// OrderStatus represnts order lifecycle state.
//
// It is a value object and guarantees that only valid states exists.
type OrderStatus struct {
	value string
}

var (
	StatusCart      = OrderStatus{"cart"}
	StatusConfirmed = OrderStatus{"confirmed"}
	StatusCancelled = OrderStatus{"canceled"}
)

var allowedStatuses = map[string]OrderStatus{
	StatusCart.value:      StatusCart,
	StatusConfirmed.value: StatusConfirmed,
	StatusCancelled.value: StatusCancelled,
}

// NewOrderStatus validates and creates OrderStatus from string value.
func NewOrderStatus(value string) (OrderStatus, error) {
	if status, ok := allowedStatuses[value]; ok {
		return status, nil
	}
	return OrderStatus{}, errors.New("invalid order status")
}

// CanConfirm returns true if order can be confirmed.
func (s OrderStatus) CanConfirm() bool {
	return s.value == StatusCart.value
}

// CanCancel returns true if order can be cancelled.
func (s OrderStatus) CanCancel() bool {
	return s.value == StatusCart.value
}

func (s OrderStatus) IsFinal() bool {
	return s == StatusCancelled
}

func (s OrderStatus) CanShip() bool {
	return s == StatusConfirmed
}

func (s OrderStatus) String() string {
	return s.value
}

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

func (s OrderStatus) Confirm() (OrderStatus, error) {
	switch s {
	case StatusCart:
		return StatusConfirmed, nil
	case StatusConfirmed:
		return s, ErrOrderAlreadyConfirmed
	case StatusCancelled:
		return s, ErrOrderAlreadyCanceled
	default:
		return s, ErrInvalidOrderState
	}
}

func (s OrderStatus) Cancel() (OrderStatus, error) {
	switch s {
	case StatusCart:
		return StatusCancelled, nil
	case StatusCancelled:
		return s, ErrOrderAlreadyCanceled
	case StatusConfirmed:
		return s, ErrInvalidOrderState
	default:
		return s, ErrInvalidOrderState
	}
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

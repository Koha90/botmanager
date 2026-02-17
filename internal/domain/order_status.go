package domain

import "errors"

type OrderStatus struct {
	value string
}

var (
	StatusCart      = OrderStatus{"cart"}
	StatusConfirmed = OrderStatus{"confirmed"}
	StatusCancelled = OrderStatus{"canceled"}
)

func (s OrderStatus) String() string {
	return s.value
}

func (s OrderStatus) CanConfirm() bool {
	return s.value == StatusCart.value
}

func (s OrderStatus) CanCancel() bool {
	return s.value == StatusCart.value
}

func NewOrderStatus(value string) (OrderStatus, error) {
	switch value {
	case "cart":
		return StatusCart, nil
	case "confirmed":
		return StatusConfirmed, nil
	case "canceled":
		return StatusCancelled, nil
	default:
		return OrderStatus{}, errors.New("invalid order status")
	}
}

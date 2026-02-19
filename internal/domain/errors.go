package domain

import "errors"

var (
	// order errors
	ErrInvalidOrderState     error = errors.New("invalid order state transition")
	ErrOrderNotFound         error = errors.New("order not found")
	ErrOrderAlreadyConfirmed error = errors.New("order already confirmed")
	ErrOrderAlreadyCanceled  error = errors.New("order already canceled")
	ErrOrderUpdate           error = errors.New("order update fail")

	// product errors
	ErrProductNotFound     error = errors.New("product not found")
	ErrInvalidProductName  error = errors.New("invalid product name")
	ErrInvalidProductPrice error = errors.New("invalid product price")
)

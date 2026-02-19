package domain

import "errors"

var (
	// order errors
	ErrInvalidState  error = errors.New("confirmed order cannot be canceled")
	ErrOrderNotFound error = errors.New("order not found")

	// product errors
	ErrProductNotFound     error = errors.New("product not found")
	ErrInvalidProductName  error = errors.New("invalid product name")
	ErrInvalidProductPrice error = errors.New("invalid product price")
)

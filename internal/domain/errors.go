package domain

import "errors"

var (
	// order errors
	ErrInvalidOrderState     error = errors.New("invalid order state transition")
	ErrOrderNotFound         error = errors.New("order not found")
	ErrOrderAlreadyConfirmed error = errors.New("order already confirmed")
	ErrOrderAlreadyCanceled  error = errors.New("order already canceled")
	ErrOrderUpdate           error = errors.New("order update fail")
	ErrOrderPublish          error = errors.New("order publish fail")

	// product errors
	ErrProductNotFound     error = errors.New("product not found")
	ErrInvalidProductName  error = errors.New("invalid product name")
	ErrInvalidProductPrice error = errors.New("invalid product price")
	ErrInvalidCategoryID   error = errors.New("invalid category id")
	ErrInvalidImageURL     error = errors.New("invalid image url")
	ErrInvalidProductID    error = errors.New("invalid product id")
	ErrInvalidPackSize     error = errors.New("invalid pack size")
	ErrInvalidDisctrictID  error = errors.New("invalid district id")
	// tx errors
)

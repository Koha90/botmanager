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
	ErrProductNotFound    error = errors.New("product not found")
	ErrInvalidProductName error = errors.New("invalid product name")
	ErrInvalidCategoryID  error = errors.New("invalid category id")
	ErrInvalidImageURL    error = errors.New("invalid image url")
	ErrInvalidProductID   error = errors.New("invalid product id")

	// product variant errors
	ErrVariantNotFound          error = errors.New("product variant not found")
	ErrInvalidProductPrice      error = errors.New("invalid product price")
	ErrInvalidPackSize          error = errors.New("invalid pack size")
	ErrInvalidDisctrictID       error = errors.New("invalid district id")
	ErrInvalidVariant           error = errors.New("invalid variant")
	ErrInvalidVariantID         error = errors.New("invalid variant id")
	ErrCannotArchiveLastVariant error = errors.New("cannot archive last variant")
	ErrVariantAlreadyExists     error = errors.New("variant already exists")
	// tx errors
)

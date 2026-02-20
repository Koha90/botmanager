package domain

import "strings"

// ProductVariant repesent a variant of product packaging.
type ProductVariant struct {
	id         int
	productID  int
	packSize   string
	districtID int
	price      int64
}

// NewProductVariant creates a new product
// packaging option with price and city district.
// Returns error if price invalid.
func NewProductVariant(
	productID int,
	packSize string,
	districtID int,
	price int64,
) (*ProductVariant, error) {
	if productID <= 0 {
		return nil, ErrInvalidProductID
	}

	if strings.TrimSpace(packSize) == "" {
		return nil, ErrInvalidPackSize
	}

	if districtID <= 0 {
		return nil, ErrInvalidDisctrictID
	}

	if price <= 0 {
		return nil, ErrInvalidProductPrice
	}

	return &ProductVariant{
		productID:  productID,
		packSize:   packSize,
		districtID: districtID,
		price:      price,
	}, nil
}

// ChangePrice changes the price of the product.
// Returns error if price invalid.
func (v *ProductVariant) ChangePrice(price int64) error {
	if price <= 0 {
		return ErrInvalidProductPrice
	}

	v.price = price
	return nil
}

// ChangePackSize changes the packaging of the product.
// Returns error if pack size invalid.
func (v *ProductVariant) ChangePackSize(packSize string) error {
	if strings.TrimSpace(packSize) == "" {
		return ErrInvalidPackSize
	}

	v.packSize = packSize
	return nil
}

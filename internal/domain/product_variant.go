package domain

import (
	"strings"
	"time"
)

// ProductVariant repesent a variant of product packaging.
type ProductVariant struct {
	id         int
	productID  int
	packSize   string
	districtID int
	price      int64
	archivedAt *time.Time
}

// NewProductVariant creates a new product
// packaging option with price and city district.
// Returns error if price invalid.
func NewProductVariant(
	id int,
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
		id:         id,
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

// ByID returns product variant.
func (v *ProductVariant) ID() int {
	return v.id
}

// Price returns price of the variant product.
func (v *ProductVariant) Price() int64 {
	return v.price
}

// ProductID returns identidier of product.
func (v *ProductVariant) ProductID() int {
	return v.productID
}

// Archive settup date of archived variant.
func (v *ProductVariant) Archive(now time.Time) {
	v.archivedAt = &now
}

// IsActive checks if the variant is in the archive.
func (v ProductVariant) IsActive() bool {
	return v.archivedAt == nil
}

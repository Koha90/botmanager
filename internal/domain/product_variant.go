package domain

import (
	"errors"
	"strings"
	"time"
)

var (
	ErrVariantNotFound          error = errors.New("product variant not found")
	ErrInvalidProductPrice      error = errors.New("invalid product price")
	ErrInvalidPackSize          error = errors.New("invalid pack size")
	ErrInvalidDistrictID        error = errors.New("invalid district id")
	ErrInvalidVariant           error = errors.New("invalid variant")
	ErrInvalidVariantID         error = errors.New("invalid variant id")
	ErrCannotArchiveLastVariant error = errors.New("cannot archive last variant")
	ErrVariantAlreadyExists     error = errors.New("variant already exists")
)

// ProductVariant repesent a variant of product packaging.
type ProductVariant struct {
	id         int
	packSize   string
	districtID int
	price      int64

	version    int
	archivedAt *time.Time
}

// NewProductVariant creates a new product
// packaging option with price and city district.
// Returns error if price invalid.
func NewProductVariant(
	packSize string,
	districtID int,
	price int64,
) (*ProductVariant, error) {
	// if productID <= 0 {
	// 	return nil, ErrInvalidProductID
	// }

	if strings.TrimSpace(packSize) == "" {
		return nil, ErrInvalidPackSize
	}

	if districtID <= 0 {
		return nil, ErrInvalidDistrictID
	}

	if price <= 0 {
		return nil, ErrInvalidProductPrice
	}

	return &ProductVariant{
		packSize:   packSize,
		districtID: districtID,
		price:      price,
		version:    1,
	}, nil
}

// NewProductVariantFromDB used only repository.
func NewProductVariantFromDB(
	id int,
	packSize string,
	distrcitID int,
	price int64,
	archivedAt *time.Time,
) ProductVariant {
	return ProductVariant{
		id:         id,
		packSize:   packSize,
		districtID: distrcitID,
		price:      price,
		archivedAt: archivedAt,
	}
}

// ---- SETTERS ----

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

// SetID is used by repository layer only.
func (v *ProductVariant) SetID(id int) {
	v.id = id
}

// ---- GETTERS ----

// ByID returns product variant.
func (v *ProductVariant) ID() int {
	return v.id
}

// Price returns price of the variant product.
func (v *ProductVariant) Price() int64 {
	return v.price
}

// DistrictID returns district id of the variant product.
func (v *ProductVariant) DistrictID() int {
	return v.districtID
}

// ArchivedAt returns time were the variant product was archived.
func (v *ProductVariant) ArchivedAt() *time.Time {
	return v.archivedAt
}

// PackSize returns pack size of the variant product.
func (v *ProductVariant) PackSize() string {
	return v.packSize
}

// Version returns version of the variant product.
func (v *ProductVariant) Version() int {
	return v.version
}

// ---- CHANGERS ----

// Archive settup date of archived variant.
func (v *ProductVariant) Archive(now time.Time) {
	v.archivedAt = &now
}

// IsActive checks if the variant is in the archive.
func (v ProductVariant) IsActive() bool {
	return v.archivedAt == nil
}

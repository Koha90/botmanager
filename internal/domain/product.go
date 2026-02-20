package domain

import (
	"strings"
)

// Product represents the product.
type Product struct {
	id          int
	name        string
	description string
	imagePath   string
	categoryID  int
}

// NewProduct creates a new product.
// The returned product is not persisted yet,
// or error if name empty or invalids
// and category id is invalid.
func NewProduct(
	name string,
	categoryID int,
	description string,
	imagePath string,
) (*Product, error) {
	if strings.TrimSpace(name) == "" {
		return nil, ErrInvalidProductName
	}

	if categoryID <= 0 {
		return nil, ErrInvalidCategoryID
	}

	return &Product{
		name:        name,
		categoryID:  categoryID,
		description: description,
		imagePath:   imagePath,
	}, nil
}

// ---- GETTERS ----

// ID return id of product.
func (p *Product) ID() int {
	return p.id
}

// Name return name of product.
func (p *Product) Name() string {
	return p.name
}

// Description return description of product.
func (p *Product) Description() string {
	return p.description
}

// ImageURL return imageurl of product.
func (p *Product) ImagePath() *string {
	return &p.imagePath
}

// ---- CHANEGERS ----

// Rename renames the product.
func (p *Product) Rename(name string) error {
	if strings.TrimSpace(name) == "" {
		return ErrInvalidProductName
	}

	p.name = name
	return nil
}

// ChangeCategory changes category.
func (p *Product) ChangeCategory(categoryID int) error {
	if categoryID <= 0 {
		return ErrInvalidCategoryID
	}
	p.categoryID = categoryID
	return nil
}

// UpdateDescription updates description the product.
func (p *Product) UpdateDescription(description string) {
	p.description = description
}

// TODO: create other getter-methods.

// ---- SETTERS ----

// SetID sets the product identifier.
// Returns error if identifier is invalid.
func (p *Product) SetID(id int) error {
	if id <= 0 {
		return ErrInvalidProductID
	}

	p.id = id
	return nil
}

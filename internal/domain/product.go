package domain

import (
	"strings"
	"time"
)

// Product represents the product.
type Product struct {
	id          int
	categoryID  int
	name        string
	description string
	imagePath   *string
	variants    []ProductVariant
}

// NewProduct creates a new product.
// The returned product is not persisted yet,
// or error if name empty or invalids
// and category id is invalid.
func NewProduct(
	id int,
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
		id:          id,
		name:        name,
		categoryID:  categoryID,
		description: description,
		imagePath:   &imagePath,
	}, nil
}

// ---- GETTERS ----

// ID return id of product.
func (p *Product) ID() int {
	return p.id
}

// CategoryID returns identifier category of product.
func (p *Product) CategoryID() int {
	return p.categoryID
}

// Name return name of product.
func (p *Product) Name() string {
	return p.name
}

// Description returns description of product.
func (p *Product) Description() string {
	return p.description
}

// ImagePath returns imageurl of product.
func (p *Product) ImagePath() string {
	return *p.imagePath
}

// VariantByID returns variant product of product
// or error if variant was not found.
func (p *Product) VariantByID(id int) (*ProductVariant, error) {
	for i := range p.variants {
		v := &p.variants[i]
		if v.ID() == id && v.IsActive() {
			return v, nil
		}
	}

	return nil, ErrVariantNotFound
}

// ActiveVariants returns of copy the variants if this active.
func (p *Product) ActiveVariants() []ProductVariant {
	var result []ProductVariant
	for _, v := range p.variants {
		if v.IsActive() {
			result = append(result, v)
		}
	}

	return result
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

func (p *Product) AddVariant(
	variantID int,
	packSize string,
	districtID int,
	price int64,
) error {
	if variantID <= 0 {
		return ErrInvalidVariantID
	}

	for _, v := range p.variants {
		if v.ID() == variantID {
			return ErrVariantAlreadyExists
		}
	}

	v, err := NewProductVariant(
		variantID,
		p.id,
		packSize,
		districtID,
		price,
	)
	if err != nil {
		return err
	}

	p.variants = append(p.variants, *v)
	return nil
}

func (p *Product) HasVariants() bool {
	return len(p.variants) > 0
}

// ArchiveVariant ...
func (p *Product) ArchiveVariant(id int, now time.Time) error {
	activeCount := 0
	var target *ProductVariant

	for i := range p.variants {
		if p.variants[i].IsActive() {
			activeCount++
		}
		if p.variants[i].ID() == id && p.variants[i].IsActive() {
			target = &p.variants[i]
		}
	}

	if target == nil {
		return ErrVariantNotFound
	}

	if activeCount <= 1 && target.IsActive() {
		return ErrCannotArchiveLastVariant
	}

	target.Archive(now)
	return nil
}

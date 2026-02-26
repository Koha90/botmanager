package domain

import "strings"

// Category represent category of the product.
type Category struct {
	id   int
	name string
}

// NewCategory create a new category or
// returns an error if name is empty.
func NewCategory(name string) (*Category, error) {
	if strings.TrimSpace(name) == "" {
		return nil, ErrInvalidCategoryName
	}

	return &Category{name: name}, nil
}

// ---- SETTERS ----

// SetID is used by repository layer only.
func (c *Category) SetID(id int) {
	c.id = id
}

// ---- GETTERS ----

// ID returns id of the category.
func (c *Category) ID() int {
	return c.id
}

// Name returns name of the category.
func (c *Category) Name() string {
	return c.name
}

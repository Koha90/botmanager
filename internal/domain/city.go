package domain

import (
	"strings"
)

// City represent the city.
type City struct {
	id   int
	name string
}

// NewCity create a new city.
// Returns error if name is empty.
func NewCity(name string) (*City, error) {
	if strings.TrimSpace(name) == "" {
		return nil, ErrInvalidCityName
	}

	return &City{name: name}, nil
}

// ---- SETTERS ----

// SetID is used by repository layer only.
func (c *City) SetID(id int) {
	c.id = id
}

// ---- GETTERS ----

// ID returns identifier of the city.
func (c *City) ID() int {
	return c.id
}

// Name returns name of the city.
func (c *City) Name() string {
	return c.name
}

package domain

import "strings"

// District represent the district of the city.
type District struct {
	id     int
	cityID int
	name   string
}

// NewDistrict creates a new district of city.
// Returns errror if the city id is wrong or name of
// the district is empty.
func NewDistrict(cityID int, name string) (*District, error) {
	if cityID <= 0 {
		return nil, ErrInvalidCityID
	}

	if strings.TrimSpace(name) == "" {
		return nil, ErrInvalidDistrictName
	}

	return &District{
		cityID: cityID,
		name:   name,
	}, nil
}

// ---- SETTERS ----

// SetID is used by repository layer only.
func (d *District) SetID(id int) {
	d.id = id
}

// ---- GETTERS ----

// ID returns id of the district.
func (d *District) ID() int {
	return d.id
}

// CityID returns id of the city in which
// the district is located.
func (d *District) CityID() int {
	return d.cityID
}

// Name returns name of the district.
func (d *District) Name() string {
	return d.name
}

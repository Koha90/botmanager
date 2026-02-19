package domain

// Product represent the Product
type Product struct {
	id          int
	name        string
	description string
	price       int64
	imageURL    *string
	distictID   int
	categoryID  int
}

// NewProduct creates a new Product.
// The returned product is not persisted yet,
// or error if name or price invalids.
func NewProduct(name string, price int64) (*Product, error) {
	if name == "" {
		return nil, ErrInvalidProductName
	}

	if price <= 0 {
		return nil, ErrInvalidProductPrice
	}
	return &Product{
		name:  name,
		price: price,
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

// Price return price of product.
func (p *Product) Price() int64 {
	return p.price
}

// ImageURL return imageurl of product.
func (p *Product) ImageURL() *string {
	return p.imageURL
}

// TODO: create other getter-methods.

// ---- SETTERS ----

// SetID settup id of product.
func (p *Product) SetID(id int) {
	p.id = id
}

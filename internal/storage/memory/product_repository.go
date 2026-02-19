package memory

import (
	"context"
	"sync"

	"botmanager/internal/domain"
)

type ProductRepository struct {
	mu       *sync.Mutex
	products map[int]*domain.Product
	nextID   int
}

func NewProductRepository(mu *sync.Mutex) *ProductRepository {
	return &ProductRepository{
		mu:       mu,
		products: make(map[int]*domain.Product),
		nextID:   1,
	}
}

func (r *ProductRepository) Create(
	ctx context.Context,
	product *domain.Product,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	product.SetID(r.nextID)
	r.products[r.nextID] = product
	r.nextID++

	return nil
}

func (r *ProductRepository) ByID(
	ctx context.Context,
	id int,
) (*domain.Product, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	product, ok := r.products[id]
	if !ok {
		return nil, domain.ErrProductNotFound
	}

	return product, nil
}

func (r *ProductRepository) Update(ctx context.Context, product *domain.Product) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.products[product.ID()] = product
	return nil
}

func (r *ProductRepository) Delete(
	ctx context.Context,
	id int,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.products, id)
	return nil
}

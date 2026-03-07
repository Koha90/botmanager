package service

import (
	"context"
	"log/slog"

	"botmanager/internal/domain"
)

type ProductService struct {
	repo   ProductRepository
	tx     TxManager
	bus    EventBus
	logger *slog.Logger
}

// NewProductService creates a new ProductService instance.
//
// logger may be nil, in that case slog.Default() is used.
func NewProductService(
	repo ProductRepository,
	tx TxManager,
	bus EventBus,
	logger *slog.Logger,
) *ProductService {
	if repo == nil {
		panic("service: ProductRepository is nil")
	}

	if tx == nil {
		panic("service: TxManager is nil")
	}

	if bus == nil {
		panic("service: EventBus is nil")
	}

	if logger == nil {
		logger = slog.Default()
	}
	return &ProductService{
		repo:   repo,
		tx:     tx,
		bus:    bus,
		logger: logger,
	}
}

func (s *ProductService) Create(
	ctx context.Context,
	name string,
	categoryID int,
	description string,
	imagePath string,
) (*domain.Product, error) {
	p, err := domain.NewProduct(name, categoryID, description, imagePath)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Save(ctx, p); err != nil {
		return nil, err
	}

	return p, nil
}

func (s *ProductService) AddVariant(
	ctx context.Context,
	productID int,
	packSize string,
	districtID int,
	price int64,
) error {
	return s.tx.WithinTransaction(ctx, func(ctx context.Context) error {
		p, err := s.repo.ByID(ctx, productID)
		if err != nil {
			s.logger.Error("failed to load product", "id", productID, "err", err)
			return err
		}

		if err := p.AddVariant(packSize, districtID, price); err != nil {
			s.logger.Warn("failed to add variant", "productID", productID, "err", err)
			return err
		}

		if err := s.repo.Save(ctx, p); err != nil {
			s.logger.Error("failed to save product", "id", productID, "err", err)
			return err
		}

		s.logger.Info(
			"variant added",
			"productID",
			productID,
			"packSize",
			packSize,
			"districtID",
			districtID,
		)
		return nil
	})
}

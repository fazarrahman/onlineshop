package service

import (
	"context"

	"github.com/fazarrahman/onlineshop/domain/product/entity"
)

// GetAllProducts ...
// Get all active products
func (s *Svc) GetAllProducts(ctx context.Context) ([]*entity.Product, error) {
	products, err := s.productRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

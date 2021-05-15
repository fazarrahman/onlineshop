package repository

import (
	"context"

	"github.com/fazarrahman/onlineshop/domain/product/entity"
)

// Repository ...
type Repository interface {
	GetAll(ctx context.Context) ([]*entity.Product, error)
}

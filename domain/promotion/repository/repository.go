package repository

import (
	"context"

	"github.com/fazarrahman/onlineshop/domain/promotion/entity"
)

// Repository ...
type Repository interface {
	GetAll(ctx context.Context) ([]*entity.Promotion, error)
}

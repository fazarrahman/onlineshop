package service

import (
	"context"

	"github.com/fazarrahman/onlineshop/domain/product/entity"
	product_repository "github.com/fazarrahman/onlineshop/domain/product/repository"
	promotion_repository "github.com/fazarrahman/onlineshop/domain/promotion/repository"
)

// Svc ...
type Svc struct {
	promotionRepository promotion_repository.Repository
	productRepository   product_repository.Repository
}

// New ...
func New(promotionRepository promotion_repository.Repository, productRepository product_repository.Repository) *Svc {
	return &Svc{promotionRepository: promotionRepository, productRepository: productRepository}
}

// Service ...
type Service interface {
	CartCheckout(ctx context.Context, skuList []string) (*float64, error)
	GetAllProducts(ctx context.Context) ([]*entity.Product, error)
}

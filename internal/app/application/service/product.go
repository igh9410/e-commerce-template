package service

import (
	"context"

	"github.com/igh9410/e-commerce-template/internal/app/domain/product"
)

type ProductService interface {
	CreateProduct(ctx context.Context) error
}

type productService struct {
	productRepo product.Repository
}

func (s *productService) CreateProduct(ctx context.Context) error {
	return nil
}

func NewProductService(productRepo product.Repository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

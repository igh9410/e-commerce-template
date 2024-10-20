package service

import (
	"context"

	"github.com/igh9410/e-commerce-template/internal/app/domain/product"
)

type ProductService interface {
	CreateProduct(ctx context.Context, params product.CreateProductParams) (*product.Product, error)
}

type productService struct {
	productRepo product.Repository
}

func (s *productService) CreateProduct(ctx context.Context, params product.CreateProductParams) (*product.Product, error) {

	product, err := product.NewProduct(params)
	if err != nil {
		return nil, err
	}
	return s.productRepo.CreateProduct(ctx, product)
}

func NewProductService(productRepo product.Repository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

package service

import "context"

type ProductService interface {
	CreateProduct(ctx context.Context) error
}

type productService struct{}

func (s *productService) CreateProduct(ctx context.Context) error {
	return nil
}

func NewProductService() ProductService {
	return &productService{}
}

package server

import (
	"context"

	"github.com/igh9410/e-commerce-template/internal/app/domain/product"
	pb "github.com/igh9410/e-commerce-template/internal/pb/v1"
	"github.com/igh9410/e-commerce-template/pkg/utils"
)

func (a *API) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	params := product.CreateProductParams{
		Name:        req.Name,
		Description: utils.ToStringPointer(req.Description),
		Category:    utils.ToStringPointer(req.Category),
		Price:       req.Price,
		Status:      req.Status.String(),
		Tags:        req.Tags,
	}

	createdProduct, err := a.productService.CreateProduct(ctx, params)
	if err != nil {
		return nil, err
	}

	return &pb.CreateProductResponse{
		Product: &pb.Product{
			Id:          createdProduct.ID,
			Name:        createdProduct.Name,
			Description: *createdProduct.Description,
			Category:    *createdProduct.Category,
			Price:       createdProduct.Price,
			Status:      pb.ProductStatus(pb.ProductStatus_value[createdProduct.Status]),
			Tags:        createdProduct.Tags,
		},
	}, nil
}

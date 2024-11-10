package server

import (
	"github.com/igh9410/e-commerce-template/internal/app/application/service"
	pb "github.com/igh9410/e-commerce-template/internal/gen/v1"
)

// API implements pb.ProductServiceServer
type API struct {
	pb.UnimplementedProductServiceServer
	productService service.ProductService
}

func NewAPI(productService service.ProductService) pb.ProductServiceServer {
	return &API{
		productService: productService,
	}
}

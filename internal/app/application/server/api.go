package server

import (
	"github.com/igh9410/e-commerce-template/internal/api"
	"github.com/igh9410/e-commerce-template/internal/app/application/service"
)

// API implements the api.StrictServerInterface
var _ api.StrictServerInterface = (*API)(nil)

type API struct {
	productService service.ProductService
}

func NewAPI(productService service.ProductService) api.StrictServerInterface {
	return &API{
		productService: productService,
	}
}

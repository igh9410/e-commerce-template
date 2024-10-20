package server

import (
	"context"

	"github.com/igh9410/e-commerce-template/internal/api"
	"github.com/igh9410/e-commerce-template/internal/app/domain/product"
	"github.com/igh9410/e-commerce-template/internal/utils"
)

// ProductServiceCreateProduct implements api.StrictServerInterface.
func (a *API) ProductServiceCreateProduct(ctx context.Context, request api.ProductServiceCreateProductRequestObject) (api.ProductServiceCreateProductResponseObject, error) {
	// Handle nullable fields from the request
	var description *string
	if request.Body.Description != nil {
		description = request.Body.Description
	}

	var category *string
	if request.Body.Category != nil {
		category = request.Body.Category
	}

	// Map the request data to CreateProductParams
	params := product.CreateProductParams{
		Name:        request.Body.Name,
		Description: description,
		Category:    category,
		Price:       utils.StringToInt64(request.Body.Price),
		Status:      string(*request.Body.Status),
		Tags:        *request.Body.Tags,
	}

	// Call the service
	createdProduct, err := a.productService.CreateProduct(ctx, params)
	if err != nil {
		return api.ProductServiceCreateProductdefaultJSONResponse{
			Body: api.Status{
				Message: err.Error(),
			},
			StatusCode: 400,
		}, err
	}
}

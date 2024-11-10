package server

import (
	"context"

	"github.com/igh9410/e-commerce-template/internal/app/domain/product"
	pb "github.com/igh9410/e-commerce-template/internal/gen/v1"
	"github.com/igh9410/e-commerce-template/internal/utils"
)

/*
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

		// Handle the status field safely
		var status string
		if request.Body.Status != nil {
			status = string(*request.Body.Status) // Convert enum to string if not nil
		}

		// Handle tags safely
		var tags []string
		if request.Body.Tags != nil {
			tags = *request.Body.Tags
		} else {
			tags = []string{} // Default to an empty slice if nil
		}

		// Map the request data to CreateProductParams
		params := product.CreateProductParams{
			Name:        request.Body.Name,
			Description: description,
			Category:    category,
			Price:       utils.StringToInt64(request.Body.Price),
			Status:      status,
			Tags:        tags,
		}

		// Call the service
		createdProduct, err := a.productService.CreateProduct(ctx, params)
		if err != nil {
			return api.ProductServiceCreateProductdefaultJSONResponse{
				Body: api.Status{
					Message: utils.ToStringPointer(err.Error()),
				},
				StatusCode: 400,
			}, err
		}

		// Convert internal status to API status using mapProductStatus helper
		productStatus := mapProductStatus(createdProduct.Status)

		// Map the response data
		return api.ProductServiceCreateProduct200JSONResponse{
			Product: &api.Product{
				Category:    createdProduct.Category,
				Description: createdProduct.Description,
				Id:          utils.ToStringPointer(createdProduct.ID),
				Name:        utils.ToStringPointer(createdProduct.Name),
				Price:       utils.Int64ToStringPointer(createdProduct.Price),
				Status:      &productStatus,
				Tags:        &tags,
			},
		}, nil

}

// Helper function to map internal status to api.ProductStatus

	func mapProductStatus(internalStatus string) api.ProductStatus {
		switch internalStatus {
		case "ACTIVE":
			return api.ProductStatusACTIVE
		case "INACTIVE":
			return api.ProductStatusINACTIVE
		default:
			// Handle unknown status
			return api.ProductStatusINACTIVE // or a default value
		}
	}
*/
func (a *API) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	params := product.CreateProductParams{
		Name:        req.Name,
		Description: utils.ToStringPointer(req.Description),
		Category:    utils.ToStringPointer(req.Category),
		Price:       req.Price,
		Status:      string(req.Status),
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

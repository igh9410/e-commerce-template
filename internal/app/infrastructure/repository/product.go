package repository

import (
	"context"

	"github.com/igh9410/e-commerce-template/internal/app/domain/product"
	"github.com/igh9410/e-commerce-template/internal/app/infrastructure/postgres"
	"github.com/igh9410/e-commerce-template/internal/app/infrastructure/postgres/sqlc"
)

var _ product.Repository = (*repository)(nil)

type repository struct {
	db *postgres.Database
}

// CreateProduct implements product.Repository.
func (r *repository) CreateProduct(ctx context.Context, product *product.Product) (*product.Product, error) {
	// Convert entity to SQLC model
	dbProduct := ProductEntityToModel(product)

	arg := sqlc.CreateProductParams{
		Name:        dbProduct.Name,
		Description: dbProduct.Description,
		Category:    dbProduct.Category,
		Price:       dbProduct.Price,
		Status:      dbProduct.Status,
		Tags:        dbProduct.Tags,
	}

	// Insert product into the database
	createdProduct, err := r.db.Querier.CreateProduct(ctx, arg)
	if err != nil {
		return nil, err
	}
	return ProductModelToEntity(createdProduct), nil
}

func NewProductRepository(db *postgres.Database) product.Repository {
	return &repository{db: db}
}

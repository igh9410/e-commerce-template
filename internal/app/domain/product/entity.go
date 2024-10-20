package product

import (
	"errors"
	"time"
)

type Product struct {
	ID          string
	Name        string
	Description *string
	Category    *string
	Price       int64
	Status      string
	Tags        []string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

// NewProduct is a constructor for the Product entity. Accepts CreateProductParams as input.
func NewProduct(params CreateProductParams) (*Product, error) {
	// Validation logic for required fields
	if params.Name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if params.Price < 0 {
		return nil, errors.New("price cannot be negative")
	}

	now := time.Now()

	// Return a new product instance without the ID (which will be generated later)
	return &Product{
		Name:        params.Name,
		Description: params.Description, // Could be nil if not provided
		Category:    params.Category,    // Could be nil if not provided
		Price:       params.Price,
		Status:      params.Status,
		Tags:        params.Tags, // No pointer needed, defaults to empty slice if empty
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

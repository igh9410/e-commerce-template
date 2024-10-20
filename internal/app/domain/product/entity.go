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

// NewProduct is a constructor for the Product entity. Since ID is autogenerated, it is left empty.
func NewProduct(name string, price int64, status string, tags []string, description, category *string) (*Product, error) {
	// Validation logic
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}
	if price < 0 {
		return nil, errors.New("price cannot be negative")
	}

	now := time.Now()

	// Return a new product instance without the ID (which will be generated later)
	return &Product{
		Name:        name,
		Description: description,
		Category:    category,
		Price:       price,
		Status:      status,
		Tags:        tags,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

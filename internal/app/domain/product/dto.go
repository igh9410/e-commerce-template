package product

// CreateProductParams represents the DTO for creating a new product.
type CreateProductParams struct {
	Name        string   `json:"name"`                  // Required field
	Description *string  `json:"description,omitempty"` // Optional field, can be null
	Price       int64    `json:"price"`                 // Required field
	Category    *string  `json:"category,omitempty"`    // Optional field, can be null
	Status      string   `json:"status"`                // Required field
	Tags        []string `json:"tags,omitempty"`        // Optional but defaults to empty slice, no pointer needed
}

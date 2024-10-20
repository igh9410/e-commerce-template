package repository

import (
	"github.com/google/uuid"
	"github.com/igh9410/e-commerce-template/internal/app/domain/product"
	"github.com/igh9410/e-commerce-template/internal/app/infrastructure/postgres/sqlc"
	"github.com/igh9410/e-commerce-template/internal/utils"
)

// EntityToSQLCProduct converts a domain Product to an SQLC Product model.
func ProductEntityToModel(p *product.Product) sqlc.Product {
	return sqlc.Product{
		ID:          uuid.MustParse(p.ID), // Convert string to UUID
		Name:        p.Name,
		Description: utils.StringToPgtypeOrNull(p.Description),
		Category:    utils.StringToPgtypeOrNull(p.Category),
		Price:       p.Price,
		Status:      utils.StringToPgtype(p.Status),
		Tags:        p.Tags,
		CreatedAt:   utils.TimeToPgtypeTimestamptz(p.CreatedAt),
		UpdatedAt:   utils.TimeToPgtypeTimestamptz(p.UpdatedAt),
		DeletedAt:   utils.TimeToPgtypeTimestamptzOrNull(p.DeletedAt),
	}
}

// SQLCProductToEntity converts an SQLC Product to a domain Product model.
func ProductModelToEntity(p sqlc.Product) *product.Product {
	return &product.Product{
		ID:          p.ID.String(), // Convert UUID to string
		Name:        p.Name,
		Description: utils.PgtypeToStringOrNull(p.Description),
		Category:    utils.PgtypeToStringOrNull(p.Category),
		Price:       p.Price,
		Status:      utils.PgtypeToString(p.Status),
		Tags:        p.Tags,
		CreatedAt:   utils.PgtypeTimestamptzToTime(p.CreatedAt),
		UpdatedAt:   utils.PgtypeTimestamptzToTime(p.UpdatedAt),
		DeletedAt:   utils.PgtypeTimestamptzToTimeOrNull(p.DeletedAt),
	}
}

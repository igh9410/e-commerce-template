package postgres

import (
	"context"
	"fmt"

	"os"

	"github.com/igh9410/e-commerce-template/internal/app/infrastructure/postgres/sqlc"
	"github.com/igh9410/e-commerce-template/internal/pkg/logger"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool    *pgxpool.Pool
	Querier sqlc.Querier
}

func NewDatabase() (*Database, error) {

	sugar := logger.GetSugaredLogger()
	// Create the connection string using the retrieved values
	// Retrieve the values from environment variables
	var username string
	var password string
	var host string

	username = os.Getenv("POSTGRES_USERNAME")
	password = os.Getenv("POSTGRES_PASSWORD")
	host = os.Getenv("POSTGRES_HOST") // Running in local Docker container

	if host == "" { // Running in local environment
		host = "localhost"
	}
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:5432/postgres?sslmode=disable", username, password, host)

	// Connect to the database
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		sugar.Fatalf("Unable to connect to database: %v", err)
		return nil, err
	}

	// Ping the database to ensure the connection is established
	if err := pool.Ping(ctx); err != nil {
		sugar.Fatalf("Unable to ping the database: %v", err)
		return nil, err
	}

	sugar.Info("Database initialized")

	querier := sqlc.New(pool)

	return &Database{Pool: pool, Querier: querier}, nil
}

// Close closes the database connection pool.
func (d *Database) Close() error {
	if d.Pool != nil {
		d.Pool.Close()
	}
	return nil
}

// GetDB returns the underlying *pgxpool.Pool instance for advanced database operations.
func (d *Database) GetDB() *pgxpool.Pool {
	return d.Pool
}

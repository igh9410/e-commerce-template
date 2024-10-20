include .env
PROTO_PATH="./proto"
DB_MIGRATION_PATH="./internal/app/infrastructure/database/migrations"

# Makefile`
.PHONY: all run docker-push docker-run linter create-migration goose-version migrate-up migrate-down test

all: generate-docs generate-server 
# Run the application
run:
	@echo "Running the application..."
	@go run cmd/e-commerce-api/main.go

# Build the application using Docker
docker-push:
#	@echo "Building the application using Docker..."
	@./scripts/docker-push.sh

# Run golangci-lint
linter:
	@golangci-lint run

# Create database migration file
create-migration:
	@goose -dir ${DB_MIGRATION_PATH} create $(filter-out $@,$(MAKECMDGOALS)) sql
	@echo "Migration created."

goose-version:
	@goose -dir ${DB_MIGRATION_PATH} postgres "host=localhost user=${POSTGRES_USERNAME} password=${POSTGRES_PASSWORD} dbname=postgres sslmode=disable port=5432" version

# Run the database migrations
migrate-up:
	@echo "Running the database migrations..."
	@goose -dir ${DB_MIGRATION_PATH} postgres "host=localhost user=${POSTGRES_USERNAME} password=${POSTGRES_PASSWORD} dbname=postgres sslmode=disable port=5432" up

# Rollback the database migrations
migrate-down:
	@echo "Rolling back the database migrations..."
	@goose -dir ${DB_MIGRATION_PATH} postgres "host=localhost user=${POSTGRES_USERNAME} password=${POSTGRES_PASSWORD} dbname=postgres sslmode=disable port=5432" down

# Run the tests
test:
	@echo "Running the tests..."
	@go test ./... -v -cover -coverprofile=coverage.out

# Generate Swagger documentation from proto files
generate-docs:
	@echo "Generating OpenAPI documentation..."
	@protoc --proto_path=${PROTO_PATH} -I=proto/third_party --openapi_out=internal/api --openapi_opt=enum_type=string ${PROTO_PATH}/*.proto
	
# Generate server code from OpenAPI specification
generate-server:
	@echo "Generating server code from OpenAPI specification..."
	@oapi-codegen --generate=gin-server,strict-server,embedded-spec --package=api -o internal/api/server.gen.go internal/api/openapi.yaml
	@oapi-codegen --generate=models --package=api -o internal/api/types.gen.go  internal/api/openapi.yaml
	@go fmt github.com/igh9410/e-commerce-api/backend/internal/api

# Generate client code from OpenAPI specification
generate-client:
		npx openapi-typescript-codegen -i internal/api/openapi.yaml -o dist/e-commerce-api-client --name ECommerceClient --client fetch 
		


# Generate SQLC code 
sqlc-generate:
	@echo "Generating SQLC code..."
	@sqlc generate -f internal/pkg/sqlc/sqlc.yaml

	
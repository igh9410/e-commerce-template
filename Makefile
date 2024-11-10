include .env
PROTO_PATH="./proto"
DB_MIGRATION_PATH="./internal/app/infrastructure/postgres/migrations"

# Makefile`
.PHONY: all run docker-push docker-run linter create-migration goose-version migrate-up migrate-down test

all: generate-docs generate-server generate-client sqlc-generate

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
	
generate-server:
	@echo "Generating gRPC and gRPC-Gateway code..."
	@protoc --proto_path=${PROTO_PATH} -I=proto/third_party \
		--go_out=. \
		--go-grpc_out=. \
		--grpc-gateway_out=. \
		--grpc-gateway_opt=logtostderr=true \
		${PROTO_PATH}/*.proto

# Generate client code from OpenAPI specification
generate-client:
	@echo "Generating client code from OpenAPI specification..."
	@npx openapi-typescript-codegen -i internal/api/openapi.yaml -o dist/e-commerce-api-client --name ECommerceClient --client fetch 
		


# Generate SQLC code 
sqlc-generate:
	@echo "Generating SQLC code..."
	@sqlc generate -f internal/app/infrastructure/postgres/sqlc.yaml



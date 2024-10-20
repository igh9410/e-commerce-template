# Build the application from source
FROM golang:1.22 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

# Copy the entire application
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -o e-commerce-api ./cmd/e-commerce-api/main.go


# Deploy the application binary into a lean image
FROM gcr.io/distroless/cc-debian12 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/e-commerce-api /e-commerce-api

EXPOSE 8080

USER nonroot:nonroot


# Specify the command to run on container start
ENTRYPOINT ["/e-commerce-api"]
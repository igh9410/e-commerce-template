package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/igh9410/e-commerce-template/internal/api/middleware"
	server "github.com/igh9410/e-commerce-template/internal/app/application/server"
	"github.com/igh9410/e-commerce-template/internal/app/application/service"
	"github.com/igh9410/e-commerce-template/internal/app/infrastructure/postgres"
	repo "github.com/igh9410/e-commerce-template/internal/app/infrastructure/repository"
	"github.com/igh9410/e-commerce-template/internal/docs"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	// Create a new zap logger
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	sugar := logger.Sugar()

	// Load environment variables
	if err := godotenv.Load(".env"); err != nil {
		sugar.Info("No .env file found. Using OS environment variables.")
	}

	// Initialize database connection
	dbConn, err := postgres.NewDatabase()
	if err != nil {
		sugar.Fatalf("Could not initialize database connection: %s", err)
	}

	productRepo := repo.NewProductRepository(dbConn)

	productService := service.NewProductService(productRepo)

	api := server.NewAPI(productService)
	grpcServer := grpc.NewServer()

	// Start gRPC server in a separate goroutine
	go func() {
		listener, err := net.Listen("tcp", ":50051")
		if err != nil {
			sugar.Fatalf("Failed to listen: %v", err)
		}
		sugar.Info("Starting gRPC server on :50051")
		if err := grpcServer.Serve(listener); err != nil {
			sugar.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	mux := runtime.NewServeMux()
	// Register all services with one call
	if err := server.RegisterAllServices(grpcServer, mux, api); err != nil {
		panic(err)
	}

	// Set up Gin for serving Swagger UI and additional routes
	r := gin.Default()
	// Use GinZapLogger middleware with zap logger
	r.Use(middleware.GinZapLogger(logger))

	// Use Swagger with the combined function
	docs.UseSwagger(r)

	// Register gRPC-Gateway as a route in Gin
	r.Any("/api/v1/*any", gin.WrapH(mux))

	// Add a home endpoint
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	// Start HTTP server (Gin)
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		sugar.Info("Starting HTTP server on :8080")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			sugar.Fatalf("Failed to serve HTTP: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 3 seconds.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	sugar.Info("Shutting down server...")

	ctxShutdown, cancelShutdown := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelShutdown()

	if err := httpServer.Shutdown(ctxShutdown); err != nil {
		sugar.Fatal("HTTP Server Shutdown:", err)
	}

	grpcServer.GracefulStop()
	sugar.Info("Server exiting")
}

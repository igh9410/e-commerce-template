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
	"github.com/igh9410/e-commerce-template/internal/api"
	server "github.com/igh9410/e-commerce-template/internal/app/application/server"
	"github.com/igh9410/e-commerce-template/internal/app/application/service"
	"github.com/igh9410/e-commerce-template/internal/app/infrastructure/postgres"
	repo "github.com/igh9410/e-commerce-template/internal/app/infrastructure/repository"
	"github.com/igh9410/e-commerce-template/internal/docs"
	pb "github.com/igh9410/e-commerce-template/internal/gen/v1"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	grpcServer := grpc.NewServer()

	// Register gRPC server
	pb.RegisterProductServiceServer(grpcServer, server.NewAPI(productService))

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

	// Set up gRPC-Gateway
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// Register gRPC-Gateway endpoints
	err = pb.RegisterProductServiceHandlerFromEndpoint(ctx, mux, ":50051", opts)
	if err != nil {
		sugar.Fatalf("Failed to register gRPC-Gateway: %v", err)
	}

	// Set up Gin for serving Swagger UI and additional routes
	r := gin.Default()

	// Serve the Swagger UI files
	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}
	swagger.Servers = nil
	docs.UseSwagger(r, swagger)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	// Register gRPC-Gateway as a route in Gin
	r.Any("/api/v1/*any", gin.WrapH(mux))

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

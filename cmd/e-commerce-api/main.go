package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/igh9410/e-commerce-template/internal/api"
	"github.com/igh9410/e-commerce-template/internal/api/middleware"
	server "github.com/igh9410/e-commerce-template/internal/app/application/server"
	"github.com/igh9410/e-commerce-template/internal/app/application/service"
	db "github.com/igh9410/e-commerce-template/internal/app/infrastructure/postgres"
	repo "github.com/igh9410/e-commerce-template/internal/app/infrastructure/repository"
	"github.com/igh9410/e-commerce-template/internal/docs"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil { // Running in local, must be run on go run . in ./cmd directory
		slog.Info("No .env file found. Using OS environment variables.")
	}

	dbConn, err := db.NewDatabase()

	if err != nil {
		log.Fatalf("Could not initialize database connection: %s", err)
	}

	r := gin.Default()

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000", "http://127.0.0.1:5500"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	swagger, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}

	// Allow all origins for swagger UI
	swagger.Servers = nil

	// Serve the Swagger UI files
	docs.UseSwagger(r, swagger)

	r.GET("/", func(c *gin.Context) {
		//time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	// This route is always accessible.
	r.GET("/api/public", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from a public endpoint! You don't need to be authenticated to see this."})
	})

	// This route is only accessible if the user has a valid access_token.
	r.GET("/api/private", middleware.EnsureValidToken(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello from a private endpoint! You need to be authenticated to see this."})
	})

	productRepo := repo.NewProductRepository(dbConn)

	productService := service.NewProductService(productRepo)

	// Create an instance of your handler that implements api.ServerInterface
	handler := api.NewStrictHandler(server.NewAPI(productService), nil)

	// Register the handlers with Gin
	api.RegisterHandlers(r, handler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Println("Server listening on http://localhost:8080")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 3 seconds.

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	<-ctx.Done()
	log.Println("Server exiting")

}

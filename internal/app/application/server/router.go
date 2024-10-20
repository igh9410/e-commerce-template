package server

/*
import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/igh9410/e-commerce-template/internal/api"
	"github.com/igh9410/e-commerce-template/internal/api/middleware"
	"github.com/igh9410/e-commerce-template/internal/docs"
)

// swagger embed files
// gin-swagger middleware

var r *gin.Engine

type RouterConfig struct {
}

func InitRouter(cfg *RouterConfig) *gin.Engine {
	r = gin.Default()

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

	// Create an instance of your handler that implements api.ServerInterface
	handler := api.NewStrictHandler(NewAPI(), nil)

	// Register the handlers with Gin
	api.RegisterHandlers(r, handler)

	return r

} */

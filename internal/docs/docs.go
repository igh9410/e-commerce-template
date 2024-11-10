package docs

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/gin-gonic/gin"
)

func byteHandler(b []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Write(b)
	}
}

func HandleSpec(swagger *openapi3.T) http.HandlerFunc {
	b, err := swagger.MarshalJSON()
	if err != nil {
		panic(err)
	}
	return byteHandler(b)
}

// loadAndServeSpec loads the OpenAPI specification from the given path and returns an http.HandlerFunc
func loadAndServeSpec(swaggerPath string) http.HandlerFunc {
	loader := openapi3.NewLoader()
	swagger, err := loader.LoadFromFile(swaggerPath)
	if err != nil {
		log.Fatalf("Failed to load OpenAPI spec: %v", err)
	}

	b, err := swagger.MarshalJSON()
	if err != nil {
		log.Fatalf("Failed to marshal OpenAPI spec to JSON: %v", err)
	}
	return func(w http.ResponseWriter, _ *http.Request) {
		w.Write(b)
	}
}

// UseSwagger combined function for Gin
func UseSwagger(r *gin.Engine) {
	// Define the path where openapi.yaml is located
	swaggerPath := filepath.Join("internal", "api", "openapi.yaml")

	// Serve the Swagger JSON
	r.GET("/swagger.json", gin.WrapH(loadAndServeSpec(swaggerPath)))

	// Serve the Swagger UI files
	r.Static("/docs", "docs/swagger-ui") // Adjust the path as per your folder structure
}

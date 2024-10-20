package middleware

import (
	"bytes"
	"io"

	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GinZapLogger is a middleware for structured logging using zap
func GinZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Capture request body (we need to read it here and replace it, since c.Request.Body is a ReadCloser)
		var requestBody string
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil {
				requestBody = string(bodyBytes)
				// Re-assign the request body as we've consumed it
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			}
		}

		// Process the request
		c.Next()

		// Log details after processing the request, without headers and response size
		logger.Info("Incoming request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("ip", c.ClientIP()),                     // Capture client IP
			zap.Int("status", c.Writer.Status()),               // Response status code
			zap.Duration("latency", time.Since(start)),         // Latency
			zap.String("user_agent", c.Request.UserAgent()),    // User agent
			zap.String("request_body", requestBody),            // Captured request body
			zap.String("query_params", c.Request.URL.RawQuery), // Query parameters
		)
	}
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

// CORSMiddleware handles CORS requests with restricted origins based on the environment
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the environment variable for allowed origins
		allowedOrigins := os.Getenv("ALLOWED_ORIGINS")

		// If there are multiple allowed origins, split them by comma
		origins := strings.Split(allowedOrigins, ",")

		// Get the request origin
		origin := c.Request.Header.Get("Origin")

		// Check if the request origin is allowed
		isAllowed := false
		for _, allowedOrigin := range origins {
			if allowedOrigin == origin {
				isAllowed = true
				break
			}
		}

		// If the origin is allowed, set the CORS headers
		if isAllowed {
			c.Header("Access-Control-Allow-Origin", origin)
		} else {
			c.Header("Access-Control-Allow-Origin", "") // Disallow if not in the allowed list
		}

		// Other CORS headers
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Header("Access-Control-Allow-Credentials", "true")

		// If the method is OPTIONS, respond with 200 OK (preflight request)
		if c.Request.Method == "OPTIONS" {
			c.JSON(http.StatusOK, gin.H{"message": "OK"})
			return
		}

		// Proceed with the request
		c.Next()
	}
}

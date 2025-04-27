package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Simulate private data
var secrets = gin.H{
	"token1": gin.H{"user": "foo", "email": "foo@bar.com", "phone": "123433"},
	"token2": gin.H{"user": "austin", "email": "austin@example.com", "phone": "666"},
	"token3": gin.H{"user": "lena", "email": "lena@guapa.com", "phone": "523443"},
}

// Middleware for Bearer Token Authentication
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		// Extract the token (remove "Bearer " prefix)
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		if secret, ok := secrets[token]; ok {
			// Token is valid, store user info in context
			c.Set("secret", secret)
			fmt.Println(token)
		} else {
			// Invalid token
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()

	// Apply CORS middleware with a custom configuration
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500"},                   // Whitelisted domains
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // Allowed HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Allowed headers
		ExposeHeaders:    []string{"Content-Length"},                          // Headers exposed to the client
		AllowCredentials: true,                                                // Allow cookies/auth headers
		MaxAge:           12 * time.Hour,                                      // Preflight request cache duration
	}))
	/*
	  router.Use(cors.Default())
	*/

	// Static
	r.Static("/static", "./static")

	// Example routes
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Protect /admin with TokenAuthMiddleware
	authorized := r.Group("/admin", TokenAuthMiddleware())

	// Endpoint to fetch secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// Retrieve user info from context
		if secret, exists := c.Get("secret"); exists {
			c.JSON(http.StatusOK, gin.H{"secret": secret})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No secret found"})
		}
	})

	// Listen and serve
	r.Run(":8080")
}

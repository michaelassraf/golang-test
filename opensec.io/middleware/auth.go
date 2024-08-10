package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"opensec.io/main/opensec.io/auth"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}
		tokenString := authHeader[7:] // Remove "Bearer " prefix
		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}
		c.Set("user", claims) // Attach user information to the context
		c.Next()
	}
}

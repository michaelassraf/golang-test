package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"opensec.io/main/opensec.io/auth"
)

func RegisterUser(c *gin.Context) {
	// ... (Handle registration logic)
}
func LoginUser(c *gin.Context) {
	// ... (Handle login logic)
	// Generate JWT token
	token, err := auth.GenerateToken(user.Username, user.Role)
	if err != nil {
		// ... (Handle error)
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}
func GetUser(c *gin.Context) {
	// Get user information from context
	claims, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	userClaims := claims.(*auth.Claims) // Type assertion
	// ... (Use userClaims.Username and userClaims.Role to access data)
}

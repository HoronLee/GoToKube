package auth

import (
	"GoToKube/database"
	"GoToKube/web/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			c.Abort()
			return
		}

		// Check if the Authorization header starts with "Bearer "
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token prefix"})
			c.Abort()
			return
		}

		// Extract the actual JWT token by removing the "Bearer " prefix
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := ParseJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		db, err := database.GetDBConnection()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
			c.Abort()
			return
		}

		var user models.User
		if err := db.First(&user, claims.UserID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		// Set the user in the context for further handlers to use
		c.Set("user", user)
		c.Next()
	}
}

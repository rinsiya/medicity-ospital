package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		userType := c.GetHeader("User-Type")
		if userType != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func AdminDashboard() gin.HandlerFunc {
	return func(c *gin.Context) {
		userType := c.GetHeader("User-Type")
		if userType != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
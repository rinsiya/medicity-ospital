package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")
		session := sessions.Default(c)
		role, ok := session.Get("role").(string)

		if !ok || role != "admin" {
			c.String(http.StatusForbidden, "Forbidden")
			c.Abort()
			return
		}

		c.Next()
	}
}

func CheckSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		role, _ := session.Get("role").(string)

		if role == "admin" {
			c.Redirect(http.StatusFound, "/admin/dashboard")
			c.Abort()
			return
		}

		if role == "user" {
			c.Redirect(http.StatusFound, "/home")
			c.Abort()
			return
		}

		c.Next()
	}
}

package handlers

import "github.com/gin-gonic/gin"



func (h *UserHandler) AdminLogin(c *gin.Context) {
	c.HTML(200, "adminLogin.html", nil)
}

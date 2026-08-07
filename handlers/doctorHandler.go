package handlers

import "github.com/gin-gonic/gin"


func (h *UserHandler) DoctorSignup(c *gin.Context) {
	c.HTML(200, "doctorRegister.html", nil)
}

func (h *UserHandler) DoctorLogin(c *gin.Context) {
	c.HTML(200, "doctorLogin.html", nil)
}

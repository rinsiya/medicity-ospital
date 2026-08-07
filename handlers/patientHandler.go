package handlers

import (
	"medicity/logger"
	"medicity/services"
	"medicity/dto"
	"medicity/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	Service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}


func (h *UserHandler) Home(c *gin.Context) {
	c.HTML(200, "home.html", nil)
}
func (h *UserHandler) PatientSignup(c *gin.Context) {
	c.HTML(200, "patientRegister.html", nil)
}

func (h *UserHandler) PatientLogin(c *gin.Context) {

	var input dto.PatientLoginInput
	if err := c.ShouldBind(&input); err != nil {
		logger.Log.Error("validation error", zap.Error(err))
		c.HTML(http.StatusBadRequest, "login.html", gin.H{
			"error": err.Error(),
		})
		return
	}

		user, err := h.Service.Login(input)
	if err != nil {
		logger.Log.Error("Login failed", zap.Error(err))
		c.HTML(http.StatusUnauthorized, "login.html", gin.H{
			"error": "Invalid credentials",
		})
		return
	}

token, err := utils.GenerateJWT(user.UserID, string(user.Role))
if err != nil {
    c.JSON(500, gin.H{
        "error": "could not generate token",
    })
    return
}

c.JSON(200, gin.H{
    "token": token,
})

	//c.HTML(200, "patientLogin.html", nil)
}


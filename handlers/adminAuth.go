package handlers

import (
	"medicity/dto"
	"medicity/logger"
	"medicity/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *UserHandler) AdminLogin(c *gin.Context) {
	var input dto.AdminLoginInput
	if err := c.ShouldBind(&input); err != nil {
		logger.Log.Error("admin login validation failed", zap.Error(err))
		c.HTML(http.StatusBadRequest, "adminLogin.html", gin.H{"error": err.Error()})
		return
	}

	user, err := h.Service.Login(dto.PatientLoginInput{Email: input.Email, Password: input.Password})
	if err != nil {
		logger.Log.Error("admin login failed", zap.Error(err))
		c.HTML(http.StatusUnauthorized, "adminLogin.html", gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.UserID, string(user.Role))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "role": string(user.Role)})
}

func (h *UserHandler) AdminDashboard(c *gin.Context) {
	doctors, err := h.Service.GetAllDoctors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	patients, err := h.Service.GetAllPatients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"doctors": doctors, "patients": patients})
}

func (h *UserHandler) ManageDepartments(c *gin.Context) {
	var payload struct {
		Name        string `form:"name" json:"name" binding:"required"`
		Description string `form:"description" json:"description"`
	}
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	department, err := h.Service.CreateDepartment(payload.Name, payload.Description)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"department": department})
}

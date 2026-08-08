package handlers

import (
	"medicity/dto"
	"medicity/logger"
	"medicity/pkg/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (h *UserHandler) DoctorSignup(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "doctorRegister.html", nil)
		return
	}

	var input dto.DoctorSignupInput
	if err := c.ShouldBind(&input); err != nil {
		logger.Log.Error("doctor signup validation failed", zap.Error(err))
		c.HTML(http.StatusBadRequest, "doctorRegister.html", gin.H{"error": err.Error()})
		return
	}

	user, err := h.Service.RegisterDoctor(input)
	if err != nil {
		logger.Log.Error("doctor registration failed", zap.Error(err))
		c.HTML(http.StatusBadRequest, "doctorRegister.html", gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "doctor account created successfully and pending admin verification",
		"user_id": user.UserID,
		"role":    string(user.Role),
	})
}

func (h *UserHandler) DoctorLogin(c *gin.Context) {
	var input dto.DoctorLoginInput
	if err := c.ShouldBind(&input); err != nil {
		logger.Log.Error("doctor login validation failed", zap.Error(err))
		c.HTML(http.StatusBadRequest, "doctorLogin.html", gin.H{"error": err.Error()})
		return
	}

	user, err := h.Service.Login(dto.PatientLoginInput{Email: input.Email, Password: input.Password})
	if err != nil {
		logger.Log.Error("doctor login failed", zap.Error(err))
		c.HTML(http.StatusUnauthorized, "doctorLogin.html", gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.UserID, string(user.Role))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "role": string(user.Role)})
}

func (h *UserHandler) DoctorDashboard(c *gin.Context) {
	patients, err := h.Service.GetAllPatients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"patients": patients})
}

func (h *UserHandler) CreatePrescription(c *gin.Context) {
	var payload struct {
		AppointmentID        uint   `form:"appointment_id" json:"appointment_id" binding:"required"`
		Complaints           string `form:"complaints" json:"complaints"`
		Diagnosis            string `form:"diagnosis" json:"diagnosis"`
		Advice               string `form:"advice" json:"advice"`
		Medicines            string `form:"medicines" json:"medicines"`
		FollowUpInstructions string `form:"follow_up_instructions" json:"follow_up_instructions"`
		FollowUpDate         string `form:"follow_up_date" json:"follow_up_date"`
	}
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var followUpDate *time.Time
	if payload.FollowUpDate != "" {
		t, err := time.Parse(time.RFC3339, payload.FollowUpDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "follow_up_date must be RFC3339"})
			return
		}
		followUpDate = &t
	}
	prescription, err := h.Service.CreatePrescription(payload.AppointmentID, payload.Complaints, payload.Diagnosis, payload.Advice, payload.Medicines, followUpDate, payload.FollowUpInstructions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"prescription": prescription})
}

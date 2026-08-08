package handlers

import (
	"medicity/dto"
	"medicity/logger"
	"medicity/pkg/utils"
	"medicity/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	Service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

func (h *UserHandler) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

func (h *UserHandler) PatientSignup(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		c.HTML(http.StatusOK, "patientRegister.html", nil)
		return
	}

	var input dto.PatientSignupInput
	if err := c.ShouldBind(&input); err != nil {
		logger.Log.Error("patient signup validation failed", zap.Error(err))
		c.HTML(http.StatusBadRequest, "patientRegister.html", gin.H{"error": err.Error()})
		return
	}

	user, err := h.Service.RegisterPatient(input)
	if err != nil {
		logger.Log.Error("patient registration failed", zap.Error(err))
		c.HTML(http.StatusBadRequest, "patientRegister.html", gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "patient account created successfully",
		"user_id": user.UserID,
		"role":    string(user.Role),
	})
}

func (h *UserHandler) PatientLogin(c *gin.Context) {
	var input dto.PatientLoginInput
	if err := c.ShouldBind(&input); err != nil {
		logger.Log.Error("patient login validation failed", zap.Error(err))
		c.HTML(http.StatusBadRequest, "patientLogin.html", gin.H{"error": err.Error()})
		return
	}

	user, err := h.Service.Login(input)
	if err != nil {
		logger.Log.Error("patient login failed", zap.Error(err))
		c.HTML(http.StatusUnauthorized, "patientLogin.html", gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.UserID, string(user.Role))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token, "role": string(user.Role)})
}

func (h *UserHandler) PatientDashboard(c *gin.Context) {
	doctors, err := h.Service.GetAllDoctors()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"doctors": doctors})
}

func (h *UserHandler) BookAppointment(c *gin.Context) {
	patientID, _ := c.Get("userID")
	patientUint, _ := strconv.ParseUint(strconv.FormatUint(uint64(patientID.(uint)), 10), 10, 10)
	var payload struct {
		DoctorID uint `form:"doctor_id" json:"doctor_id" binding:"required"`
		SlotID   uint `form:"slot_id" json:"slot_id" binding:"required"`
		Fee      int  `form:"consultation_fee" json:"consultation_fee"`
	}
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	appointment, err := h.Service.CreateAppointment(uint(patientUint), payload.DoctorID, payload.SlotID, payload.Fee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"appointment": appointment})
}

func (h *UserHandler) AddDoctorReview(c *gin.Context) {
	patientID, _ := c.Get("userID")
	var payload struct {
		DoctorID      uint   `form:"doctor_id" json:"doctor_id" binding:"required"`
		AppointmentID uint   `form:"appointment_id" json:"appointment_id" binding:"required"`
		Review        string `form:"review" json:"review"`
		Rating        uint8  `form:"rating" json:"rating" binding:"required"`
	}
	if err := c.ShouldBind(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	review, err := h.Service.AddDoctorReview(uint(patientID.(uint)), payload.DoctorID, payload.AppointmentID, payload.Review, payload.Rating)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"review": review})
}

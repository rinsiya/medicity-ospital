package handler

import (
	"errors"
	"net/http"
	"net/url"

	"medicity/internal/dto"
	"medicity/internal/models"
	"medicity/internal/service"
	"medicity/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SignupHandler struct {
	signupService service.SignupService
}

func NewSignupHandler(signupService service.SignupService) *SignupHandler {
	return &SignupHandler{
		signupService: signupService,
	}
}

func (h *SignupHandler) Signup(c *gin.Context) {

	var input dto.SignupInput

	// Log that signup request was received
	logger.Log.Info(
		"Signup request received",
		zap.String("role", c.Param("role")),
	)

	// Bind and validate form
	if err := c.ShouldBind(&input); err != nil {

		logger.Log.Warn(
			"Signup validation failed",
			zap.String("role", c.Param("role")),
			zap.Error(err),
		)

		c.HTML(http.StatusBadRequest, c.Param("role")+"Signup.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	// Determine role from URL
	var role models.UserRole

	switch c.Param("role") {

	case "patient":
		role = models.RolePatient

	case "doctor":
		role = models.RoleDoctor

	default:
		logger.Log.Warn(
			"Invalid signup role",
			zap.String("role", c.Param("role")),
			zap.String("email", input.Email),
		)

		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Invalid signup role",
		})
		return
	}

	// Log signup attempt
	// Never log password or OTP.
	logger.Log.Info(
		"Signup attempt",
		zap.String("role", string(role)),
		zap.String("email", input.Email),
		zap.String("phone", input.Phone),
	)

	// Save signup
	err := h.signupService.Signup(&input, role)

	if err != nil {

		switch {

		case errors.Is(err, service.ErrEmailAlreadyExists):

			logger.Log.Warn(
				"Signup failed: email already exists",
				zap.String("role", string(role)),
				zap.String("email", input.Email),
			)

			c.HTML(http.StatusConflict, string(role)+"Signup.html", gin.H{
				"error": "Email already registered",
			})

		case errors.Is(err, service.ErrPhoneAlreadyExists):

			logger.Log.Warn(
				"Signup failed: phone already exists",
				zap.String("role", string(role)),
				zap.String("phone", input.Phone),
			)

			c.HTML(http.StatusConflict, string(role)+"Signup.html", gin.H{
				"error": "Phone number already registered",
			})

		default:

			logger.Log.Error(
				"Signup failed",
				zap.String("role", string(role)),
				zap.String("email", input.Email),
				zap.Error(err),
			)

			c.HTML(http.StatusInternalServerError, string(role)+"Signup.html", gin.H{
				"error": "Unable to process signup",
			})
		}

		return
	}

	// Signup successful
	logger.Log.Info(
		"Signup successful",
		zap.String("role", string(role)),
		zap.String("email", input.Email),
	)




c.Redirect(
    http.StatusSeeOther,
    "/"+string(role)+"/verify-otp?phone="+url.QueryEscape(input.Phone),
)
}
func (h *SignupHandler) ShowOTPPage(c *gin.Context) {

	role := c.Param("role")
	email := c.Query("email")

	logger.Log.Info(
		"OTP verification page requested",
		zap.String("role", role),
		zap.String("email", email),
	)

	c.HTML(http.StatusOK, "verify-otp.html", gin.H{
		"role":  role,
		"email": email,
	})
}
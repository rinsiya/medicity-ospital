package handler

import (
	"medicity/internal/dto"
	"medicity/internal/models"
	"medicity/internal/service"
	"medicity/logger"
	"medicity/pkg/utils"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) Home(c *gin.Context) {
	c.HTML(200, "home.html", nil)
}

func (h *UserHandler) PatientLogin(c *gin.Context) {
	c.HTML(200, "patientLogin.html", nil)
}

func (h *UserHandler) DoctorLogin(c *gin.Context) {
	c.HTML(200, "doctorLogin.html", nil)
}


func (h *UserHandler) AdminLogin(c *gin.Context) {
	c.HTML(200, "adminLogin.html", nil)
}

func (h *UserHandler) Login(c *gin.Context) {

     role := c.Param("role")

    switch role {
    case "patient", "doctor", "admin":
        // valid role
    default:
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "message": "Invalid role",
        })
        return
    }

input := dto.LoginInput{}
	if err := c.ShouldBind(&input); err != nil {
		logger.Log.Error("Validation error", zap.Error(err))

		c.HTML(http.StatusBadRequest, role+"Login.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.Service.Login(input,role)
	if err != nil {
		logger.Log.Error("Login failed", zap.Error(err))

		c.HTML(http.StatusUnauthorized, role+"Login.html", gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	token, err := utils.GenerateJWT(
		user.UserID,
		string(user.Role),
	)


	if err != nil {
		logger.Log.Error("Token generation failed", zap.Error(err))
c.HTML(http.StatusUnauthorized, role+"Login.html", gin.H{
			"error": "Token generation failed",
		})
		return
	}



	// Store JWT in HttpOnly cookie
	c.SetCookie(
		"access_token", // cookie name
		token,          // token
		86400,          // 24 hours
		"/",            // available throughout the application
		"",             // domain
		false,          // secure - set true when using HTTPS
		true,           // HttpOnly
	)

	// Redirect based on role
	switch user.Role {

	case models.RoleAdmin:
		c.Redirect(http.StatusSeeOther, "/admin/dashboard")

	case models.RolePatient:
		c.Redirect(http.StatusSeeOther, "/patient/home")

	case models.RoleDoctor:
		c.Redirect(http.StatusSeeOther, "/doctor/dashboard")

	default:
		logger.Log.Warn(
			"Unknown user role",
			zap.String("role", string(user.Role)),
		)

		c.HTML(http.StatusForbidden, "login.html", gin.H{
			"error": "Invalid user role",
		})
	}
}

//logout function to clear the JWT cookie and redirect to login page
func (h *UserHandler) Logout(c *gin.Context) {
	// Clear the JWT cookie
	c.SetCookie(
		"access_token",
		"",
		-1, // expire immediately
		"/",
		"",
		false,
		true,
	)

	logger.Log.Info("User logged out")

	c.Redirect(http.StatusSeeOther, "/login")
}
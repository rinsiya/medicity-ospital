package handlers

import (
	"gorm_projectR/dto"
	"gorm_projectR/logger"
	"gorm_projectR/services"
	"net/http"

	"github.com/gin-contrib/sessions"
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
func (h *UserHandler) Signup(c *gin.Context) {
	var input dto.SignupInput
	if err := c.ShouldBind(&input); err != nil {
		logger.Log.Error("validation error", zap.Error(err))
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	err := h.Service.Signup(input)
	if err != nil {
		logger.Log.Error("Signup failed", zap.Error(err))
		c.String(500, "signup failed")
		return
	}
	c.Redirect(http.StatusFound, "/")
}

func (h *UserHandler) Login(c *gin.Context) {
	session := sessions.Default(c)

	var input dto.LoginInput
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

	session.Set("id", user.ID)
	session.Set("name", user.Name)
	session.Set("role", user.Role)
	if err := session.Save(); err != nil {
		logger.Log.Error("Failed to save session", zap.Error(err))
		c.String(http.StatusInternalServerError, "Session error")
		return
	}
	if user.Role == "admin" {
		c.Redirect(http.StatusFound, "/admin/dashboard")
	} else {
		c.Redirect(http.StatusFound, "/home")
	}
}
func (h *UserHandler) Home(c *gin.Context) {
	session := sessions.Default(c)
	id := session.Get("id")
	if id == nil {
		logger.Log.Warn("Unauthorized access attempt")
		c.Redirect(http.StatusFound, "/login")
		return
	}
	logger.Log.Info("Fetching user from database",
		zap.Any("user_id", id))
	user, err := h.Service.GetUserById(id)
	if err != nil {
		logger.Log.Error("Failed to fetch user details",
			zap.Any("user_id", id),
			zap.Error(err))
		c.String(500, "Error fetching user details")
		return
	}
	logger.Log.Info("Home page loaded successfully",
		zap.Int("user_id", user.ID),
		zap.String("email", user.Email),
	)

	c.HTML(200, "home.html", gin.H{
		"users": user,
	})
}
func (h *UserHandler) Logout(c *gin.Context) {
	session := sessions.Default(c)
	logger.Log.Info("User logged out", zap.Any("user_id", session.Get("id")), zap.Any("name", session.Get("name")))
	session.Clear()
	if err := session.Save(); err != nil {
		logger.Log.Error("Failed to clear session", zap.Error(err))
		c.String(http.StatusInternalServerError, "Logout failed")
		return
	}
	c.Redirect(http.StatusFound, "/login")
}

func (h *UserHandler) UpdateProfile(c *gin.Context) {

	session := sessions.Default(c)

	id := session.Get("id")
	if id == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	var input dto.UpdateUserInput
	if err := c.ShouldBind(&input); err != nil {
		logger.Log.Warn("Validation failed", zap.Error(err))
		user, _ := h.Service.GetUserById(id)

		c.HTML(http.StatusBadRequest, "home.html", gin.H{
			"error": err.Error(),
			"users": user,
		})
		return
	}

	err := h.Service.UpdateUser(id, input)
	if err != nil {
		logger.Log.Error("Failed to update user data", zap.Error(err))
		user, _ := h.Service.GetUserById(id)

		c.HTML(http.StatusInternalServerError, "home.html", gin.H{
			"error": err,
			"users": user,
		})
		return
	}
	c.Redirect(http.StatusFound, "/home")
}

func (h *UserHandler) ShowLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
func (h *UserHandler) ShowSignup(c *gin.Context) {
	c.HTML(200, "signup.html", nil)
}

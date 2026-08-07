package routes

import (
	"medicity/handlers"
	"medicity/middleware"
	"medicity/repository"
	"medicity/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	// Repository
	userRepo := repository.NewUserRepository()

	// Service
	userService := services.NewUserService(userRepo)

	// Handler
	userHandler := handlers.NewUserHandler(userService)

	// ----------------------------
	// Public Routes
	// ----------------------------

	router.GET("/", userHandler.Home)

	// Patient Routes
	patient := router.Group("/patient")
	{
		patient.GET("/signup", userHandler.PatientSignup)
		patient.POST("/signup", userHandler.PatientSignup)

		patient.GET("/login", userHandler.PatientLogin)
		patient.POST("/login", userHandler.PatientLogin)
	}

	// Doctor Routes
	doctor := router.Group("/doctor")
	{
		doctor.GET("/signup", userHandler.DoctorSignup)
		doctor.POST("/signup", userHandler.DoctorSignup)

		doctor.GET("/login", userHandler.DoctorLogin)
		doctor.POST("/login", userHandler.DoctorLogin)
	}

	// Admin Login
	admin := router.Group("/admin")
	{
		admin.GET("/login", userHandler.AdminLogin)
		admin.POST("/login", userHandler.AdminLogin)
	}

	// ----------------------------
	// Protected Routes
	// ----------------------------

	auth := router.Group("/")
	auth.Use(middleware.JWTAuth())
	{
		auth.GET("/home", userHandler.Home)
	}

	// Protected Patient Routes
	patientAuth := router.Group("/patient")
	patientAuth.Use(middleware.JWTAuth())
	{
		// Example:
		// patientAuth.GET("/profile", userHandler.PatientProfile)
		// patientAuth.GET("/appointments", userHandler.PatientAppointments)
	}

	// Protected Doctor Routes
	doctorAuth := router.Group("/doctor")
	doctorAuth.Use(middleware.JWTAuth())
	{
		// Example:
		// doctorAuth.GET("/dashboard", userHandler.DoctorDashboard)
	}

	// Protected Admin Routes
	adminAuth := router.Group("/admin")
	adminAuth.Use(middleware.JWTAuth())
	adminAuth.Use(middleware.AdminOnly())
	{
		//adminAuth.GET("/dashboard", userHandler.AdminDashboard)
	}
}
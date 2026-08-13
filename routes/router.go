package routes

import (
	"medicity/handlers"
	"medicity/middleware"
	"medicity/repository"
	"medicity/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	userRepo := repository.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router.GET("/", userHandler.Home)

	patient := router.Group("/patient")
	{
		patient.GET("/signup", userHandler.PatientSignup)
		patient.POST("/signup", userHandler.PatientSignup)
		patient.GET("/login", userHandler.PatientLogin)
		patient.POST("/login", userHandler.PatientLogin)

		patientAuth := patient.Group("")
		patientAuth.Use(middleware.JWTAuth())
		{
			patientAuth.GET("/dashboard", userHandler.PatientDashboard)
			patientAuth.POST("/appointments", userHandler.BookAppointment)
			patientAuth.POST("/review", userHandler.AddDoctorReview)
		}
	}

	doctor := router.Group("/doctor")
	{
		doctor.GET("/signup", userHandler.DoctorSignup)
		doctor.POST("/signup", userHandler.DoctorSignup)
		doctor.GET("/login", userHandler.DoctorLogin)
		doctor.POST("/login", userHandler.DoctorLogin)

		doctorAuth := doctor.Group("")
		doctorAuth.Use(middleware.JWTAuth())
		{
			doctorAuth.GET("/dashboard", userHandler.DoctorDashboard)
			doctorAuth.POST("/prescriptions", userHandler.CreatePrescription)
		}
	}

	admin := router.Group("/admin")
	{
		admin.GET("/login", userHandler.AdminLogin)
		admin.POST("/login", userHandler.AdminLogin)

		adminAuth := admin.Group("")
		adminAuth.Use(middleware.JWTAuth())
		adminAuth.Use(middleware.AdminOnly())
		{
			adminAuth.GET("/dashboard", userHandler.AdminDashboard)
			adminAuth.POST("/departments", userHandler.ManageDepartments)
		}
	}

	auth := router.Group("/api")
	auth.Use(middleware.JWTAuth())
	{
		auth.GET("/home", userHandler.Home)
	}
}

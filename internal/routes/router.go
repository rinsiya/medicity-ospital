package routes

import (
	"medicity/internal/handler"
	"medicity/internal/repository"
	"medicity/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Repository
	userRepo := repository.NewUserRepository()
pendingUserRepo := repository.NewPendingUserSignupRepository()
patientRepo := repository.NewPatientRepository()
doctorRepo := repository.NewDoctorRepository()
	// Service
	userService := service.NewUserService(userRepo)
signupService := service.NewSignupService(pendingUserRepo,userRepo,patientRepo,doctorRepo)
patientService := service.NewPatientService(patientRepo)
doctorService := service.NewDoctorService(doctorRepo)

	// Handler
	patientHandler := handler.NewPatientHandler(patientService)
	doctorHandler := handler.NewDoctorHandler(doctorService)
	
	userHandler := handler.NewUserHandler(userService)
signupHandler:= handler.NewSignupHandler(signupService)


	router.GET("/", userHandler.Home)
	router.GET("/home", userHandler.Home)
	router.POST("/:role/login", userHandler.Login)
	router.POST("/logout", userHandler.Logout)
	router.POST("/signup/:role", signupHandler.Signup)
	
	router.GET("/:role/verify-otp", signupHandler.ShowOTPPage)
	router.POST("/verify-otp", signupHandler.ValidateOTP)

	router.GET("/:role/verify-otp-pending", signupHandler.ShowOTPPendingPage)
router.POST("/resend-otp", signupHandler.ResendOTP)
router.GET("/change-phone",signupHandler.ChangePhone)
router.POST("/change-phone",signupHandler.UpdatePhone)

router.GET("/patient/verification-success", signupHandler.PatientVerificationSuccess)
router.GET("/doctor/verification-success", signupHandler.DoctorVerificationSuccess)
router.GET("/patient/login", userHandler.PatientLogin)
router.GET("/doctor/login", userHandler.DoctorLogin)
router.GET("/admin/login", userHandler.AdminLogin)







	// Patient specific Routes
    router.GET("/patient/signup", patientHandler.PatientSignup)
    router.GET("/doctor/signup", doctorHandler.DoctorSignup)

}
package routes

import (
	"gorm_projectR/handlers"
	"gorm_projectR/middleware"
	"gorm_projectR/repository"
	"gorm_projectR/services"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	userRepo := repository.NewUserRepository()

	userService := services.NewUserService(userRepo)

	userHandler := handlers.NewUserHandler(userService)

	r.GET("/", middleware.CheckSession(), userHandler.ShowLogin)

	r.GET("/login", middleware.CheckSession(), userHandler.ShowLogin)

	r.GET("/signup", userHandler.ShowSignup)
	r.POST("/signup", userHandler.Signup)
	r.POST("/login", userHandler.Login)
	r.GET("/home", userHandler.Home)
	r.GET("/logout", userHandler.Logout)
	r.POST("/profile/update", userHandler.UpdateProfile)
	admin := r.Group("/admin")
	admin.Use(middleware.AdminOnly())
	{
		admin.GET("/dashboard", userHandler.Dashboard)
		admin.GET("/dashboard/:search", userHandler.Dashboard)
		admin.GET("/dashboard/:search/:clear", userHandler.Dashboard)
		admin.GET("/create", func(c *gin.Context) {
			c.HTML(200, "createUser.html", gin.H{})

		})
		admin.POST("/createUser", userHandler.CreateUser)
		admin.GET("/editUsers/:id", userHandler.EditUsers)
		admin.GET("/deleteUser/:id", userHandler.DeleteUser)
		admin.GET("/editRole/:id/:role", userHandler.UpdateUser)

	}
}


// db.Table(orders)
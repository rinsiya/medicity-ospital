package handlers

import (
	//"gorm_projectR/config"
	"gorm_projectR/dto"
	"gorm_projectR/logger"
	//"gorm_projectR/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	//"golang.org/x/crypto/bcrypt"
)

func (h *UserHandler) Dashboard(c *gin.Context) {

	search := c.Query("search")
	logger.Log.Info("Dashboard accessed",
		zap.String("search", search))
	//clear := c.Query("clear")
	users,err := h.Service.GetUsers(search)
	if err!=nil {
			logger.Log.Error("Failed to fetch user data",
	        zap.Error(err),)
		c.String(500,"Error fetching users")
		
		return
	}
	if len(users) == 0{
				logger.Log.Warn("No match found for search key : ",
	         zap.String("search-key",search),)
		c.HTML(200, "dashboard.html",gin.H{
			"empty" : "No match found",
		})
		return
	}
logger.Log.Info("Users fetched successfully ",zap.Int("Count",len(users)),)

	c.HTML(200, "dashboard.html", gin.H{
		"users": users,
	})

}

func (h *UserHandler) EditUsers(c *gin.Context) {

	id := c.Param("id")
	user,err := h.Service.GetUser(id)
	if err != nil{
	c.String(500, "user not exist")
		return
	}

	c.HTML(200, "editUser.html", gin.H{
		"users": user,
	})
}

func (h *UserHandler) UpdateUser(c *gin.Context) {

	id := c.Param("id")
	role := c.Param("role")
logger.Log.Info("Updating user role",
		zap.String("user_id", id),
		zap.String("old_role", role),
		)
	err := h.Service.ToggleRole(id,role)

	if err != nil {
		logger.Log.Error("User Role update failed",zap.String("user_id", id),
		zap.String("old_role", role),
		)
		c.String(500, "Error updating user role")
		return
	}
	c.Redirect(302, "/admin/editUsers/"+id)
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	logger.Log.Info("deleting user ",
		zap.String("user_id", id),
	)
	err := h.Service.DeleteUser(id)
	if err != nil {
		logger.Log.Error("deletion failed ",
		zap.String("user_id", id),
	)
		c.String(500, "Error deleting user")
		return
	}
	c.Redirect(302, "/admin/dashboard")
	//c.HTML(200, "/dashboard", nil)
	//c.Redirect(http.StatusFound, "/admin/dashboard")
}
func (h *UserHandler) CreateUser(c *gin.Context) {

	// name := c.PostForm("name")
	// email := c.PostForm("email")
	// password := c.PostForm("password")
	// role := c.PostForm("role")


	// user := models.User{
	// 	Name:     name,
	// 	Email:    email,0,
	// 	Password: string(hash),
	// 	Role:     role,
	// }

	var input dto.AdminCreateUserInput
	if err:= c.ShouldBind(&input);err!= nil {
logger.Log.Warn("Validation failed", zap.Error(err))
			c.HTML(http.StatusBadRequest, "createUser.html", gin.H{
			"error": err.Error(),
		})
		return
	}
	logger.Log.Info("inserting new user ",
		zap.String("user_email", input.Email),
		zap.String("user_name", input.Name),
		zap.String("user_role", input.Role),

	)	
err := h.Service.CreateUser(input)

	if err != nil {
		logger.Log.Info("inserting new user failed!!! ",
		zap.String("user_email", input.Email),
		zap.String("user_name", input.Name),
		zap.String("user_role", input.Role),

	)
		c.String(500, "failed to create user", err)
		return
	}
	c.Redirect(http.StatusFound, "/admin/dashboard")
}

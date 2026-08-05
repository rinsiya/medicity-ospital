package dto

type AdminCreateUserInput struct {
	Name            string `form:"name" binding:"required,min=3"`
	Email           string `form:"email" binding:"required,email"`
	Password        string `form:"password"`
	Role            string `form:"role" binding:"required,oneof=admin user"`
}

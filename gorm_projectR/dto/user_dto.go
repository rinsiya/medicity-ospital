package dto

type UpdateUserInput struct {
	Name            string `form:"name" binding:"required,min=3"`
	Email           string `form:"email" binding:"required,email"`
	Password        string `form:"password"`
	ConfirmPassword string `form:"confirm_password"`
}

package dto

type SignupInput struct {
	FirstName      string `form:"firstname" json:"firstname" binding:"required,min=3"`
	LastName       string `form:"lastname" json:"lastname" binding:"required,min=3"`
	Email          string `form:"email" json:"email" binding:"required,email"`
	Phone          string `form:"phone" json:"phone" binding:"required"`
	Password       string `form:"password" json:"password" binding:"required,min=6"`
	ConfirmPassword string `form:"confirm_password" json:"confirm_password" binding:"required,eqfield=Password"`
}
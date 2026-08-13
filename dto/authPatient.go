package dto


type PatientLoginInput struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type PatientSignupInput struct {
	FirstName     string `form:"firstname" binding:"required,min=3"`
	LastName     string `form:"lastname" binding:"required,min=3"`
	Email    string `form:"email" binding:"required,email"`
	Phone	string `form:"phone" binding:"required"`
	Password string `form:"password" binding:"required,min=6"`
	ConfirmPassword string `form:"confirm_password" binding:"required,eqfield=Password"`

}
package dto

type VerifyPatientOTPRequest struct {
	Phone string `json:"phone" binding:"required"`
	OTP   string `json:"otp" binding:"required,len=6"`
}

type ResendPatientOTPRequest struct {
	Phone string `json:"phone" binding:"required"`
}

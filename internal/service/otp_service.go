package service
import "log"

type OTPService interface {
	SendOTP(phone string, otp string) error
}



type otpService struct {
}

func NewOTPService() OTPService {
	return &otpService{}
}

func (s *otpService) SendOTP(phone string, otp string) error {

	// Replace this with your SMS provider later.
	log.Printf("OTP for %s: %s", phone, otp)

	return nil
}
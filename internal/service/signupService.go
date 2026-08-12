package service

import (
	"errors"
	"strings"
	"time"

	"medicity/internal/dto"
	"medicity/internal/models"
	"medicity/internal/repository"
	"medicity/logger"
	"medicity/pkg/utils"

	"go.uber.org/zap"
)

var (
	ErrEmailAlreadyExists = errors.New("email already registered")
	ErrPhoneAlreadyExists = errors.New("phone already registered")
)

type SignupService interface {
	Signup(input *dto.SignupInput, role models.UserRole) error
	
}

type signupService struct {
	pendingUserRepo repository.PendingUserSignupRepository
	userRepo        repository.UserRepository
}

func NewSignupService(
	pendingUserRepo repository.PendingUserSignupRepository,
	userRepo repository.UserRepository,
) SignupService {
	return &signupService{
		pendingUserRepo: pendingUserRepo,
		userRepo:        userRepo,
	}
}

func (s *signupService) Signup(
	input *dto.SignupInput,
	role models.UserRole,
) error {

	email := strings.ToLower(strings.TrimSpace(input.Email))
	phone := strings.TrimSpace(input.Phone)

	logger.Log.Info(
		"Signup service started",
		zap.String("role", string(role)),
		zap.String("email", email),
	)

	// Check existing user by email

	user, err := s.userRepo.FindByEmail(email)
	if err != nil {

		logger.Log.Error(
			"Failed to check existing user by email",
			zap.String("email", email),
			zap.Error(err),
		)

		return err
	}

	if user != nil {

		logger.Log.Warn(
			"Signup rejected: email already registered",
			zap.String("email", email),
			zap.String("role", string(role)),
		)

		return ErrEmailAlreadyExists
	}

	// Check existing user by phone
	
	user, err = s.userRepo.FindByPhone(phone)
	if err != nil {

		logger.Log.Error(
			"Failed to check existing user by phone",
			zap.String("phone", phone),
			zap.Error(err),
		)

		return err
	}

	if user != nil {

		logger.Log.Warn(
			"Signup rejected: phone already registered",
			zap.String("phone", phone),
			zap.String("role", string(role)),
		)

		return ErrPhoneAlreadyExists
	}

	// Check pending signup by email

	pendingUser, err := s.pendingUserRepo.FindByEmail(email)
	if err != nil {

		logger.Log.Error(
			"Failed to check pending signup by email",
			zap.String("email", email),
			zap.Error(err),
		)

		return err
	}

	if pendingUser != nil {

		logger.Log.Warn(
			"Signup rejected: email has pending verification",
			zap.String("email", email),
			zap.String("role", string(role)),
		)

		return ErrEmailAlreadyExists
	}

	// Check pending signup by phone

	pendingUser, err = s.pendingUserRepo.FindByPhone(phone)
	if err != nil {

		logger.Log.Error(
			"Failed to check pending signup by phone",
			zap.String("phone", phone),
			zap.Error(err),
		)

		return err
	}

	if pendingUser != nil {

		logger.Log.Warn(
			"Signup rejected: phone has pending verification",
			zap.String("phone", phone),
			zap.String("role", string(role)),
		)

		return ErrPhoneAlreadyExists
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {

		logger.Log.Error(
			"Failed to hash signup password",
			zap.String("email", email),
			zap.Error(err),
		)

		return err
	}

	// Generate OTP

	otp, err := utils.GenerateOTP()
	if err != nil {

		logger.Log.Error(
			"Failed to generate signup OTP",
			zap.String("email", email),
			zap.Error(err),
		)

		return err
	}

	// Hash password

	otpHash, err := utils.HashPassword(otp)
	if err != nil {

		logger.Log.Error(
			"Failed to hash signup OTP",
			zap.String("email", email),
			zap.Error(err),
		)

		return err
	}


	pendingSignup := &models.PendingUserSignup{
		Role:         role,
		FirstName:    input.FirstName,
		LastName:     input.LastName,
		Email:        email,
		Phone:        phone,
		Password:     hashedPassword,
		OTPHash:      otpHash,
		OTPExpiresAt: time.Now().Add(5 * time.Minute),
		CreatedAt:    time.Now(),
	}

	err = s.pendingUserRepo.Create(pendingSignup)
	if err != nil {

		logger.Log.Error(
			"Failed to create pending signup",
			zap.String("email", email),
			zap.String("role", string(role)),
			zap.Error(err),
		)

		return err
	}


	println("OTP:", otp)

	logger.Log.Info(
		"Signup pending OTP verification",
		zap.String("email", email),
		zap.String("role", string(role)),
	)

	return nil
}
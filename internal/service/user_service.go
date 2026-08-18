package service

import (
	"errors"

	"medicity/internal/dto"
	"medicity/internal/models"
	"medicity/internal/repository"
	"medicity/logger"
	"medicity/pkg/utils"

	"go.uber.org/zap"
)

type UserService interface {
	Login(input dto.LoginInput, role string) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) Login(input dto.LoginInput, role string) (*models.User, error) {

	logger.Log.Info("Login attempt",
		zap.String("username", input.Username),
		zap.String("requested_role", role),
	)

	// Find user by email or phone
	user, err := s.repo.FindByUsername(input.Username)

	if err != nil {
		logger.Log.Error("Failed to find user during login",
			zap.String("username", input.Username),
			zap.Error(err),
		)
		return nil, err
	}

	// User not found
	if user == nil {
		logger.Log.Warn("Login failed: user not found",
			zap.String("username", input.Username),
		)
		return nil, ErrInvalidCredentials
	}

	// Check password
	if !utils.CheckPassword(user.Password, input.Password) {
		logger.Log.Warn("Login failed: invalid password",
			zap.String("username", input.Username),
		)
		return nil, ErrInvalidCredentials
	}

	// Check requested role
	if string(user.Role) != role {
		logger.Log.Warn("Login failed: role mismatch",
			zap.String("username", input.Username),
			zap.String("user_role", string(user.Role)),
			zap.String("requested_role", role),
		)
		return nil, ErrInvalidCredentials
	}

	// Check account status
	if user.Status != models.UserActive {
		logger.Log.Warn("Login failed: account is not active",
			zap.String("username", input.Username),
			zap.String("status", string(user.Status)),
		)
		return nil, errors.New("user account is not active")
	}

	logger.Log.Info("Login successful",
		zap.Int("user_id", int(user.UserID)),
		zap.String("role", string(user.Role)),
	)

	return user, nil
}

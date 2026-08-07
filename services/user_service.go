package services

import (
	//"errors"
	"medicity/dto"
	"medicity/models"
	"medicity/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	// Signup(input dto.SignupInput) error
	 Login(input dto.PatientLoginInput) (*models.User, error)
	// GetUserById(id interface{}) (*models.User, error)
	// GetUsers(search string) ([]models.User, error)
	// GetUser(id string) (*models.User, error)
	// ToggleRole(id string, currentRole string) error
	// DeleteUser(id string) error
	// CreateUser(input dto.AdminCreateUserInput) error
	// UpdateUser(id interface{}, input dto.UpdateUserInput) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) Login(input dto.PatientLoginInput) (*models.User, error) {
	user, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, err
	}
	return user, nil
}

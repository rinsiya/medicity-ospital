package services

import (
	"errors"
	"gorm_projectR/dto"
	"gorm_projectR/models"
	"gorm_projectR/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Signup(input dto.SignupInput) error
	Login(input dto.LoginInput) (*models.User, error)
	GetUserById(id interface{}) (*models.User, error)
	GetUsers(search string) ([]models.User, error)
	GetUser(id string) (*models.User, error)
	ToggleRole(id string, currentRole string) error
	DeleteUser(id string) error
	CreateUser(input dto.AdminCreateUserInput) error
	UpdateUser(id interface{}, input dto.UpdateUserInput) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) Signup(input dto.SignupInput) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return err
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hash),
		Role:     "user",
	}
	return s.repo.Create(&user)
}

func (s *userService) Login(input dto.LoginInput) (*models.User, error) {
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

func (s *userService) GetUserById(id interface{}) (*models.User, error) {
	return s.repo.FindById(id)
}
func (s *userService) GetUsers(search string) ([]models.User, error) {
	return s.repo.GetAll(search)
}
func (s *userService) GetUser(id string) (*models.User, error) {
	return s.repo.FindById(id)
}
func (s *userService) ToggleRole(id string, currentRole string) error {
	var newRole string
	if currentRole == "admin" {
		newRole = "user"
	} else {
		newRole = "admin"
	}
	err := s.repo.UpdateRole(id, newRole)
	return err
}
func (s *userService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}
func (s *userService) CreateUser(input dto.AdminCreateUserInput) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return err
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hash),
		Role:     input.Role,
	}
	return s.repo.Create(&user)
}

func (s *userService) UpdateUser(id interface{}, input dto.UpdateUserInput) error {
	if input.Name == "" {
		return errors.New("name is required")
	}

	if input.Email == "" {
		return errors.New("email is required")
	}
	updates := map[string]interface{}{
		"name":  input.Name,
		"email": input.Email,
	}

	if input.Password != "" {
		if len(input.Password) < 6 {
			return errors.New("Password must contain atleast 6 characters")
		}

		if input.Password != input.ConfirmPassword {
			return errors.New("Passwords do not match")
		}
		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(input.Password), bcrypt.DefaultCost,
		)
		if err != nil {
			return err
		}
		updates["password"] = string(hashedPassword)
	}
	return s.repo.UpdateUser(id, updates)
}

package repository

import (
	"errors"
	"medicity/database"
	"medicity/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByEmailOrPhone(username string) (*models.User, error)
	//Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByPhone(phone string) (*models.User, error)
}

  type userRepository struct{}

  func NewUserRepository() UserRepository {
  	return &userRepository{}
  }
func (r *userRepository) FindByEmailOrPhone(username string) (*models.User, error) {
    var user models.User

    err := database.DB.
        Where("email = ? OR phone = ?", username, username).
        First(&user).Error

    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }

    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	err := database.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) FindByPhone(phone string) (*models.User, error) {
	var user models.User

	err := database.DB.Where("phone = ?", phone).First(&user).Error	


	if err != nil {
		return nil, err
	}

	return &user, nil
}

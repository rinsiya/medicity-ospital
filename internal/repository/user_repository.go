package repository

import (
	"errors"
	"medicity/database"
	"medicity/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByUsername(username string) (*models.User, error)
	FindByEmailAndPhone(email,phone string) (*models.User, error)
	FindByEmailOrPhone(email,phone string) (*models.User, error)

	FindByEmail(email string) (*models.User, error)
	FindByPhone(phone string) (*models.User, error)
	EmailExist(email string) (bool, error)		
	PhoneExist(phone string) (bool, error)
	Create(tx *gorm.DB, user *models.User) error
	FetchUser(userID uint) (*models.User, error)
}

  type userRepository struct{}

  func NewUserRepository() UserRepository {
  	return &userRepository{}
  }

  
func (r *userRepository) FindByEmailOrPhone(email, phone string) (*models.User, error) {
    var user models.User

    err := database.DB.
        Where("email = ? OR phone = ?", email, phone).
        First(&user).Error

    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }

    if err != nil {
        return nil, err
    }

    return &user, nil
}
func (r *userRepository) FindByUsername(username string) (*models.User, error) {
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

func (r *userRepository) EmailExist(email string) (bool, error) {
	var count int64

		err := database.DB.Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *userRepository) PhoneExist(phone string) (bool, error) {
	var count int64

		err := database.DB.Model(&models.User{}).Where("phone = ?", phone).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *userRepository) FindByEmailAndPhone(email, phone string) (*models.User, error) {
    var user models.User

    err := database.DB.
        Where("email = ? AND phone = ?", email, phone).
        First(&user).Error

    if errors.Is(err, gorm.ErrRecordNotFound) {
        return nil, nil
    }

    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (r *userRepository) Create(tx *gorm.DB,user *models.User) error {

    return tx.Create(user).Error
}

func (r *userRepository) FetchUser(userID uint) (*models.User, error) {

	var user models.User

	err := database.DB.
		Where("user_id = ?", userID).
		First(&user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &user, nil
}
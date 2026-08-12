package repository

import (
	"errors"

	"medicity/database"
	"medicity/internal/models"

	"gorm.io/gorm"
)

type PendingUserSignupRepository interface {
	Create(pendingUser *models.PendingUserSignup) error
	FindByEmail(email string) (*models.PendingUserSignup, error)
	FindByPhone(phone string) (*models.PendingUserSignup, error)
}

type pendingUserSignupRepository struct {}

func NewPendingUserSignupRepository() PendingUserSignupRepository {
	return &pendingUserSignupRepository{}
}

func (r *pendingUserSignupRepository) Create(pendingUser *models.PendingUserSignup) error {
	return database.DB.Create(pendingUser).Error
}

func (r *pendingUserSignupRepository) FindByEmail(
	email string,
) (*models.PendingUserSignup, error) {

	var pendingUser models.PendingUserSignup

	err := database.DB.
		Where("email = ?", email).
		First(&pendingUser).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &pendingUser, nil
}

func (r *pendingUserSignupRepository) FindByPhone(
	phone string,
) (*models.PendingUserSignup, error) {

	var pendingUser models.PendingUserSignup

	err := database.DB.
		Where("phone = ?", phone).
		First(&pendingUser).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &pendingUser, nil
}
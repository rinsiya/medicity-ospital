package repository

import (
	"errors"

	"medicity/database"
	"medicity/internal/models"

	"gorm.io/gorm"
)

type PendingUserSignupRepository interface {
	Create(pendingUser *models.PendingUserSignup) error
	FindByEmailAndPhone(email ,phone string) (*models.PendingUserSignup, error)
	FindByEmailOrPhone(email ,phone string) (*models.PendingUserSignup, error)
	FindByPhone(phone string) (*models.PendingUserSignup, error)
	Update(pendingUser *models.PendingUserSignup) error
	Delete(tx *gorm.DB, pendingID uint) error
	PhoneExist(phone string)(bool,error)
}

type pendingUserSignupRepository struct {}

func NewPendingUserSignupRepository() PendingUserSignupRepository {
	return &pendingUserSignupRepository{}
}

func (r *pendingUserSignupRepository) Create(pendingUser *models.PendingUserSignup) error {
	return database.DB.Create(pendingUser).Error
}

func (r *pendingUserSignupRepository) FindByEmailAndPhone(email,phone string) (*models.PendingUserSignup, error) {

	var pendingUser models.PendingUserSignup

	err := database.DB.Where("email = ? OR phone= ?", email,phone).First(&pendingUser).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		
		return nil, err
	}

	return &pendingUser, nil
}
func (r *pendingUserSignupRepository) FindByEmailOrPhone(email,phone string) (*models.PendingUserSignup, error) {

	var pendingUser models.PendingUserSignup

	err := database.DB.Where("email = ? OR phone= ?", email,phone).First(&pendingUser).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		
		return nil, err
	}

	return &pendingUser, nil
}
func (r *pendingUserSignupRepository) FindByPhone(phone string) (*models.PendingUserSignup, error) {

	var pendingUser models.PendingUserSignup

	err := database.DB.Where("phone = ?", phone).First(&pendingUser).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &pendingUser, nil
}

func (r *pendingUserSignupRepository) Update(pendingUser *models.PendingUserSignup) error {
	return database.DB.Save(pendingUser).Error
}

func (r *pendingUserSignupRepository) Delete(tx *gorm.DB,pendingID uint) error {


    return tx.Delete(&models.PendingUserSignup{}, pendingID).Error
}

func (r *pendingUserSignupRepository) PhoneExist(phone string) (bool, error) {
	var count int64

		err := database.DB.Model(&models.PendingUserSignup{}).Where("phone = ?", phone).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
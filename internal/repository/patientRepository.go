package repository

import (
	//"errors"

	"gorm.io/gorm"
	"medicity/database"
	"medicity/internal/models"
)


type PatientRepository interface {
	//FindByEmailOrPhone(username string) (*models.User, error)
	//Create(user *models.User) error
}

  type patientRepository struct{}

  func NewPatientRepository() PatientRepository {
  	return &patientRepository{}
  }


func (r *patientRepository) EmailExists(email string) (bool, error) {
	var count int64

	err := database.DB.
		Model(&models.User{}).
		Where("email = ?", email).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *patientRepository) PhoneExists(phone string) (bool, error) {
	var count int64

	err := database.DB.
		Model(&models.User{}).
		Where("phone = ?", phone).
		Count(&count).
		Error

	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *userRepository) CreatePatient(
	user *models.User,
	patient *models.Patient,
) error {

	return database.DB.Transaction(func(tx *gorm.DB) error {

		// Create user first
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		// Assign generated UserID to patient
		patient.UserID = user.UserID

		// Create patient
		if err := tx.Create(patient).Error; err != nil {
			return err
		}

		return nil
	})
}


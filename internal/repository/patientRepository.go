package repository

import (
	//"errors"

	"gorm.io/gorm"
	"medicity/database"
	"medicity/internal/models"
)


type PatientRepository interface {
	//FindByEmailOrPhone(username string) (*models.User, error)
	CreatePatient(tx *gorm.DB,patient *models.Patient) error
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

func (r *patientRepository) CreatePatient(tx *gorm.DB,patient *models.Patient) error {
		// Create patient
		return tx.Create(patient).Error;
			
}


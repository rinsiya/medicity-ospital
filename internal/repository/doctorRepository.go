package repository

import (
	"medicity/internal/models"

	"gorm.io/gorm"
)



type DoctorRepository interface {
	CreateDoctor(tx *gorm.DB, doctor *models.Doctor) error
}

  type doctorRepository struct{}

  func NewDoctorRepository() DoctorRepository {
  	return &doctorRepository{}
  }


  func(r *doctorRepository) CreateDoctor(tx *gorm.DB, doctor *models.Doctor) error{

		return tx.Create(doctor).Error;


  }
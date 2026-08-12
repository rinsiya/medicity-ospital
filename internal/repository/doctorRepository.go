package repository

import (

)
type DoctorRepository interface {
	
}

  type doctorRepository struct{}

  func NewDoctorRepository() DoctorRepository {
  	return &doctorRepository{}
  }
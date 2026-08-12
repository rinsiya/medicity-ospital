package service

import (

	"medicity/internal/repository"

)

type PatientService interface {
}

type patientService struct {
	patientRepo        repository.PatientRepository
}

func NewPatientService(patientRepo repository.PatientRepository) PatientService {
	return &patientService{
		patientRepo: patientRepo,
	}

}


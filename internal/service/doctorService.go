package service

import "medicity/internal/repository"

type DoctorService interface {
}

type doctorService struct {
	doctorRepo        repository.DoctorRepository
}

func NewDoctorService(doctorRepo repository.DoctorRepository) DoctorService {
	return &doctorService{
		doctorRepo: doctorRepo,
	}

}


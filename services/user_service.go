package services

import (
	"errors"
	"medicity/dto"
	"medicity/models"
	"medicity/repository"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
)

type UserService interface {
	RegisterPatient(input dto.PatientSignupInput) (*models.User, error)
	RegisterDoctor(input dto.DoctorSignupInput) (*models.User, error)
	Login(input dto.PatientLoginInput) (*models.User, error)
	GetAllDoctors() ([]models.Doctor, error)
	GetAllPatients() ([]models.Patient, error)
	CreateDepartment(name string, description string) (*models.Department, error)
	CreateAppointment(patientID, doctorID, slotID uint, consultationFee int) (*models.Appointment, error)
	CreatePrescription(appointmentID uint, complaints, diagnosis, advice, medicines string, followUpDate *time.Time, followUpInstructions string) (*models.Prescription, error)
	AddDoctorReview(patientID, doctorID, appointmentID uint, review string, rating uint8) (*models.DoctorReview, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) RegisterPatient(input dto.PatientSignupInput) (*models.User, error) {
	if input.Password != input.ConfirmPassword {
		return nil, errors.New("passwords do not match")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Role:     models.RolePatient,
		Email:    input.Email,
		Phone:    input.Phone,
		Password: string(hash),
		Status:   models.UserActive,
	}
	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	patient := &models.Patient{
		UserID:    user.UserID,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Gender:    "not_set",
	}
	if err := s.repo.CreatePatient(patient); err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}

func (s *userService) RegisterDoctor(input dto.DoctorSignupInput) (*models.User, error) {
	if input.Password != input.ConfirmPassword {
		return nil, errors.New("passwords do not match")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Role:     models.RoleDoctor,
		Email:    input.Email,
		Phone:    input.Phone,
		Password: string(hash),
		Status:   models.UserActive,
	}
	if err := s.repo.CreateUser(user); err != nil {
		return nil, err
	}

	doctor := &models.Doctor{
		UserID:             user.UserID,
		FirstName:          input.FirstName,
		LastName:           input.LastName,
		DepartmentID:       1,
		ConsultationFee:    500,
		VerificationStatus: models.PendingVerification,
	}
	if err := s.repo.CreateDoctor(doctor); err != nil {
		return nil, err
	}

	user.Password = ""
	return user, nil
}

func (s *userService) Login(input dto.PatientLoginInput) (*models.User, error) {
	user, err := s.repo.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) GetAllDoctors() ([]models.Doctor, error) {
	return s.repo.GetDoctors()
}

func (s *userService) GetAllPatients() ([]models.Patient, error) {
	return s.repo.GetPatients()
}

func (s *userService) CreateDepartment(name string, description string) (*models.Department, error) {
	department := &models.Department{
		DepartmentName: name,
		Description:    description,
		Status:         models.DepartmentActive,
	}
	if err := s.repo.CreateDepartment(department); err != nil {
		return nil, err
	}
	return department, nil
}

func (s *userService) CreateAppointment(patientID, doctorID, slotID uint, consultationFee int) (*models.Appointment, error) {
	appointment := &models.Appointment{
		PatientID:       patientID,
		DoctorID:        doctorID,
		SlotID:          slotID,
		ConsultationFee: consultationFee,
		Status:          models.AppointmentConfirmed,
	}
	if err := s.repo.CreateAppointment(appointment); err != nil {
		return nil, err
	}
	return appointment, nil
}

func (s *userService) CreatePrescription(appointmentID uint, complaints, diagnosis, advice, medicines string, followUpDate *time.Time, followUpInstructions string) (*models.Prescription, error) {
	prescription := &models.Prescription{
		AppointmentID:        appointmentID,
		Complaints:           complaints,
		Diagnosis:            diagnosis,
		Advice:               advice,
		Medicines:            datatypes.JSON([]byte(medicines)),
		FollowUpDate:         followUpDate,
		FollowUpInstructions: followUpInstructions,
	}
	if err := s.repo.CreatePrescription(prescription); err != nil {
		return nil, err
	}
	return prescription, nil
}

func (s *userService) AddDoctorReview(patientID, doctorID, appointmentID uint, review string, rating uint8) (*models.DoctorReview, error) {
	if rating > 5 {
		return nil, errors.New("rating must be between 1 and 5")
	}
	docReview := &models.DoctorReview{
		PatientID:     patientID,
		DoctorID:      doctorID,
		AppointmentID: appointmentID,
		Review:        review,
		Rating:        rating,
	}
	if err := s.repo.CreateReview(docReview); err != nil {
		return nil, err
	}
	return docReview, nil
}

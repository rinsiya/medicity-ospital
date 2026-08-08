package repository

import (
	"medicity/database"
	"medicity/models"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindByID(id uint) (*models.User, error)
	CreatePatient(patient *models.Patient) error
	CreateDoctor(doctor *models.Doctor) error
	GetPatients() ([]models.Patient, error)
	GetDoctors() ([]models.Doctor, error)
	CreateDepartment(department *models.Department) error
	ListDepartments() ([]models.Department, error)
	CreateAppointment(appointment *models.Appointment) error
	CreatePrescription(prescription *models.Prescription) error
	CreateReview(review *models.DoctorReview) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *userRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := database.DB.First(&user, id).Error
	return &user, err
}

func (r *userRepository) CreatePatient(patient *models.Patient) error {
	return database.DB.Create(patient).Error
}

func (r *userRepository) CreateDoctor(doctor *models.Doctor) error {
	return database.DB.Create(doctor).Error
}

func (r *userRepository) GetPatients() ([]models.Patient, error) {
	var patients []models.Patient
	err := database.DB.Find(&patients).Error
	return patients, err
}

func (r *userRepository) GetDoctors() ([]models.Doctor, error) {
	var doctors []models.Doctor
	err := database.DB.Find(&doctors).Error
	return doctors, err
}

func (r *userRepository) CreateDepartment(department *models.Department) error {
	return database.DB.Create(department).Error
}

func (r *userRepository) ListDepartments() ([]models.Department, error) {
	var departments []models.Department
	err := database.DB.Find(&departments).Error
	return departments, err
}

func (r *userRepository) CreateAppointment(appointment *models.Appointment) error {
	return database.DB.Create(appointment).Error
}

func (r *userRepository) CreatePrescription(prescription *models.Prescription) error {
	return database.DB.Create(prescription).Error
}

func (r *userRepository) CreateReview(review *models.DoctorReview) error {
	return database.DB.Create(review).Error
}

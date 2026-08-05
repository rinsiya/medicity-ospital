package repository

import (
	"errors"
	"gorm_projectR/config"
	"gorm_projectR/models"
	"strings"

	"gorm.io/gorm"
	//"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User, error)
	FindById(id interface{}) (*models.User, error)
	GetAll(search string) ([]models.User, error)
	UpdateRole(id string, newRole string) error
	Delete(id string) error
	UpdateUser(id interface{}, input map[string]interface{}) error
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) error {
	return config.DB.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
func (r *userRepository) FindById(id interface{}) (*models.User, error) {
	var user models.User
	err := config.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}
func (r *userRepository) GetAll(search string) ([]models.User, error) {
	var users []models.User
	query := config.DB.Select("id,name,email,role")

	if search != "" {
		query = query.Where(
			"name ILIKE ? OR email ILIKE ? OR CAST(id AS TEXT) ILIKE ?",
			"%"+search+"%",
			"%"+search+"%",
			"%"+search+"%",
		)
	}
	err := query.Find(&users).Error
	return users, err
}
func (r *userRepository) UpdateRole(id string, newRole string) error {
	return config.DB.Model(&models.User{}).
		Where("id = ?", id).
		Update("role", newRole).Error
}
func (r *userRepository) Delete(id string) error {
	return config.DB.Delete(&models.User{}, id).Error
}

func (r *userRepository) UpdateUser(id interface{}, updates map[string]interface{}) error {

	result := config.DB.
		Model(&models.User{}).
		Where("id = ?", id).
		Updates(updates)

	if result.Error != nil {

		// duplicate email
		if strings.Contains(result.Error.Error(), "duplicate key") {
			return errors.New("email already exists")
		}

		// record not found
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}

		return result.Error
	}

	// no rows updated
	if result.RowsAffected == 0 {
		return errors.New("no user updated")
	}

	return nil
}

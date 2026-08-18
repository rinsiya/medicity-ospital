package models

import "time"

type UserRole string

const (
	RoleAdmin   UserRole = "admin"
	RoleDoctor  UserRole = "doctor"
	RolePatient UserRole = "patient"
)

type UserStatus string

const (
	UserActive  UserStatus = "active"
	UserBlocked UserStatus = "blocked"
	UserDeleted UserStatus = "deleted"
)

type User struct {
	UserID uint `gorm:"primaryKey;column:user_id"`

	Role UserRole `gorm:"type:varchar(20);not null"`

	Email string `gorm:"size:150;not null;uniqueIndex"`

	Phone string `gorm:"size:20;not null;uniqueIndex"`

	Password string `gorm:"size:255;not null"`

	Status UserStatus `gorm:"type:varchar(20);not null;default:'active'"`

	LastLogin *time.Time

	CreatedAt time.Time `gorm:"autoCreateTime"`

	UpdatedAt time.Time `gorm:"autoUpdateTime"`

}
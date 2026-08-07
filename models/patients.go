package models

import "time"

type Patient struct {
	PatientID uint `gorm:"primaryKey;column:patient_id"`

	UserID uint `gorm:"not null;uniqueIndex"`

	FirstName string `gorm:"size:50;not null"`

	LastName string `gorm:"size:50;not null"`

	Gender string `gorm:"size:10;not null"`

	DOB time.Time `gorm:"type:date;not null"`

	ProfilePhotoID *uint

	CreatedAt time.Time `gorm:"autoCreateTime"`

	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	
	// User User `gorm:"foreignKey:UserID;references:UserID"`

	// ProfilePhoto *File `gorm:"foreignKey:ProfilePhotoID;references:FileID"`

	// Address *Address `gorm:"foreignKey:PatientID"`

	// Appointments []Appointment `gorm:"foreignKey:PatientID"`

	// VitalData []VitalData `gorm:"foreignKey:PatientID"`

	// Reviews []DoctorReview `gorm:"foreignKey:PatientID"`
}
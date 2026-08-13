package models

import "time"

type DoctorReview struct {
	ReviewID uint `gorm:"primaryKey;column:review_id"`

	PatientID uint `gorm:"not null;index"`

	DoctorID uint `gorm:"not null;index"`

	AppointmentID uint `gorm:"not null;uniqueIndex"`

	Review string `gorm:"type:text"`

	Rating uint8 `gorm:"not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`

	// Patient Patient `gorm:"foreignKey:PatientID;references:PatientID"`

	// Doctor Doctor `gorm:"foreignKey:DoctorID;references:DoctorID"`

	// Appointment Appointment `gorm:"foreignKey:AppointmentID;references:AppointmentID"`
}

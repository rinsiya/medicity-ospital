package models

import (
	"time"

	"gorm.io/datatypes"
)

type Prescription struct {
	PrescriptionID uint `gorm:"primaryKey;column:prescription_id"`

	AppointmentID uint `gorm:"not null;uniqueIndex"`

	Complaints string `gorm:"type:text"`

	Diagnosis string `gorm:"type:text"`

	Advice string `gorm:"type:text"`

	Medicines datatypes.JSON `gorm:"type:jsonb;not null"`

	FollowUpDate *time.Time

	FollowUpInstructions string `gorm:"type:text"`

	CreatedAt time.Time `gorm:"autoCreateTime"`

  	// Appointment Appointment `gorm:"foreignKey:AppointmentID;references:AppointmentID"`
}

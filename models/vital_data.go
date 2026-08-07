package models

import (
	"time"

	"gorm.io/datatypes"
)

type VitalData struct {
	EntryID uint `gorm:"primaryKey;column:entry_id"`

	PatientID uint `gorm:"not null;index"`

	Vitals datatypes.JSON `gorm:"type:jsonb;not null"`

	CreatedAt time.Time `gorm:"autoCreateTime"`

	//Patient Patient `gorm:"foreignKey:PatientID;references:PatientID"`
}
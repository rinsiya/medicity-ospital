package models

import "time"

type SlotStatus string

const (
	SlotAvailable SlotStatus = "available"
	SlotBooked    SlotStatus = "booked"
	SlotCompleted SlotStatus = "completed"
	SlotCancelled SlotStatus = "cancelled"
)

type DoctorTimeSlot struct {
	SlotID uint `gorm:"primaryKey;column:slot_id"`

	DoctorID uint `gorm:"not null;index"`

	Date time.Time `gorm:"type:date;not null"`

	StartTime time.Time `gorm:"type:time;not null"`

	EndTime time.Time `gorm:"type:time;not null"`

	Status SlotStatus `gorm:"type:varchar(20);not null;default:'available'"`

	CreatedAt time.Time `gorm:"autoCreateTime"`

	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	//Doctor Doctor `gorm:"foreignKey:DoctorID;references:DoctorID"`

	// Appointment *Appointment `gorm:"foreignKey:SlotID"`
}
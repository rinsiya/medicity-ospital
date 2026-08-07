package models

import "time"

type CancelledBy string

const (
	CancelledByPatient CancelledBy = "patient"
	CancelledByDoctor  CancelledBy = "doctor"
)

type RefundStatus string

const (
	RefundPending     RefundStatus = "pending"
	RefundTransferred RefundStatus = "transferred"
)

type CancelledAppointment struct {
	CancelledID uint `gorm:"primaryKey;column:cancelled_id"`

	AppointmentID uint `gorm:"not null;uniqueIndex"`

	CancelledBy CancelledBy `gorm:"type:varchar(20);not null"`

	CancelledAt time.Time `gorm:"autoCreateTime"`

	RefundStatus RefundStatus `gorm:"type:varchar(20);not null;default:'pending'"`

//	Appointment Appointment `gorm:"foreignKey:AppointmentID;references:AppointmentID"`
}

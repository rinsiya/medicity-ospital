package models

import "time"

type PaymentMethod string

const (
	Razorpay PaymentMethod = "razorpay"
	Stripe   PaymentMethod = "stripe"
	Cash     PaymentMethod = "cash"
)

type PaymentStatus string

const (
	PaymentPending PaymentStatus = "pending"
	PaymentSuccess PaymentStatus = "success"
	PaymentFailed  PaymentStatus = "failed"
)

type Payment struct {
	PaymentID uint `gorm:"primaryKey;column:payment_id"`

	AppointmentID uint `gorm:"not null;uniqueIndex"`

	Amount int `gorm:"not null"`

	PaymentMethod PaymentMethod `gorm:"type:varchar(20);not null"`

	PaymentGateway string `gorm:"size:30"`

	GatewayPaymentID string `gorm:"size:100"`

	GatewayOrderID string `gorm:"size:100"`

	GatewaySignature string `gorm:"size:255"`

	Status PaymentStatus `gorm:"type:varchar(20);not null;default:'pending'"`

	PaidAt *time.Time

	CreatedAt time.Time `gorm:"autoCreateTime"`

	//Appointment Appointment `gorm:"foreignKey:AppointmentID;references:AppointmentID"`

	//Refund *Refund `gorm:"foreignKey:PaymentID"`
}
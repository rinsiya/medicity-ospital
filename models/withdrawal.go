package models

import "time"

type WithdrawalStatus string

const (
	WithdrawalRequested   WithdrawalStatus = "requested"
	WithdrawalProcessing  WithdrawalStatus = "processing"
	WithdrawalTransferred WithdrawalStatus = "transferred"
	WithdrawalRejected    WithdrawalStatus = "rejected"
)

type Withdrawal struct {
	RequestID uint `gorm:"primaryKey;column:request_id"`

	DoctorID uint `gorm:"not null;index"`

	AccountID uint `gorm:"not null;index"`

	Amount int `gorm:"not null"`

	CurrentBalance int `gorm:"not null"`

	Status WithdrawalStatus `gorm:"type:varchar(20);not null;default:'requested'"`

	CreatedAt time.Time `gorm:"autoCreateTime"`

	ProcessedAt *time.Time

	 //Doctor Doctor `gorm:"foreignKey:DoctorID;references:DoctorID"`

	 //Account BankAccount `gorm:"foreignKey:AccountID;references:AccountID"`
}
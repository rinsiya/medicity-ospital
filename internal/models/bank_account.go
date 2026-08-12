package models

import "time"

type BankAccount struct {
	AccountID uint `gorm:"primaryKey;column:account_id"`

	DoctorID uint `gorm:"not null;index"`

	AccountHolderName string `gorm:"size:100;not null"`

	BankName string `gorm:"size:100;not null"`

	AccountNumber string `gorm:"size:30;not null"`

	AccountType string `gorm:"size:20;not null"`

	IFSCCode string `gorm:"size:20;not null"`

	UPIID string `gorm:"size:100"`

	Verified bool `gorm:"default:false"`

	CreatedAt time.Time `gorm:"autoCreateTime"`

	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Doctor Doctor `gorm:"foreignKey:DoctorID;references:DoctorID"`

	// Withdrawals []Withdrawal `gorm:"foreignKey:AccountID"`
}

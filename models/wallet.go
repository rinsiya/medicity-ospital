package models

import "time"

type Wallet struct {
	WalletID uint `gorm:"primaryKey;column:wallet_id"`

	DoctorID uint `gorm:"not null;uniqueIndex"`

	AvailableBalance int `gorm:"not null;default:0"`

	PendingBalance int `gorm:"not null;default:0"`

	TotalEarned int `gorm:"not null;default:0"`

	CreatedAt time.Time `gorm:"autoCreateTime"`

	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Doctor Doctor `gorm:"foreignKey:DoctorID;references:DoctorID"`
}
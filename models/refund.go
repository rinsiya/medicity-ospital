package models

import "time"

type RefundProcessStatus string

const (
	RefundPendingStatus    RefundProcessStatus = "pending"
	RefundProcessingStatus RefundProcessStatus = "processing"
	RefundSuccessStatus    RefundProcessStatus = "success"
	RefundFailedStatus     RefundProcessStatus = "failed"
)

type Refund struct {
	RefundID uint `gorm:"primaryKey;column:refund_id"`

	PaymentID uint `gorm:"not null;uniqueIndex"`

	RefundAmount int `gorm:"not null"`

	GatewayRefundID string `gorm:"size:100"`

	Status RefundProcessStatus `gorm:"type:varchar(20);not null;default:'pending'"`

	RefundedAt *time.Time

	ProcessedAt *time.Time

	CreatedAt time.Time `gorm:"autoCreateTime"`

	//Payment Payment `gorm:"foreignKey:PaymentID;references:PaymentID"`
}

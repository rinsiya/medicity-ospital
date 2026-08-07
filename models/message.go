package models

import "time"

type Message struct {
	MessageID uint `gorm:"primaryKey;column:message_id"`

	SenderID uint

	ReceiverID uint

	Message string `gorm:"type:text"`

	SentAt time.Time

	ReadAt *time.Time

	//Sender User `gorm:"foreignKey:SenderID;references:UserID"`

	// Receiver User `gorm:"foreignKey:ReceiverID;references:UserID"`
}
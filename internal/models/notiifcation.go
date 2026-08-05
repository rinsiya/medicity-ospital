package models

import "time"

type NotificationStatus string

const (
	NotificationRead   NotificationStatus = "read"
	NotificationUnread NotificationStatus = "unread"
)

type Notification struct {
	NotificationID uint `gorm:"primaryKey;column:notification_id"`

	UserID uint `gorm:"not null;index"`

	Title string `gorm:"size:200;not null"`

	Message string `gorm:"type:text;not null"`

	Status NotificationStatus `gorm:"type:varchar(20);not null;default:'unread'"`

	CreatedAt time.Time `gorm:"autoCreateTime"`

	User User `gorm:"foreignKey:UserID;references:UserID"`
}
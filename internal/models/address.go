package models

type Address struct {
	AddressID uint `gorm:"primaryKey;column:address_id"`

	PatientID uint `gorm:"unique;not null"`

	Address string `gorm:"size:100"`

	Place string `gorm:"size:50"`

	Country string

	Patient Patient `gorm:"foreignKey:PatientID;references:PatientID"`
}
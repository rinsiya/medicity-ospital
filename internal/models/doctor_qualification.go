package models

type DoctorQualification struct {
	QualificationID uint `gorm:"primaryKey;column:qualification_id"`

	DoctorID uint `gorm:"not null;index"`

	Qualification string `gorm:"size:100;not null"`

	Branch string `gorm:"size:150;not null"`

	University string `gorm:"size:150;not null"`

	Year uint16 `gorm:"not null"`

	//Doctor Doctor `gorm:"foreignKey:DoctorID;references:DoctorID"`
}
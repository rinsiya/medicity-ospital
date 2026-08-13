package models

import "time"

type FileCategory string

const (
	DoctorCertificate FileCategory = "doctor_certificate"
	MedicalReport     FileCategory = "medical_report"
	IdentityProof     FileCategory = "identity_proof"
	ProfilePhoto      FileCategory = "profile_photo"
	Other             FileCategory = "other"
)

type File struct {
	FileID uint `gorm:"primaryKey;column:file_id"`

	UserID uint `gorm:"not null;index"`

	Category FileCategory `gorm:"type:varchar(30);not null"`

	StoragePath string `gorm:"size:255;not null"`

	FileName string `gorm:"size:255;not null"`

	FileType string `gorm:"size:100;not null"`

	Remarks string `gorm:"type:text"`

	UploadedAt time.Time `gorm:"autoCreateTime"`

	//User User `gorm:"foreignKey:UserID;references:UserID"`
}
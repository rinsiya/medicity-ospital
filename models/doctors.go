package models


type VerificationStatus string

const (
	PendingVerification VerificationStatus = "pending"
	Verified            VerificationStatus = "verified"
	Rejected            VerificationStatus = "rejected"
)

type Doctor struct {
	DoctorID uint `gorm:"primaryKey;column:doctor_id"`

	UserID uint `gorm:"not null;uniqueIndex"`

	FirstName string `gorm:"size:50;not null"`

	LastName string `gorm:"size:50;not null"`

	DepartmentID uint `gorm:"not null;index"`

	ConsultationFee int `gorm:"not null"`

	VerificationStatus VerificationStatus `gorm:"type:varchar(20);not null;default:'pending'"`

	User User `gorm:"foreignKey:UserID;references:UserID"`

	Department Department `gorm:"foreignKey:DepartmentID;references:DepartmentID"`

	Profile *DoctorProfile `gorm:"foreignKey:DoctorID"`

	Qualifications []DoctorQualification `gorm:"foreignKey:DoctorID"`

	TimeSlots []DoctorTimeSlot `gorm:"foreignKey:DoctorID"`

	Appointments []Appointment `gorm:"foreignKey:DoctorID"`

	Reviews []DoctorReview `gorm:"foreignKey:DoctorID"`

	Wallet *Wallet `gorm:"foreignKey:DoctorID"`

	BankAccount *BankAccount `gorm:"foreignKey:DoctorID"`

	Withdrawals []Withdrawal `gorm:"foreignKey:DoctorID"`
}
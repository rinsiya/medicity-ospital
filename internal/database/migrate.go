package database

import (
	"medicity/internal/logger"
	"medicity/internal/models"

	"go.uber.org/zap"
)

func MigrateDatabase() error {

	err := DB.AutoMigrate(

		&models.User{},
		&models.Patient{},
		&models.Doctor{},
		&models.Address{},
		&models.Department{},

		&models.File{},
		&models.DoctorProfile{},
		&models.DoctorQualification{},
		&models.DoctorTimeSlot{},
		&models.Appointment{},
		&models.CancelledAppointment{},
		&models.Prescription{},
		&models.VitalData{},
		&models.Message{},
		&models.DoctorReview{},
		&models.BankAccount{},
		&models.Wallet{},
		&models.Payment{},
		&models.Refund{},
		&models.Withdrawal{},
		&models.Notification{},
	)

	if err != nil {
		logger.Log.Fatal("Migration failed", zap.Error(err))
		return err
	}

	logger.Log.Info("Database migrated successfully")

	return nil
}

package database

import (
	"medicity/logger"
	"medicity/models"

	"go.uber.org/zap"
)

func MigrateDatabase() error {
	logger.Log.Info("Database migration initialized")

	err := DB.AutoMigrate(
		&models.User{},
		&models.Patient{},
		&models.Department{},
		&models.Doctor{},
		&models.Wallet{},
		&models.Appointment{},
		&models.Prescription{},
		&models.DoctorReview{},
		&models.Message{},
	)

	if err != nil {
		logger.Log.Fatal("Migration failed", zap.Error(err))
		return err
	}

	logger.Log.Info("Database migrated successfully")

	return nil
}

package config

import (
	"gorm_projectR/logger"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

func ConnectDB() {

	connStr := "user=ginuser password=Password dbname=gormdb host=localhost port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		
	})

	if err != nil {
		logger.Log.Fatal("Database connection failed ",zap.Error( err))
	}
	DB = db
	logger.Log.Info("Database connected successsfully")
}

package main

import (
	"fmt"
	"log"
	"medicity/config"
	"medicity/internal/database"
	"medicity/internal/logger"
)

func main() {

	confg := config.LoadConfig()
	if err := logger.InitLogger(); err != nil {
		log.Fatal(err)
	}

	defer logger.Sync()

	logger.Log.Info("Logger initialized")
	//connect DB
	database.ConnectDatabase(confg)

	// Run migrations
    if err := database.MigrateDatabase(); err != nil {
        panic(err)
    }
	fmt.Println(confg.AppName)
	fmt.Println(confg.Port)

}

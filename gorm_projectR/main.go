package main

import (
	"gorm_projectR/config"
	"gorm_projectR/logger"

	//"gorm_project/middleware"
	"gorm_projectR/models"
	"gorm_projectR/routes"

	//"log/slog"
	//"os"

	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger.InitLogger()
	defer logger.Log.Sync()
	//connect DB
	config.ConnectDB()

	var dbName string
	config.DB.Raw("SELECT current_database()").Scan(&dbName)
	logger.Log.Info("Connected DB", zap.String("db", dbName))
	err := config.DB.AutoMigrate(&models.User{})
	if err != nil {
		logger.Log.Fatal("Migration failed", zap.Error(err))
	}
	// slogger := slog.New(
	// 	slog.NewTextHandler(os.Stdout, nil),
	// )
	//create Gin router
	r := gin.New()
	//r.Use(gin.Recovery())

	//	r.SetTrustedProxies(nil)
	//	r.Use(middleware.RequestLogger(logger.Log))
	//r.Use(middleware.RequestDataLogger(logger.Log))
	r.Use(ginzap.Ginzap(logger.Log, time.RFC3339, true))
	r.Use(ginzap.Ginzap(logger.Log, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger.Log, true))

	//load HTML templates
	r.LoadHTMLGlob("templates/*")
	//load static files
	r.Static("/static", "./static")
	store := cookie.NewStore([]byte("secret"))

	r.Use(sessions.Sessions("mysession", store))
	routes.SetupRoutes(r)
	logger.Log.Info("Server started", zap.String("port", "8080"))
	//Run server
	r.Run(":8081")
}

package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"uplift-force-backend/config"
	"uplift-force-backend/models"
	"uplift-force-backend/routes"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// 连接数据库
	config.ConnectDatabase()

	// 自动迁移数据库
	config.DB.AutoMigrate(&models.User{})

	// 设置路由
	r := routes.SetupRoutes()

	// 启动服务器
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}

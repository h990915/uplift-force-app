package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"uplift-force-backend/config"
	_ "uplift-force-backend/docs"
	"uplift-force-backend/models"
	"uplift-force-backend/routes"
)

// @title           Uplift Force Backend API
// @version         1.0
// @description     This is the backend API for Uplift Force DApp
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      45.32.67.85:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// 连接数据库
	config.ConnectDatabase()

	// 自动迁移数据库
	config.DB.AutoMigrate(&models.User{})

	// 设置路由（包含Swagger路由）
	r := routes.SetupRoutes()

	// 启动服务器
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)
}

package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
	"uplift-force-backend/controllers"
	"uplift-force-backend/middleware"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// 配置CORS中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://45.32.67.85:3000", "https://uplift.yoosmart.top", "https://api.yoosmart.top"}, // 允许的前端域名
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Swagger文档路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 公开路由（无需认证）
	public := r.Group("/api/v1")
	{
		// 钱包认证相关
		public.POST("/auth/login", controllers.Login)             // 签名登录
		public.POST("/auth/register", controllers.Register)       // 钱包注册
		public.POST("/auth/refresh", controllers.RefreshToken)    // 刷新token
		public.POST("/auth/verify", controllers.VerifyWallet)     // 验证钱包（仅验签）
		public.POST("/auth/checkWallet", controllers.CheckWallet) // 检查钱包是否已注册
	}

	// 需要认证的路由
	protected := r.Group("/api/v1")
	protected.Use(middleware.JWTMiddleware())
	{
		// 用户相关
		protected.GET("/auth/profile", controllers.GetProfile)
		protected.POST("/auth/logout", controllers.Logout)
		protected.GET("/users", controllers.GetUsers)
		protected.GET("/users/:id", controllers.GetUserByID)
		protected.GET("/users/wallet/:address", controllers.GetUserByWallet)
		protected.PUT("/users/:id", controllers.UpdateUser)
		protected.PUT("/users/:id/login", controllers.UpdateLastLogin)

		// 管理员专用路由
		admin := protected.Group("")
		admin.Use(middleware.RequireAdmin())
		{
			admin.DELETE("/users/:id", controllers.DeleteUser)
		}
	}

	return r
}

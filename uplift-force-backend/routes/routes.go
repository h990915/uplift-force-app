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
		//AllowOrigins:     []string{"http://localhost:3000", "http://45.32.67.85:3000", "https://uplift.yoosmart.top", "https://api.yoosmart.top"},
		AllowOrigins:     []string{"*"}, // 测试环境允许所有跨域请求
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

		public.GET("/orders", controllers.GetOrders)          // 获取订单列表（支持筛选可接单的订单）
		public.GET("/orders/:id", controllers.GetOrderDetail) // 获取订单详情
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

		// 订单相关（需要登录）
		orderRoutes := protected.Group("/orders")
		{
			// 订单基础操作
			orderRoutes.POST("", controllers.CreateOrder)              // 创建订单（玩家）
			orderRoutes.POST("/accept", controllers.AcceptOrder)       // 接受订单（代练）
			orderRoutes.POST("/confirm", controllers.ConfirmOrder)     // 确认订单并支付剩余金额（玩家）
			orderRoutes.POST("/cancel", controllers.CancelOrder)       // 取消订单（玩家/代练）
			orderRoutes.POST("/complete", controllers.CompleteOrder)   // 手动完成订单（代练）
			orderRoutes.PUT("/:id/dispute", controllers.CreateDispute) // 创建争议（玩家/代练）

			// 订单状态查询
			orderRoutes.GET("/my", controllers.GetMyOrders)               // 获取我的订单（我发布的+我接取的）
			orderRoutes.GET("/available", controllers.GetAvailableOrders) // 获取可接单的订单
			orderRoutes.GET("/:id/logs", controllers.GetOrderLogs)        // 获取订单操作日志
		}

		// 订单相关（需要登录）
		riotRoutes := protected.Group("/riot")
		{
			handler := &controllers.GameApiHandler{}
			riotRoutes.GET("/getSummonerPUUID", handler.GetSummonerPUUID)
			riotRoutes.GET("/getLeagueEntries", handler.GetLeagueEntries)
			riotRoutes.GET("/getWithRank", handler.GetSummonerProfileWithRank)
		}

	}

	return r
}

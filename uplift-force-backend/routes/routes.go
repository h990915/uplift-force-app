package routes

import (
	"github.com/gin-gonic/gin"
	"uplift-force-backend/controllers"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		api.GET("/users", controllers.GetUsers)
		api.POST("/users", controllers.CreateUser)
	}

	return r
}

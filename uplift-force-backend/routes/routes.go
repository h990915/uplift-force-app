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
		api.POST("/createUsers", controllers.CreateUser)
		api.GET("/getUserByID", controllers.GetUserByID)
		api.GET("getUserByWallet", controllers.GetUserByWallet)
		api.POST("/updateUser", controllers.UpdateUser)
		api.POST("/deleteUser", controllers.DeleteUser)
	}

	return r
}

package router

import (
	"biller/controllers"
	"biller/database"
	"biller/repositories"
	"github.com/gin-gonic/gin"
)

func InitAuthRoute(app *gin.Engine) {
	userController := controllers.InitUserController(repositories.InitUserRepository(database.Get()))
	api := app.Group("/api/")
	{
		api.POST("refresh", controllers.RefreshToken)
		api.POST("login", controllers.Login)
		api.POST("user", userController.Create())
	}
}

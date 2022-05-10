package router

import (
	"biller/repositories"
	"biller/services"
	"github.com/gin-gonic/gin"
)

func InitUserRoute(app *gin.Engine, userRepo *repositories.UserRepository) {
	userController := services.InitUserController(userRepo)
	api := app.Group("/api/")
	{
		api.GET("/user", userController.Get())
		api.POST("/user", userController.Create())
	}
}

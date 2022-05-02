package router

import (
	"biller/repositories"
	"biller/services"
	"github.com/gin-gonic/gin"
)

func InitAuthRoute(app *gin.Engine, userRepo *repositories.UserRepository) {
	api := app.Group("/api/")
	{
		api.POST("login", services.Login(userRepo))
	}
}

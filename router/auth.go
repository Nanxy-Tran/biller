package router

import (
	"biller/middleware"
	"biller/repositories"
	"github.com/gin-gonic/gin"
)

func InitAuthRoute(app *gin.Engine, userRepo *repositories.UserRepository) {
	api := app.Group("/api/")
	{
		api.POST("login", middleware.Login(userRepo))
	}
}

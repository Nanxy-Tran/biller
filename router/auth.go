package router

import (
	"biller/database"
	"biller/repositories"
	"biller/services"
)

func InitAuthRoute(app *database.InjectDBApp) {
	userRepo := repositories.InitUserRepository(app.DB)
	userController := services.InitUserController(userRepo)
	api := app.Instance.Group("/api/")
	{
		api.POST("refresh", services.RefreshToken())
		api.POST("login", services.Login(userRepo))
		api.POST("user", userController.Create())
	}
}

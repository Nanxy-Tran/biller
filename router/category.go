package router

import (
	"biller/database"
	"biller/repositories"
	"biller/services"
	"github.com/gin-gonic/gin"
)

func InitCategoryRoute(app *database.InjectDBApp, middlewares ...gin.HandlerFunc) {
	billRepo := repositories.InitCategoryRepository(app.DB)
	categoryController := services.InitCategoryController(billRepo)

	api := app.Instance.Group("/api/", middlewares...)
	{
		api.GET("/category", categoryController.GetCategories())
		api.POST("/category", categoryController.CreateCategory())
	}
}

package router

import (
	"biller/controllers"
	"biller/database"
	"biller/repositories"
	"github.com/gin-gonic/gin"
)

func InitCategoryRoute(app *gin.Engine) {
	billRepo := repositories.InitCategoryRepository(database.Get())
	categoryController := controllers.InitCategoryController(billRepo)

	api := app.Group("/api/")
	{
		api.GET("/category", categoryController.GetCategories)
		api.POST("/category", categoryController.CreateCategory)
	}
}

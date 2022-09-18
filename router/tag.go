package router

import (
	"biller/controllers"
	"biller/database"
	"biller/repositories"
	"github.com/gin-gonic/gin"
)

func InitTagRoute(app *gin.Engine) {
	tagRepo := repositories.InitTagRepository(database.Get())
	tagController := controllers.InitTagController(tagRepo)
	api := app.Group("/api/")
	{
		api.GET("/tags", tagController.GetTags)
		api.POST("/tag", tagController.CreatTag)
	}
}

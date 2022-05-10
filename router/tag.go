package router

import (
	"biller/repositories"
	"biller/services"
	"github.com/gin-gonic/gin"
)

func InitTagRoutes(app *gin.Engine, tag *repositories.TagRepository) *gin.Engine {
	tagController := services.InitTagController(tag)
	api := app.Group("/api/")
	{
		api.GET("/tags", tagController.GetTags())
		api.POST("/tag", tagController.CreatTag())
	}
	return app
}

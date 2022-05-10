package router

import (
	"biller/repositories"
	"biller/services"
	"github.com/gin-gonic/gin"
)

func InitTagRoute(app *gin.Engine, tag *repositories.TagRepository, middleware gin.HandlerFunc) {
	tagController := services.InitTagController(tag)
	api := app.Group("/api/").Use(middleware)
	{
		api.GET("/tags", tagController.GetTags())
		api.POST("/tag", tagController.CreatTag())
	}
}

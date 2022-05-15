package router

import (
	"biller/database"
	"biller/repositories"
	"biller/services"
	"github.com/gin-gonic/gin"
)

func InitTagRoute(app *database.InjectDBApp, middleware ...gin.HandlerFunc) {
	tagRepo := repositories.InitTagRepository(app.DB)
	tagController := services.InitTagController(tagRepo)
	api := app.Instance.Group("/api/", middleware...)
	{
		api.GET("/tags", tagController.GetTags())
		api.POST("/tag", tagController.CreatTag())
	}
}

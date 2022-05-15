package router

import (
	"biller/repositories"
	"biller/services"
	"github.com/gin-gonic/gin"
)

func InitBillRoute(app *gin.Engine) {
	billRepo := repositories.InitBillRepository(db)
	billController := services.InitBillController(billRepo)
	api := app.Group("/api/").Use(middlewares)
	{
		api.GET("/bills", billController.GetBills())
		api.GET("bill/:id", billController.GetBill())
		api.POST("/bill", billController.Save())
	}
}

func InitPageApp(app *gin.Engine) {
	app.Static("/resource", "./resource")
	app.StaticFile("/app.js", "./client/app.js")
	//app.StaticFile("/", "./client/base.html")
}

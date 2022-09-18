package router

import (
	"biller/controllers"
	"biller/database"
	"biller/repositories"
	"github.com/gin-gonic/gin"
)

func InitBillRoute(app *gin.Engine, middlewares ...gin.HandlerFunc) {
	billRepo := repositories.InitBillRepository(database.Get())
	billController := controllers.InitBillController(billRepo)

	api := app.Group("/api/", middlewares...).Use(middlewares...)
	{
		api.GET("/bills", billController.GetBills)
		api.GET("bill/:id", billController.GetBill)
		api.POST("/bill", billController.Save)
	}
}

func InitPageApp(app *gin.Engine) {
	app.Static("/dist", "./client/dist")
	app.StaticFile("/", "./client/public/index.html")
}

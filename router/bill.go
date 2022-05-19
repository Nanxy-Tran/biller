package router

import (
	"biller/database"
	"biller/repositories"
	"biller/services"
	"github.com/gin-gonic/gin"
)

func InitBillRoute(app *database.InjectDBApp, middlewares ...gin.HandlerFunc) {
	billRepo := repositories.InitBillRepository(app.DB)
	billController := services.InitBillController(billRepo)

	api := app.Instance.Group("/api/", middlewares...)
	{
		api.GET("/bills", billController.GetBills())
		api.GET("bill/:id", billController.GetBill())
		api.POST("/bill", billController.Save())
	}
}

func InitPageApp(app *gin.Engine) {
	app.Static("/dist", "./client/dist")
	app.StaticFile("/", "./client/public/index.html")
}

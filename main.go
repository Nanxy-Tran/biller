package main

import (
	"biller/database"
	"biller/middleware"
	"biller/repositories"
	"biller/router"
	"github.com/gin-gonic/gin"
)

var emptyMiddleware = func(context *gin.Context) {
	context.Next()
}

func main() {
	app := gin.Default()
	app.Use(middleware.CORSMiddleware)

	userName, password, dbName := "root", "admin", "biller"

	db := database.ConnectDB(userName, password, dbName)

	userRepo := repositories.InitUserRepository(db)
	billRepo := repositories.InitBillRepository(db)

	router.InitPageApp(app)
	router.InitBillRoute(app, billRepo)
	router.InitUserRoute(app, userRepo, emptyMiddleware)
	router.InitAuthRoute(app, userRepo)

	err := app.Run("localhost:8080")
	if err != nil {
		panic(err.Error())
	}
}

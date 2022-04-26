package main

import (
	"biller/database"
	"biller/middleware"
	"biller/repositories"
	"biller/router"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.Use(middleware.CORSMiddleware)

	userName, password, dbName := "root", "admin", "biller"

	db := database.ConnectDB(userName, password, dbName)

	billRepository := repositories.InitBillRepository(db)
	userRepository := repositories.InitUserRepository(db)

	router.InitPageApp(app)
	router.InitBillRoute(app, billRepository)
	router.InitUserRoute(app, userRepository)
	router.InitAuthRoute(app, userRepository)

	err := app.Run("localhost:8080")
	if err != nil {
		panic(err.Error())
	}
}

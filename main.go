package main

import (
	"biller/database"
	"biller/repositories"
	"biller/router"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	userName, password, dbName := "root", "admin", "biller"

	db := database.ConnectDB(userName, password, dbName)

	billRepository := repositories.InitBillRepository(db)
	router.InitBillRoute(app, billRepository)

	err := app.Run("localhost:8080")
	if err != nil {
		panic(err.Error())
	}
}

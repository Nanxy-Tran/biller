package main

import (
	"biller/database"
	"biller/models"
	"biller/repositories"
	"biller/router"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	userName, password, dbName := "root", "admin", "biller"

	db := database.ConnectDB(userName, password, dbName)
	err := db.AutoMigrate(&models.Bill{})

	if err != nil {
		panic(err.Error())
	}

	billRepository := repositories.BillRepository{DB: db}
	router.InitBillRoute(app, billRepository)

	err = app.Run("localhost:8080")
	if err != nil {
		panic(err.Error())
	}
}

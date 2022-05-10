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

	userRepo := repositories.InitUserRepository(db)
	billRepo := repositories.InitBillRepository(db)
	tagRepo := repositories.InitTagRepository(db)

	router.InitPageApp(app)
	router.InitBillRoute(app, billRepo).Use(middleware.AuthMiddleware(db))
	router.InitUserRoute(app, userRepo).Use(middleware.AuthMiddleware(db))
	router.InitTagRoutes(app, tagRepo).Use(middleware.AuthMiddleware(db))
	router.InitAuthRoute(app, userRepo)

	err := app.Run("localhost:8080")
	if err != nil {
		panic(err.Error())
	}
}

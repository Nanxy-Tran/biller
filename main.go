package main

import (
	"biller/database"
	"biller/middleware"
	"biller/router"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	app.Use(middleware.CORSMiddleware)

	userName, password, dbName := "root", "admin", "biller"

	db := database.ConnectDB(userName, password, dbName)
	injectedApp := database.InjectDB(app, db)

	//router.InitPageApp(app)
	router.InitAuthRoute(injectedApp)

	injectedApp.Instance.Use(middleware.Authentication)
	injectedApp.Instance.Use(middleware.Authorization(injectedApp.DB))

	router.InitBillRoute(injectedApp)
	router.InitTagRoute(injectedApp)

	err := app.Run("localhost:8080")
	if err != nil {
		panic(err.Error())
	}
}

/**
1. generic error
2. utils helper
3. tag tag tag & bill bill bill
4. Feature honda dream
*/

package main

import (
	"biller/database"
	"biller/middleware"
	"biller/router"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	app := gin.Default()
	env := os.Getenv("mode")
	fmt.Println(env)
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
	router.InitCategoryRoute(injectedApp)

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

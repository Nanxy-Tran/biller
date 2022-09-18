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
	database.Connect()
	app := gin.Default()
	env := os.Getenv("mode")
	fmt.Println(env)
	app.Use(middleware.CORSMiddleware)
	router.SetupRoutes(app)

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

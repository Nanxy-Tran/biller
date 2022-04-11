package main

import "github.com/gin-gonic/gin"

func main() {
	app := gin.Default()

	app.GET("", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "Hello"})
	})

	err := app.Run("localhost:8080")
	if err != nil {
		panic(err.Error())
	}
}

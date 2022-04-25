package router

import (
	"biller/models"
	"biller/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUserRoute(app *gin.Engine, repo *repositories.UserRepository) {
	api := app.Group("/api/")
	{
		//api.GET("/user", func(context *gin.Context) {
		//	bills := repo.GetBills()
		//	context.JSON(http.StatusOK, gin.H{
		//		"data": bills.Result,
		//	})
		//})

		api.POST("/user", func(context *gin.Context) {
			var user models.User
			if err := context.ShouldBindJSON(&user); err == nil {
				fmt.Printf("user email: %s", user.Email)
			} else {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			result := repo.Creat(&user)

			if result.Error != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			} else {
				context.JSON(http.StatusCreated, gin.H{"user_id": result.Result})
			}
		})
	}

}

package router

import (
	"biller/models"
	"biller/repositories"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func InitRouter(g *gin.Context) {
//	g.Get("/")
//}

func InitBillRoute(app *gin.Engine, repo *repositories.BillRepository) {
	app.GET("/bills", func(context *gin.Context) {
		bills := repo.GetBills()
		context.JSON(http.StatusOK, gin.H{
			"data": bills.Result,
		})
	})

	app.GET("bill/:id", func(context *gin.Context) {
		id := context.Param("id")
		bill := repo.GetBill(id)
		context.JSON(http.StatusOK, gin.H{"data": bill.Result})
		return
	})

	app.POST("/bill", func(context *gin.Context) {
		var bill models.Bill
		if err := context.ShouldBindJSON(&bill); err == nil {
			fmt.Printf("Bill: name -> %s, amount -> %d", bill.Name, bill.Amount)
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := repo.Save(&bill)

		if result.Error != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		} else {
			context.JSON(http.StatusCreated, gin.H{"data": result.Result})
		}
	})
}

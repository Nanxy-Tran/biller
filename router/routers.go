package router

import (
	"biller/models"
	"biller/repositories"
	"biller/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func InitRouter(g *gin.Context) {
//	g.Get("/")
//}

func InitBillRoute(app *gin.Engine, repo repositories.BillRepository) {
	app.GET("/bills", func(context *gin.Context) {
		bills := services.GetBills(repo)
		context.JSON(http.StatusOK, gin.H{
			"data": bills.Result,
		})
	})

	app.POST("/bill", func(context *gin.Context) {
		var bill models.Bill
		if err := context.ShouldBindJSON(&bill); err == nil {
			fmt.Printf("Bill: name -> %s, amount -> %d", bill.Name, bill.Amount)
		} else {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := services.CreateBill(&bill, repo)

		if result.Error != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		} else {
			context.JSON(http.StatusCreated, gin.H{"data": bill})
		}
	})
}

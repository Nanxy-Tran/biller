package router

import (
	"biller/models"
	"biller/repositories"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func InitRouter(g *gin.Context) {
//	g.Get("/")
//}

func InitBillRoute(app *gin.Engine, repo *repositories.BillRepository) {
	api := app.Group("/api/")
	{
		api.GET("/bills", func(context *gin.Context) {
			bills := repo.GetBills()
			context.JSON(http.StatusOK, gin.H{
				"data": bills.Result,
			})
		})

		api.GET("bill/:id", func(context *gin.Context) {
			id := context.Param("id")
			bill := repo.GetBill(id)
			if bill.Result != nil {
				context.JSON(http.StatusOK, gin.H{"data": bill.Result})
			} else if bill.Error == sql.ErrNoRows {
				context.JSON(http.StatusNotFound, gin.H{"error": "No bill found"})
			} else {
				context.JSON(http.StatusBadRequest, gin.H{"error": bill.Error.Error()})
			}
		})

		api.POST("/bill", func(context *gin.Context) {
			var bill models.Bill
			if err := context.ShouldBindJSON(&bill); err == nil {
				fmt.Printf("Bill: name -> %s, amount -> %d", bill.Name, bill.Amount)
			} else {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			result := repo.Save(&bill)

			if result.Error != nil {
				context.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			} else {
				context.JSON(http.StatusCreated, gin.H{"bill_id": result.Result})
			}
		})
	}

}

func InitBillPages(app *gin.Engine, repo *repositories.BillRepository) {
	app.LoadHTMLGlob("templates/*.tmpl")
	app.GET("/bills", func(context *gin.Context) {
		bills := repo.GetBills()
		context.HTML(http.StatusOK, "base.tmpl", gin.H{
			"bills": bills.Result,
		})
	})

	//app.GET("bill/:id", func(context *gin.Context) {
	//	id := context.Param("id")
	//	bill := repo.GetBill(id)
	//	if bill.Result != nil {
	//		context.JSON(http.StatusOK, gin.H{"data": bill.Result})
	//	} else if bill.Error == sql.ErrNoRows {
	//		context.JSON(http.StatusNotFound, gin.H{"error": "No bill found"})
	//	} else {
	//		context.JSON(http.StatusBadRequest, gin.H{"error": bill.Error.Error()})
	//	}
	//})
	//
	//app.POST("/bill", func(context *gin.Context) {
	//	var bill models.Bill
	//	if err := context.ShouldBindJSON(&bill); err == nil {
	//		fmt.Printf("Bill: name -> %s, amount -> %d", bill.Name, bill.Amount)
	//	} else {
	//		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//		return
	//	}
	//
	//	result := repo.Save(&bill)
	//
	//	if result.Error != nil {
	//		context.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
	//	} else {
	//		context.JSON(http.StatusCreated, gin.H{"bill_id": result.Result})
	//	}
	//})
}

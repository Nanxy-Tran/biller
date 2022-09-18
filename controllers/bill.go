package controllers

import (
	"biller/models"
	"biller/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BillController struct {
	repository *repositories.BillRepository
}

func InitBillController(repo *repositories.BillRepository) *BillController {
	return &BillController{
		repository: repo,
	}
}

func (controller *BillController) Save(context *gin.Context) {
	bill := models.Bill{}
	if err := context.ShouldBindJSON(&bill); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdBill := controller.repository.Save(&bill)

	if createdBill.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": createdBill.Error.Error()})
	} else {
		context.JSON(http.StatusCreated, gin.H{"bill_id": createdBill.Result})
	}
}

func (controller *BillController) GetBills(context *gin.Context) {
	billsOptions := repositories.BillsOptions{Limit: 10, CurrentPage: 1}

	if err := context.ShouldBindQuery(&billsOptions); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Something wrong with pages !"})
		return
	}

	bills := controller.repository.GetBills(&billsOptions)
	context.JSON(http.StatusOK, bills)
}

func (controller *BillController) GetBill(context *gin.Context) {
	billID := context.Param("id")

	if billID == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
		return
	}

	billResult := controller.repository.GetBill(billID)

	if billResult.Error != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": billResult.Error.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": billResult.Result})
}

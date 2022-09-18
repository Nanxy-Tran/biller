package controllers

import (
	"biller/models"
	"biller/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	repository *repositories.UserRepository
}

func InitUserController(repo *repositories.UserRepository) *UserController {
	return &UserController{
		repository: repo,
	}
}

func (controller *UserController) Get() gin.HandlerFunc {
	return func(context *gin.Context) {
		var user models.User
		if err := context.ShouldBind(&user); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user credentials"})
			return
		}

		result := controller.repository.Get(user.Email)
		if result.Error != nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": result.Result})
	}
}

func (controller *UserController) Create() gin.HandlerFunc {
	return func(context *gin.Context) {
		var user models.User
		if err := context.ShouldBindJSON(&user); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := controller.repository.Creat(&user)

		if result.Error != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
			return
		}
		context.JSON(http.StatusCreated, gin.H{"user_id": result.Result})
	}
}

//func (controller *UserController) Update() gin.HandlerFunc {
//	return func(context *gin.Context) {
//		bill := models.Bill{}
//		if err := context.ShouldBindQuery(&bill); err != nil {
//			context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid params"})
//			return
//		}
//
//		billResult := controller.repository.GetBill(strconv.Itoa(int(bill.ID)))
//		if billResult.Error != nil {
//			context.JSON(http.StatusNotFound, gin.H{"error": "Not found any bill"})
//			return
//		}
//
//		context.JSON(http.StatusOK, gin.H{"data": billResult.Result})
//
//	}
//}

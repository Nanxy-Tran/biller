package services

import (
	"biller/models"
	"biller/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryController struct {
	repository *repositories.CategoryRepository
}

func InitCategoryController(repo *repositories.CategoryRepository) *CategoryController {
	return &CategoryController{
		repository: repo,
	}
}

func (controller *CategoryController) GetCategories() gin.HandlerFunc {
	return func(context *gin.Context) {
		result := controller.repository.GetCategories()

		if result.Error != nil {
			context.JSON(http.StatusBadGateway, gin.H{"error": "No categories found"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": result.Result})
	}
}

func (controller *CategoryController) CreateCategory() gin.HandlerFunc {
	return func(context *gin.Context) {
		var category models.Category
		if err := context.ShouldBindJSON(&category); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := controller.repository.CreateCategory(&category)

		if result.Error != nil {
			context.JSON(http.StatusConflict, gin.H{"error": "Tag already existed !"})
			return
		}

		context.JSON(http.StatusCreated, gin.H{"tag_id": result.Result})
	}
}

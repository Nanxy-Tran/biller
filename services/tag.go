package services

import (
	"biller/models"
	"biller/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TagController struct {
	repository *repositories.TagRepository
}

func InitTagController(repo *repositories.TagRepository) *TagController {
	return &TagController{
		repository: repo,
	}
}

func (controller *TagController) GetTags() gin.HandlerFunc {
	return func(context *gin.Context) {
		//This should be a function to handle user authorization
		user, exists := context.Get("user")
		if !exists {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Internal error"})
			return
		}
		result := controller.repository.GetTags(user.(models.User).ID)
		if result.Error != nil {
			context.JSON(http.StatusBadGateway, gin.H{"error": "Can not get tags"})
			return
		}

		context.JSON(http.StatusOK, gin.H{"data": result.Result})
	}
}

func (controller *TagController) CreatTag() gin.HandlerFunc {
	return func(context *gin.Context) {
		var tag models.Tag
		if err := context.ShouldBindJSON(&tag); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		result := controller.repository.CreateTag(&tag)

		if result.Error != nil {
			context.JSON(http.StatusConflict, gin.H{"error": "Tag already existed !"})
			return
		}

		context.JSON(http.StatusCreated, gin.H{"tag_id": result.Result})
	}
}

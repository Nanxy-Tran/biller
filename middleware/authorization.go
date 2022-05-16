package middleware

import (
	"biller/models"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Authorization(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		if claims, exist := context.Get("claims"); exist {
			var user models.User
			db.First(&user).Where("email = ?", claims.(Claims).Email)
			context.Set("user", user)
			context.Next()
		} else {
			context.AbortWithError(http.StatusInternalServerError, errors.New("claims not found"))
			return
		}
	}
}

package middleware

import "github.com/gin-gonic/gin"

type Login struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Auth(app *gin.Engine) {

}

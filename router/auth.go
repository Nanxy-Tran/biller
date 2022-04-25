package router

import "github.com/gin-gonic/gin"

func InitAuthRoute(app *gin.Engine) {
	app.POST("login")
}

func login(ctx *gin.Context) {

}

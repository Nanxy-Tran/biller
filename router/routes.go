package router

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine) {
	InitAuthRoute(app)
	InitBillRoute(app)
	InitCategoryRoute(app)
	InitTagRoute(app)
	InitPageApp(app)
}

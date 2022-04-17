package middleware

import (
	"github.com/gin-contrib/cors"
	"time"
)

var CORSMiddleware = cors.New(cors.Config{
	AllowOrigins:     []string{"http://localhost:3000"},
	AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE"},
	AllowHeaders:     []string{"Origin"},
	ExposeHeaders:    []string{"Content-Length"},
	AllowCredentials: true,
	MaxAge:           12 * time.Hour,
})

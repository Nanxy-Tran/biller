package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
)

var JwtSecret = []byte("0A5456A8E91294BB5664BF0F2B08A016D70E88D8A226E2E828AAC175927EF9F2")

type Credentials struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	Password []byte `json:"password"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

type AuthHeader struct {
	TokenID string `header:"Authorization"`
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var header AuthHeader

		if err := c.ShouldBindHeader(&header); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid token"})
			c.Abort()
		}

		tokenString := strings.Split(header.TokenID, "Bearer ")
		claims := Claims{}

		fmt.Println(tokenString)

		token, err := jwt.ParseWithClaims(tokenString[1], &claims, func(token *jwt.Token) (interface{}, error) {
			return JwtSecret, nil
		})

		if err != nil {
			err = c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		if !token.Valid {
			fmt.Println(err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			c.Abort()
			return
		}
		c.Next()
	}

}

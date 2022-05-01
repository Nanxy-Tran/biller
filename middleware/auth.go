package middleware

import (
	"biller/models"
	"biller/repositories"
	"biller/services"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

var jwtSecret = []byte("0A5456A8E91294BB5664BF0F2B08A016D70E88D8A226E2E828AAC175927EF9F2")

type Credentials struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Claims struct {
	Password []byte `json:"password"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func Login(r *repositories.UserRepository) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		var credentials Claims
		err := json.NewDecoder(ctx.Request.Body).Decode(&credentials)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid credentials"})
			return
		}

		var authUser = r.Get(credentials.Email)
		fmt.Println(authUser)

		if authUser.Error != nil || credentials.Valid() != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized action"})
			return
		}

		if err := services.IsPasswordMatched(credentials.Password, authUser.Result.(*models.User).Password); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user credentials"})
			return
		}

		expirationTime := time.Now().Add(10 * time.Minute)

		claim := &Claims{
			Email:    credentials.Email,
			Password: credentials.Password,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: &jwt.NumericDate{Time: expirationTime},
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

		tokenString, err := token.SignedString(jwtSecret)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something conflict"})
			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{"token": tokenString})
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return token, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Action unauthorized"})
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.Next()
	}

}

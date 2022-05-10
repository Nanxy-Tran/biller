package services

import (
	"biller/middleware"
	"biller/models"
	"biller/repositories"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strings"
	"time"
)

func IsPasswordMatched(comingPassword []byte, hashedPassword []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, comingPassword)
}

func Login(r *repositories.UserRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var credentials middleware.Claims
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

		if err := IsPasswordMatched(credentials.Password, authUser.Result.(*models.User).Password); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user credentials"})
			return
		}

		expirationTime := time.Now().Add(3 * time.Hour)

		claim := &middleware.Claims{
			Email:    credentials.Email,
			Password: credentials.Password,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: &jwt.NumericDate{Time: expirationTime},
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

		tokenString, err := token.SignedString(middleware.JwtSecret)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Something conflict"})
			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{"token": tokenString, "username": authUser.Result.(*models.User).UserName})
	}
}

func RefreshToken() gin.HandlerFunc {
	return func(context *gin.Context) {
		var header middleware.AuthHeader

		if err := context.ShouldBindHeader(&header); err != nil {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid token"})
			context.Abort()
			return
		}

		tokenString := strings.Split(header.TokenID, "Bearer ")

		if len(tokenString) < 2 {
			context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid header params"})
			context.Abort()
			return
		}

		claims := middleware.Claims{}

		_, err := jwt.ParseWithClaims(tokenString[1], &claims, func(token *jwt.Token) (interface{}, error) {
			return middleware.JwtSecret, nil
		})

		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "invalid params"})
			context.Abort()
			return
		}

		expirationTime := time.Now().Add(3 * time.Hour)

		newClaims := &middleware.Claims{
			Email:    claims.Email,
			Password: claims.Password,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: &jwt.NumericDate{Time: expirationTime},
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)

		newTokenString, err := token.SignedString(middleware.JwtSecret)

		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Something wrong happened"})
			context.Abort()
			return
		}

		context.JSON(http.StatusOK, gin.H{"token": newTokenString})

	}
}

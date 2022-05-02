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
	"time"
)

func IsPasswordMatched(comingPassword []byte, hashedPassword []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, comingPassword)
}

func Login(r *repositories.UserRepository) func(ctx *gin.Context) {
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

package repositories

import (
	"biller/models"
	"biller/utils"
	"errors"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func InitUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Get(email string) RepositoryResult {
	result := r.DB.First(&models.User{}, "email = ?", email)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return RepositoryResult{Error: &ApiError{e: "No user found for this email"}}
	}
	return RepositoryResult{Result: result}
}

func (r *UserRepository) Creat(user *models.User) RepositoryResult {
	password, err := utils.HashPassword(user.Password)
	if err != nil {
		return RepositoryResult{Error: &ApiError{e: "Something wrong with password"}}
	}

	user.Password = password
	result := r.DB.Create(&user)

	if result.Error != nil {
		return RepositoryResult{Error: &ApiError{e: "User already existed"}}
	}

	return RepositoryResult{Result: user.ID}
}

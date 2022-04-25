package repositories

import (
	"biller/models"
	"biller/services"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func InitUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Get(email string) RepositoryResult {
	var user models.User
	row := r.DB.QueryRow("SELECT 'user_name', 'password' from users where email = ?", email)

	err := row.Scan(&user.UserName, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		return RepositoryResult{Error: &ApiError{e: "No user found for this email"}}
	}
	return RepositoryResult{Result: user}
}

func (r *UserRepository) Creat(user *models.User) RepositoryResult {
	password, err := services.HashPassword(user.Password)
	if err != nil {
		return RepositoryResult{Error: &ApiError{e: "Something wrong with password"}}
	}
	result, err := r.DB.Exec("INSERT INTO users (user_name, email, password) VALUE (?, ?, ?)", user.UserName, user.Email, password)

	if err != nil {
		return RepositoryResult{Error: &ApiError{e: "User already existed"}}
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return RepositoryResult{Error: &ApiError{e: "Cannot create user"}}
	}

	return RepositoryResult{Result: userId}
}

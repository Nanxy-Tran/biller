package repositories

import (
	"biller/models"
	"biller/utils"
	"database/sql"
	"fmt"
)

type UserRepository struct {
	DB *sql.DB
}

func InitUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) getUserStatement(field string) (*sql.Stmt, error) {
	return r.DB.Prepare(fmt.Sprintf("SELECT %s FROM users where email = ?", field))
}

func (r *UserRepository) Get(email string) RepositoryResult {
	var user = new(models.User)
	statement, _ := r.getUserStatement("username, password")
	row := statement.QueryRow(email)

	err := row.Scan(&user.UserName, &user.Password)

	if err == sql.ErrNoRows {
		return RepositoryResult{Error: &ApiError{e: "No user found for this email"}}
	}
	return RepositoryResult{Result: user}
}

func (r *UserRepository) Creat(user *models.User) RepositoryResult {
	password, err := utils.HashPassword(user.Password)
	if err != nil {
		return RepositoryResult{Error: &ApiError{e: "Something wrong with password"}}
	}
	result, err := r.DB.Exec(
		"INSERT INTO users (username, email, password) VALUE (?, ?, ?)",
		user.UserName, user.Email, password,
	)

	if err != nil {
		return RepositoryResult{Error: &ApiError{e: "User already existed"}}
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return RepositoryResult{Error: &ApiError{e: "Cannot create user"}}
	}

	return RepositoryResult{Result: userId}
}

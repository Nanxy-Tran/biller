package database

import (
	"biller/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func Connect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error when parsing env: ", err.Error())
	}

	userName := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE")

	dbDNS := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, dbName)

	db, err := gorm.Open(mysql.Open(dbDNS), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting database: ", err.Error())
	}

	//TODO: replace auto migrate
	err = db.AutoMigrate(
		&models.User{},
		&models.Bill{},
		&models.Tag{},
		&models.Category{},
	)

	if err != nil {
		panic(err.Error())
	}

	return db
}

var db *gorm.DB

func Get() *gorm.DB {
	if db != nil {
		return db
	} else {
		db = Connect()
		return db
	}
}

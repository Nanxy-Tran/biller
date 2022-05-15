package database

import (
	"biller/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(userName string, password string, dbName string) (db *gorm.DB) {
	dbDNS := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", userName, password, dbName)
	db, err := gorm.Open(mysql.Open(dbDNS), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

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

type InjectDBApp struct {
	Instance *gin.Engine
	DB       *gorm.DB
}

func InjectDB(app *gin.Engine, db *gorm.DB) *InjectDBApp {
	return &InjectDBApp{
		Instance: app,
		DB:       db,
	}
}

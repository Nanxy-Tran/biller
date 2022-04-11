package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB(userName string, password string, dbName string) (db *gorm.DB) {
	DSN := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8&parseTime=true&loc=Local", userName, password, dbName)
	if db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{}); err != nil {
		panic(err.Error())
	} else {
		return db
	}
}

func MigrateDB(db *gorm.DB) {
	if err := db.AutoMigrate(); err != nil {
		panic(err.Error())
	}
}

package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	DSN := "root:admin@tcp(127.0.0.1:3306/biller?charset=utf8&parseTime=true&loc=Local)"
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

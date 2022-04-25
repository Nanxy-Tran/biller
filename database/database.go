package database

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

func ConnectDB(userName string, password string, dbName string) (db *sql.DB) {

	dbConfig := mysql.Config{
		User:      userName,
		Passwd:    password,
		Net:       "tcp",
		Addr:      "127.0.0.1:3306",
		DBName:    dbName,
		ParseTime: true,
	}

	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		panic(err.Error())
	}
	return db
}

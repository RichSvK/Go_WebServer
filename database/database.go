package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var PoolDB *sql.DB = nil

func GetConnection(username string, password string, databaseName string, path string) {
	// Open connection
	PoolDB, err := sql.Open("mysql", username+":"+password+"@tcp("+path+":3306)/"+databaseName)
	if err != nil {
		log.Fatal("Failed to connect")
		return
	}

	// Check Connection
	err = PoolDB.Ping()
	if err != nil {
		log.Fatal("Inconnect credential")
		return
	}

	PoolDB.SetMaxIdleConns(5)
	PoolDB.SetMaxOpenConns(10)
	PoolDB.SetConnMaxIdleTime(3 * time.Minute)
	PoolDB.SetConnMaxLifetime(2 * time.Hour)
}

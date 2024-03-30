package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetConnection(username string, password string, databaseName string, path string) *sql.DB {
	// Open connection
	PoolDB, err := sql.Open("mysql", username+":"+password+"@tcp("+path+":3306)/"+databaseName)
	if err != nil {
		fmt.Println("Failed to connect")
		return nil
	}

	// Check Connection
	err = PoolDB.Ping()
	if err != nil {
		fmt.Println("Inconnect credential")
		// log.Fatal("Inconnect credential")
		return nil
	}

	PoolDB.SetMaxIdleConns(5)
	PoolDB.SetMaxOpenConns(10)
	PoolDB.SetConnMaxIdleTime(3 * time.Minute)
	PoolDB.SetConnMaxLifetime(2 * time.Hour)
	fmt.Println("Success make connection")
	return PoolDB
}

package database

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/RichSvK/Go_WebServer/models"
	_ "github.com/go-sql-driver/mysql"
)

var PoolDB *sql.DB = nil

func GetConnection(username string, password string, host string, databaseName string) error {
	// Open connection
	source := username + ":" + password + "@tcp(" + host + ":3306)/" + databaseName
	var err error = nil
	PoolDB, err = sql.Open("mysql", source)
	if err != nil {
		fmt.Println("Failed to connect")
		return err
	}

	// Check Connection
	err = PoolDB.Ping()
	if err != nil {
		fmt.Println("Incorrect credential")
		return err
	}

	PoolDB.SetMaxIdleConns(5)
	PoolDB.SetMaxOpenConns(10)
	PoolDB.SetConnMaxIdleTime(5 * time.Minute)
	PoolDB.SetConnMaxLifetime(1 * time.Hour)
	fmt.Println("Success make connection")
	return nil
}

func GetStudent(NIM string, student *models.Student) int {
	query := "SELECT * FROM Students WHERE NIM = ?"
	ctx := context.Background()
	statement, err := PoolDB.PrepareContext(ctx, query)
	if err != nil {
		fmt.Println("error 0")
		fmt.Println(err)
		return http.StatusInternalServerError
	}
	defer statement.Close()

	result, err := statement.QueryContext(ctx, NIM)
	if err != nil {
		fmt.Println("error 1")
		return http.StatusInternalServerError
	}

	if !result.Next() {
		fmt.Println("error 2")
		return http.StatusNotFound
	}
	defer result.Close()

	err = result.Scan(&student.NIM, &student.Name, &student.Age)
	if err != nil {
		fmt.Println("error 3")
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

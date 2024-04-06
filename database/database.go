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

	PoolDB.SetMaxIdleConns(3)
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
		return http.StatusInternalServerError
	}
	defer statement.Close()

	result, err := statement.QueryContext(ctx, NIM)
	if err != nil {
		return http.StatusInternalServerError
	}

	if !result.Next() {
		return http.StatusNotFound
	}
	defer result.Close()

	err = result.Scan(&student.NIM, &student.Name, &student.Age)
	if err != nil {
		return http.StatusInternalServerError
	}
	return http.StatusOK
}

func GetStudents() ([]models.Student, int) {
	query := "SELECT * FROM Students"
	ctx := context.Background()
	statement, err := PoolDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	defer statement.Close()

	result, err := statement.QueryContext(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	defer result.Close()

	if !result.Next() {
		return nil, http.StatusNotFound
	}

	listStudent := []models.Student{}
	student := models.Student{}
	for {
		err = result.Scan(&student.NIM, &student.Name, &student.Age)
		if err != nil {
			return nil, http.StatusInternalServerError
		}
		listStudent = append(listStudent, student)
		if !result.Next() {
			break
		}
	}

	return listStudent, http.StatusOK
}

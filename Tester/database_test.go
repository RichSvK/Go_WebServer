package Tester

import (
	"context"
	"fmt"
	"testing"

	"github.com/RichSvK/Go_WebServer/database"
	"github.com/RichSvK/Go_WebServer/models"
)

func TestDatabase(t *testing.T) {
	db := database.GetConnection("root", "12345678", "Testing", "database-rds.cbyi6oqugc5k.us-east-1.rds.amazonaws.com")
	query := "SELECT * FROM Students"
	ctx := context.Background()
	statement, err := db.PrepareContext(ctx, query)
	fmt.Println("Test")
	if err != nil {
		fmt.Println("Error 1")
		t.FailNow()
	}

	result, err := statement.QueryContext(ctx)
	if err != nil {
		fmt.Println("Error 2")
		t.FailNow()
		return
	}

	mahasiswa := models.Student{}
	for result.Next() {
		err := result.Scan(&mahasiswa.NIM, &mahasiswa.Name, &mahasiswa.Age)
		if err != nil {
			fmt.Println("Error 3")
			t.FailNow()
		}
		// response := "NIM: " + mahasiswa.NIM + "\nNama: " + mahasiswa.Name + "\nAge: " + strconv.Itoa(mahasiswa.Age)
		fmt.Println(mahasiswa)
	}
}

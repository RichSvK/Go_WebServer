package Tester

import (
	"context"
	"fmt"
	"testing"

	"github.com/RichSvK/Go_WebServer/database"
	"github.com/RichSvK/Go_WebServer/models"
)

func TestDatabase(t *testing.T) {
	err := database.GetConnection("root", "12345678", "database-rds.cbyi6oqugc5k.us-east-1.rds.amazonaws.com", "Testing")
	if err != nil {
		t.FailNow()
	}
	query := "SELECT * FROM Students"
	ctx := context.Background()
	statement, err := database.PoolDB.PrepareContext(ctx, query)
	if err != nil {
		t.FailNow()
	}

	result, err := statement.QueryContext(ctx)
	if err != nil {
		t.FailNow()
		return
	}

	mahasiswa := models.Student{}
	for result.Next() {
		err := result.Scan(&mahasiswa.NIM, &mahasiswa.Name, &mahasiswa.Age)
		if err != nil {
			t.FailNow()
		}
		fmt.Println(mahasiswa)
	}
}

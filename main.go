package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/RichSvK/Go_WebServer/database"
	"github.com/RichSvK/Go_WebServer/models"
)

func GetData(w http.ResponseWriter, request *http.Request) {
	db := database.GetConnection("root", "12345678", "Testing", "database-rds.cbyi6oqugc5k.us-east-1.rds.amazonaws.com")
	NIM := request.URL.Query().Get("nim")
	query := "SELECT * FROM Students WHERE NIM = ?"
	ctx := context.Background()
	statement, err := db.PrepareContext(ctx, query)
	if err != nil {
		fmt.Fprint(w, "Error")
		return
	}

	result, err := statement.QueryContext(ctx, NIM)
	if err != nil {
		fmt.Fprint(w, "Error")
		return
	}

	mahasiswa := models.Student{}
	if !result.Next() {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Students With that NIM not found")
		return
	}

	err = result.Scan(&mahasiswa.NIM, &mahasiswa.Name, &mahasiswa.Age)
	if err != nil {
		fmt.Fprint(w, "Error")
		return
	}
	response := "NIM: " + mahasiswa.NIM + "\nNama: " + mahasiswa.Name + "\nAge: " + strconv.Itoa(mahasiswa.Age)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(response))
}

func RootHandler(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Course Distributed Cloud Computing")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", RootHandler)
	mux.HandleFunc("/studentInfo", GetData)
	webServer := http.Server{
		Addr:    "ec2-3-83-103-232.compute-1.amazonaws.com:8080",
		Handler: mux,
	}

	err := webServer.ListenAndServe()
	if err != nil {
		log.Fatal("Error")
		return
	}
}

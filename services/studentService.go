package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/RichSvK/Go_WebServer/database"
	"github.com/RichSvK/Go_WebServer/models"
	"github.com/julienschmidt/httprouter"
)

func GetStudentInfo(w http.ResponseWriter, request *http.Request, p httprouter.Params) {
	NIM := p.ByName("NIM")
	status := 0
	mahasiswa := &models.Student{}
	if len(NIM) != 10 {
		status = http.StatusBadRequest
	} else {
		status = database.GetStudent(NIM, mahasiswa)
	}

	switch status {
	case http.StatusInternalServerError:
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	case http.StatusNotFound:
		http.Error(w, fmt.Sprintf("Student with NIM %s is not found", NIM), http.StatusNotFound)
	case http.StatusOK:
		response, err := json.Marshal(mahasiswa)
		if err != nil {
			log.Fatal("Error: ", err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	case http.StatusBadRequest:
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}

func RootHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Home")
}

func GetStudents(w http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	listStudent, status := database.GetStudents()
	switch status {
	case http.StatusInternalServerError:
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	case http.StatusNotFound:
		http.Error(w, "No Student in Database", http.StatusNotFound)
	case http.StatusOK:
		response, err := json.Marshal(listStudent)
		if err != nil {
			log.Fatal("Error: ", err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func DeleteStudents(w http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	query := "TRUNCATE TABLE Students"
	ctx := context.Background()
	statement, err := database.PoolDB.PrepareContext(ctx, query)
	if err != nil {
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}

	_, err = statement.ExecContext(ctx)
	if err != nil {
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success Delete Data in Database"))
}

func PostStudents(w http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	listStudent := []models.Student{
		{Name: "Richard Sugiharto", Age: 20, NIM: "2602061561"},
		{Name: "Allen", Age: 21, NIM: "2602061562"},
		{Name: "Yonathan", Age: 22, NIM: "2602061563"},
		{Name: "Paulus", Age: 19, NIM: "2602061564"},
		{Name: "Antheo", Age: 18, NIM: "2602061565"},
	}

	query := "INSERT INTO Students (NIM, Name, Age) VALUES(?, ?, ?)"
	ctx := context.Background()
	stmt, err := database.PoolDB.PrepareContext(ctx, query)
	if err != nil {
		http.Error(w, "Failed to Insert", http.StatusInternalServerError)
		return
	}
	for _, student := range listStudent {
		_, err := stmt.ExecContext(ctx, student.NIM, student.Name, student.Age)
		if err != nil {
			http.Error(w, "Failed to Insert", http.StatusInternalServerError)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success Post to Database"))
}

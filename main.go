package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/RichSvK/Go_WebServer/database"
	"github.com/RichSvK/Go_WebServer/models"
)

func GetData(w http.ResponseWriter, request *http.Request) {
	NIM := request.URL.Query().Get("NIM")
	status := 0
	mahasiswa := &models.Student{}
	if len(NIM) != 10 {
		status = http.StatusBadRequest
	} else {
		status = database.GetStudent(NIM, mahasiswa)
	}

	switch status {
	case http.StatusInternalServerError:
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Server Error")
	case http.StatusNotFound:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Student with NIM %s not found\n", NIM)
	case http.StatusOK:
		w.WriteHeader(http.StatusOK)
		response := "NIM: " + mahasiswa.NIM + "\nNama: " + mahasiswa.Name + "\nAge: " + strconv.Itoa(mahasiswa.Age)
		w.Write([]byte(response))
	case http.StatusBadRequest:
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
	}
}

func RootHandler(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Course Distributed Cloud Computing")
}

func main() {
	// err := database.GetConnection("root", "root", "localhost", "Testing")
	// if err != nil {
	// 	fmt.Println("Failed to connect")
	// 	return
	// }

	mux := http.NewServeMux()
	mux.HandleFunc("/", RootHandler)
	mux.HandleFunc("/studentInfo/", GetData)
	err := database.GetConnection("root", "12345678", "database-rds.cbyi6oqugc5k.us-east-1.rds.amazonaws.com", "Testing")
	// err := database.GetConnection("root", "root", "localhost", "Testing")
	if err != nil {
		fmt.Println("Failed to connect")
		return
	}
	defer database.PoolDB.Close()
	webServer := http.Server{
		Addr: "ec2-34-201-134-21.compute-1.amazonaws.com:8080",
		// Addr:    "localhost:8080",
		Handler: mux,
	}

	err = webServer.ListenAndServe()
	if err != nil {
		log.Fatal("Error")
		return
	}
}

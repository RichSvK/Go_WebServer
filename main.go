package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RichSvK/Go_WebServer/database"
	"github.com/RichSvK/Go_WebServer/routes"
)

func main() {
	router := routes.SetupRouter()

	// The RDS instance endpoint may be changed
	err := database.GetConnection("root", "12345678", "database-rds.cbyi6oqugc5k.us-east-1.rds.amazonaws.com", "Testing")
	if err != nil {
		fmt.Println("Failed to connect")
		return
	}
	defer database.PoolDB.Close()

	webServer := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	err = webServer.ListenAndServe()
	if err != nil {
		log.Fatal("Error")
		return
	}
}

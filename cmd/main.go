package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang-rest/routes"
)

func main() {
	// router := mux.NewRouter()
	PORT := os.Getenv("PORT")
	HOST := os.Getenv("HOST")
	routes := routes.NewRouter()
	log.Printf("Server Starting at %s", HOST+":"+PORT)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", HOST, PORT), routes)

	if err != nil {
		log.Fatal("Error starting the server: ", err)
	}
}

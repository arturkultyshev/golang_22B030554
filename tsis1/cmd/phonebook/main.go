package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	http.Handle("/", router)
	router.HandleFunc("/", Home).Methods("GET")
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonById).Methods("GET")
	router.HandleFunc("/people/name/{name}", GetPersonByName).Methods("GET")
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", router)
}

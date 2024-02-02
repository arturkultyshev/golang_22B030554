package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"tsis1/pkg/handlers"
	_ "tsis1/pkg/handlers"
)

func main() {
	router := mux.NewRouter()

	http.Handle("/", router)
	router.HandleFunc("/", handlers.Home).Methods("GET")
	router.HandleFunc("/people", handlers.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", handlers.GetPersonById).Methods("GET")
	router.HandleFunc("/people/name/{name}", handlers.GetPersonByName).Methods("GET")
	router.HandleFunc("/health-check", handlers.HealthCheck).Methods("GET")

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", router)
}

package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
	"tsis1/api"
	_ "tsis1/api"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello to my phone book!\n")
	fmt.Fprintf(w, "You can use route /people to find all people in this phonebook.\n")
	fmt.Fprintf(w, "You can use route /people/{id} to find person by ID.\n")
	fmt.Fprintf(w, "You can use route /people/name/{name} to find person by name.\n")
	fmt.Fprintf(w, "You can use route /health-check to find some info about this phonebook.\n")
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("entering health check end point")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This app is like usefull phonebook. Creator: Kultyshev Artur :)")
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	for _, person := range api.People {
		json.NewEncoder(w).Encode(person)
	}
}

func GetPersonById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, person := range api.People {
		if person.ID == params["id"] {
			json.NewEncoder(w).Encode(person)
			return
		}
	}
	json.NewEncoder(w).Encode(&api.Person{})
}

func GetPersonByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var results []api.Person
	for _, person := range api.People {
		if strings.ToLower(person.Name) == params["name"] {
			results = append(results, person)
		}
	}
	json.NewEncoder(w).Encode(results)
}

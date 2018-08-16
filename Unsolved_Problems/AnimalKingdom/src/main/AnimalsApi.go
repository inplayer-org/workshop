package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/structures"

	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/db"

	"github.com/gorilla/mux"
)

func GetAnimals(DB *sql.DB) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		animals := db.SelectAllAnimals(DB)
		json.NewEncoder(w).Encode(animals)
	}
}

func GetAnimal(DB *sql.DB) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		parameter := mux.Vars(r)
		w.Header().Set("Content-Type", "application/json")
		animal := structures.Animal{}
		animal = db.SelectAnimal(DB, parameter["name"])
		log.Println(parameter["name"])
		json.NewEncoder(w).Encode(animal)
	}
}

func AddAnimal(DB *sql.DB) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		parameter := mux.Vars(r)
		w.Header().Set("Content-Type", "application/json")
		animal := structures.Animal{}
		animal.Name = parameter["name"]
		_ = json.NewDecoder(r.Body).Decode(&animal)
		db.InsertAnimal(DB, animal)
		json.NewEncoder(w).Encode(animal)
	}
}

func DeleteAnimal(DB *sql.DB) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		parameter := mux.Vars(r)
		w.Header().Set("Content-Type", "application/json")
		animalName := parameter["name"]
		db.DeleteAnimal(DB, animalName)

	}
}

func UpdateAnimal(DB *sql.DB) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		parameter := mux.Vars(r)
		w.Header().Set("Content-Type", "application/json")
		var animal structures.Animal
		animal.Name = parameter["name"]
		_ = json.NewDecoder(r.Body).Decode(&animal)
		db.UpdateAnimal(DB, animal)
		json.NewEncoder(w).Encode(animal)
	}
}

func main() {
	cfg, port := processFlags()
	DB := db.ConnectDB(cfg)
	router := mux.NewRouter()

	//Route Handlers/Endpoints
	router.HandleFunc("/animals", GetAnimals(DB)).Methods("GET")
	router.HandleFunc("/animals/{name}", GetAnimal(DB)).Methods("GET")
	router.HandleFunc("/animals/{name}", AddAnimal(DB)).Methods("POST")
	router.HandleFunc("/animals/{name}", DeleteAnimal(DB)).Methods("DELETE")
	router.HandleFunc("/animals/{name}", UpdateAnimal(DB)).Methods("PUT")
	log.Fatal(http.ListenAndServe(port, router))
}

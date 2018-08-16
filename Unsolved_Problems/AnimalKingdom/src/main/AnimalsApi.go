package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/structures"

	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/db"

	"github.com/gorilla/mux"
)

var IsString = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

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
		if !IsString(parameter["name"]) {
			respondWithError(w, http.StatusBadRequest, "Invalid animal name")
			return
		}
		animal := structures.Animal{}
		animal, err := db.SelectAnimal(DB, parameter["name"])
		fmt.Println(animal)
		fmt.Println("stignav")
		if err != nil {
			switch err {
			//ErrNoRows is returned by Scan when QueryRow doesn't return a row
			case sql.ErrNoRows:
				respondWithError(w, http.StatusNotFound, "Animal NOT FOUND")

			default:
				respondWithJSON(w, http.StatusInternalServerError, err.Error())
				return
			}

		}
		respondWithJSON(w, http.StatusOK, animal)

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

// func NotFound(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusLocked)
// }

// func Found(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusAlreadyReported)
// }

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
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
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
	//router.HandleFunc("/animals/{[0-9]+}", Found)
	//router.HandleFunc("/animals/{([0-9]+[a-zA-Z]+)|([a-zA-Z]+[0-9]+)}", NotFound)

	log.Fatal(http.ListenAndServe(port, router))
}

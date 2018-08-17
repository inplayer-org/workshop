package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/controllinput"
	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/structures"

	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/db"

	"github.com/gorilla/mux"
)

var IsString = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

func GetAnimals(DB *sql.DB) func(w http.ResponseWriter, req *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		animals := db.SelectAllAnimals(DB)
		respondWithJSON(w, http.StatusOK, animals)
	}
}

func GetAnimal(DB *sql.DB) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		parameter := mux.Vars(req)
		entry := parameter["entry"]
		if controllinput.IntOnly(entry) {
			GetAnimalByID(DB, entry, w)
		} else if controllinput.CheckString(&entry) {
			GetAnimalByName(DB, entry, w)
		} else {
			respondWithError(w, http.StatusBadRequest, "Request is a mix of ints and chars or is shorter than 2 characters or longer than 30 characters. Please provide correct request")
		}
	}

}

func GetAnimalByName(DB *sql.DB, name string, w http.ResponseWriter) {

	animal, err := db.SelectAnimalByName(DB, name)
	if err == sql.ErrNoRows {
		respondWithError(w, http.StatusNotFound, "Food name ("+name+") not present in database")
		return
	}
	respondWithJSON(w, http.StatusOK, animal)

}

func GetAnimalByID(DB *sql.DB, id string, w http.ResponseWriter) {
	i, _ := strconv.Atoi(id)
	animal, err := db.SelectAnimalByID(DB, i)
	if err == sql.ErrNoRows {
		respondWithError(w, http.StatusNotFound, "Animal index ("+id+") not present in database")
		return
	}
	respondWithJSON(w, http.StatusOK, animal)

}

func AddAnimal(DB *sql.DB) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		parameter := mux.Vars(r)
		w.Header().Set("Content-Type", "application/json")
		animal := structures.Animal{}
		json.NewDecoder(r.Body).Decode(&animal)
		if animal.Name != parameter["name"] {
			respondWithError(w, http.StatusBadRequest, "Your input for animal name does not match with requested animal name ")
			return
		}
		exist := db.Exists(DB, animal.Name)
		if exist == "1" {
			respondWithError(w, http.StatusUnsupportedMediaType, "This animal already exist")
			return
		}

		err := db.InsertAnimal(DB, animal)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, animal)
	}
}

func DeleteAnimal(DB *sql.DB) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		parameter := mux.Vars(r)
		w.Header().Set("Content-Type", "application/json")
		animalName := parameter["name"]
		exist := db.Exists(DB, animalName)
		fmt.Println(exist)
		if exist == "0" {
			respondWithError(w, http.StatusUnsupportedMediaType, "This animal does not exist")
			return
		}
		err := db.DeleteAnimalByName(DB, animalName)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithError(w, http.StatusOK, "Successfuly deleted")
	}
}

func UpdateAnimal(DB *sql.DB) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		//parameter := mux.Vars(r)
		w.Header().Set("Content-Type", "application/json")
		var animal structures.Animal

		_ = json.NewDecoder(r.Body).Decode(&animal)
		exist := db.Exists(DB, animal.Name)
		fmt.Println("dddddddd", exist)
		if exist == "0" {
			respondWithError(w, http.StatusUnsupportedMediaType, "This animal does not exist")
			return
		}

		animals, _ := db.UpdateAnimal(DB, animal)
		json.NewEncoder(w).Encode(animals)
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
	router.HandleFunc("/animals/{entry}", GetAnimal(DB)).Methods("GET")
	router.HandleFunc("/animals/{name}", AddAnimal(DB)).Methods("POST")
	router.HandleFunc("/animals/{name}", DeleteAnimal(DB)).Methods("DELETE")
	router.HandleFunc("/animals/{name}", UpdateAnimal(DB)).Methods("PUT")
	//router.HandleFunc("/animals/{[0-9]+}", Found)
	//router.HandleFunc("/animals/{([0-9]+[a-zA-Z]+)|([a-zA-Z]+[0-9]+)}", NotFound)

	log.Fatal(http.ListenAndServe(port, router))
}

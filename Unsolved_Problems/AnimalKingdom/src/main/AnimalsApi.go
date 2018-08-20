package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/controllinput"
	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/structures"

	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/db"

	"github.com/gorilla/mux"
)

func GetAnimals(DB *sql.DB) func(w http.ResponseWriter, req *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		animals := db.SelectAllAnimals(DB)
		respondWithJSON(w, http.StatusOK, animals)
	}
}

//Get animal by name or id
func GetAnimal(DB *sql.DB) func(w http.ResponseWriter, req *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		parameter := mux.Vars(req)
		entry := parameter["entry"]
		if !(controllinput.IntOnly(entry) || controllinput.CheckString(&entry)) {
			respondWithError(w, http.StatusBadRequest, "Request is a mix of ints and chars or is shorter than 2 characters or longer than 30 characters. Please provide correct request")
			return
		}
		animal, err := db.SelectAnimal(DB, entry)
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusNotFound, "Animal name ("+entry+") not present in database")
			return
		}
		respondWithJSON(w, http.StatusOK, animal)

	}

}

//Delete animal by name or id
func DeleteAnimal(DB *sql.DB) func(w http.ResponseWriter, req *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		parameter := mux.Vars(req)
		entry := parameter["entry"]
		if !(controllinput.IntOnly(entry) || controllinput.CheckString(&entry)) {
			respondWithError(w, http.StatusBadRequest, "Request is a mix of ints and chars or is shorter than 2 characters or longer than 30 characters. Please provide correct request")
			return
		}
		_, err := db.SelectAnimal(DB, entry)
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusBadRequest, "Name ("+entry+") not present in database")
			return
		}
		err1 := db.DeleteAnimal(DB, entry)
		if err1 != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithError(w, http.StatusOK, "Successfuly deleted")
	}

}

//Add animal by name
func AddAnimal(DB *sql.DB) func(w http.ResponseWriter, req *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		parameter := mux.Vars(req)
		entry := parameter["entry"]
		if controllinput.IntOnly(entry) {
			respondWithError(w, http.StatusBadRequest, "You can add animal just by name")
		} else if controllinput.CheckString(&entry) {
			AddAnimalByName(DB, entry, w, req)
		} else {
			respondWithError(w, http.StatusBadRequest, "Request is a mix of ints and chars or is shorter than 2 characters or longer than 30 characters. Please provide correct request")
		}
	}

}

func AddAnimalByName(DB *sql.DB, animalName string, w http.ResponseWriter, req *http.Request) {
	animal := structures.Animal{}
	json.NewDecoder(req.Body).Decode(&animal)
	if animal.Name != animalName {
		respondWithError(w, http.StatusMultipleChoices, "Entry name("+animalName+") and JSON name("+animal.Name+") has to be equal")
		return
	}
	exist := db.Exists(DB, animalName)
	if exist == "1" {
		respondWithError(w, http.StatusAlreadyReported, "Animal ("+animalName+") already exists in database, If you want to update entry use the PUT method")
		return
	}
	err := db.InsertAnimal(DB, animal)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, animal)
}

//Update animal by name
func UpdateAnimal(DB *sql.DB) func(w http.ResponseWriter, req *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		parameter := mux.Vars(req)
		entry := parameter["entry"]
		if controllinput.IntOnly(entry) {
			respondWithError(w, http.StatusBadRequest, "You can update animal just by name")
		} else if controllinput.CheckString(&entry) {
			UpdateAnimalByName(DB, entry, w, req)
		} else {
			respondWithError(w, http.StatusBadRequest, "Request is a mix of ints and chars or is shorter than 2 characters or longer than 30 characters. Please provide correct request")
		}
	}

}

func UpdateAnimalByName(DB *sql.DB, animalName string, w http.ResponseWriter, req *http.Request) {

	var animal structures.Animal
	_ = json.NewDecoder(req.Body).Decode(&animal)
	if animal.Name != animalName {
		respondWithError(w, http.StatusMultipleChoices, "Entry name("+animalName+") and JSON name("+animal.Name+") has to be equal")
		return
	}
	exist := db.Exists(DB, animal.Name)
	if exist == "0" {
		respondWithError(w, http.StatusBadRequest, "Animal ("+animalName+") not present in database")
		return
	}
	animals, _ := db.UpdateAnimal(DB, animal)
	json.NewEncoder(w).Encode(animals)

}

//Select all food that an animal eat
func GetAnimalFoods(DB *sql.DB) func(w http.ResponseWriter, req *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		parameter := mux.Vars(req)
		entryAnimal := parameter["animal"]
		if !(controllinput.IntOnly(entryAnimal) || controllinput.CheckString(&entryAnimal)) {
			respondWithError(w, http.StatusBadRequest, "Request is a mix of ints and chars or is shorter than 2 characters or longer than 30 characters. Please provide correct request")
			return
		}
		_, err := db.SelectAnimal(DB, entryAnimal)
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusNotFound, "Animal name ("+entryAnimal+") not present in database")
			return
		}
		listOfFood := db.SelectFoodAnimalEat(DB, entryAnimal)
		respondWithJSON(w, http.StatusOK, listOfFood)

	}

}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
func respondWithJSON(w http.ResponseWriter, code int, dataStruct interface{}) {
	response, err := json.Marshal(dataStruct)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)

	} else {
		w.WriteHeader(code)
	}

	w.Write(response)
}
func main() {
	cfg, port := processFlags()
	DB := db.ConnectDB(cfg)
	router := mux.NewRouter()

	//Route Handlers/Endpoints
	router.HandleFunc("/animals", GetAnimals(DB)).Methods("GET")
	router.HandleFunc("/animals/{entry}", GetAnimal(DB)).Methods("GET")
	router.HandleFunc("/animals/{entry}", AddAnimal(DB)).Methods("POST")
	router.HandleFunc("/animals/{entry}", DeleteAnimal(DB)).Methods("DELETE")
	router.HandleFunc("/animals/{entry}", UpdateAnimal(DB)).Methods("PUT")
	//router.HandleFunc("/food/{animal}", GetAnimalFoods(DB)).Methods("GET")
	router.HandleFunc("/eat/{animal}", GetAnimalFoods(DB)).Methods("GET")
	log.Fatal(http.ListenAndServe(port, router))
}

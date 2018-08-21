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

var IsMix = regexp.MustCompile(`([a-zA-Z+]+[0-9+]){1,30}$`).MatchString
var IsString = regexp.MustCompile(`^[a-zA-Z+]+$`).MatchString

//Get all animals
func GetAnimals(DB *sql.DB) func(w http.ResponseWriter, req *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		animals := db.SelectAllAnimals(DB)
		respondWithJSON(w, http.StatusOK, animals)
	}
}

//Get animal by name or id
func GetAnimal(DB *sql.DB) func(w http.ResponseWriter, req *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		parameter := mux.Vars(req)
		entry := parameter["entry"]
		if IsMix(entry) {
			respondWithError(w, http.StatusBadRequest, "Request is a mix of ints and chars ")
			return
		}
		animal, err := db.SelectAnimal(DB, entry)
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusNotFound, "Animal ("+entry+") not present in database")
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
		if IsMix(entry) {
			respondWithError(w, http.StatusBadRequest, "Request is a mix of ints and chars")
			return
		}
		_, err := db.SelectAnimal(DB, entry)
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusBadRequest, "Name ("+entry+") not present in database")
			return
		}
		err1 := db.DeleteAnimal(DB, entry)
		if err1 != nil {
			respondWithError(w, http.StatusInternalServerError, err1.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, "Successfuly deleted")
	}

}

//Add animal by name
func AddAnimal(DB *sql.DB) func(w http.ResponseWriter, req *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		parameter := mux.Vars(req)
		entry := parameter["entry"]
		if IsMix(entry) {
			respondWithError(w, http.StatusBadRequest, "Request is a mix of ints and chars")
			return
		} else if !IsString(entry) {
			respondWithError(w, http.StatusBadRequest, "You can add animal just by name")
			return
		}
		animal := structures.Animal{}
		json.NewDecoder(req.Body).Decode(&animal)
		if animal.Name != entry {
			respondWithError(w, http.StatusMultipleChoices, "Entry ("+entry+") and JSON ("+animal.Name+") has to be equal")
			return
		}
		err := db.InsertAnimal(DB, animal)
		fmt.Println(err)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}
		respondWithJSON(w, http.StatusOK, animal)
	}

}

//Update animal by name
func UpdateAnimal(DB *sql.DB) func(w http.ResponseWriter, req *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		parameter := mux.Vars(req)
		entry := parameter["entry"]
		if IsMix(entry) {
			respondWithError(w, http.StatusBadRequest, "Request is a mix of ints and chars")
			return
		} else if !IsString(entry) {
			respondWithError(w, http.StatusBadRequest, "You can add animal just by name")
			return
		}
		var animal structures.Animal
		_ = json.NewDecoder(req.Body).Decode(&animal)
		if animal.Name != entry {
			respondWithError(w, http.StatusMultipleChoices, "Entry name("+entry+") and JSON name("+animal.Name+") has to be equal")
			return
		}
		_, err := db.SelectAnimal(DB, entry)
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusBadRequest, "Animal ("+entry+") not present in database")
			return
		}
		animals, _ := db.UpdateAnimal(DB, animal)
		json.NewEncoder(w).Encode(animals)
	}

}

//Select all food that an animal eat
func GetAnimalFoods(DB *sql.DB) func(w http.ResponseWriter, req *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		parameter := mux.Vars(req)
		entryAnimal := parameter["animal"]
		if IsMix(entryAnimal) {
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
	router.HandleFunc("/eat/{animal}", GetAnimalFoods(DB)).Methods("GET")
	log.Fatal(http.ListenAndServe(port, router))
}

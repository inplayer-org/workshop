package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

// func GetAnimal(DB *sql.DB) func(w http.ResponseWriter, req *http.Request) {

// 	return func(w http.ResponseWriter, req *http.Request) {
// 		parameter := mux.Vars(req)
// 		entry := parameter["entry"]
// 		if controllinput.IntOnly(entry) {
// 			GetAnimalByID(DB, entry, w)
// 		} else if controllinput.CheckString(&entry) {
// 			GetAnimalByName(DB, entry, w)
// 		} else {
// 			respondWithError(w, http.StatusBadRequest, "Request is a mix of ints and chars or is shorter than 2 characters or longer than 30 characters. Please provide correct request")
// 		}
// 	}

// }

func SAnimal(DB *sql.DB) func(w http.ResponseWriter, req *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		parameter := mux.Vars(req)
		entry := parameter["entry"]
		var query string
		if controllinput.IntOnly(entry) {
			query = "animalID"
		} else if controllinput.CheckString(&entry) {
			query = "name"
		} else {
			respondWithError(w, http.StatusBadRequest, "Request is a mix of ints and chars or is shorter than 2 characters or longer than 30 characters. Please provide correct request")
			return
		}
		animal, err := db.Select(DB, entry, query)
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusNotFound, "Animal name ("+entry+") not present in database")
			return
		}
		respondWithJSON(w, http.StatusOK, animal)

	}

}

// func GetAnimalByName(DB *sql.DB, name string, w http.ResponseWriter) {

// 	animal, err := db.SelectAnimalByName(DB, name)
// 	if err == sql.ErrNoRows {
// 		respondWithError(w, http.StatusNotFound, "Animal name ("+name+") not present in database")
// 		return
// 	}
// 	respondWithJSON(w, http.StatusOK, animal)

// }

// func GetAnimalByID(DB *sql.DB, id string, w http.ResponseWriter) {
// 	i, _ := strconv.Atoi(id)
// 	animal, err := db.SelectAnimalByID(DB, i)
// 	if err == sql.ErrNoRows {
// 		respondWithError(w, http.StatusNotFound, "Animal index ("+id+") not present in database")
// 		return
// 	}
// 	respondWithJSON(w, http.StatusOK, animal)

// }

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

func DeleteAnimal(DB *sql.DB) func(w http.ResponseWriter, req *http.Request) {

	return func(w http.ResponseWriter, req *http.Request) {
		parameter := mux.Vars(req)
		entry := parameter["entry"]
		if controllinput.IntOnly(entry) {
			DeleteAnimalByID(DB, entry, w)
		} else if controllinput.CheckString(&entry) {
			DeleteAnimalByName(DB, entry, w)
		} else {
			respondWithError(w, http.StatusBadRequest, "Request is a mix of ints and chars or is shorter than 2 characters or longer than 30 characters. Please provide correct request")
		}
	}

}

func DeleteAnimalByName(DB *sql.DB, animalName string, w http.ResponseWriter) {

	exist := db.Exists(DB, animalName)
	if exist == "0" {
		respondWithError(w, http.StatusBadRequest, "Name ("+animalName+") not present in database")
		return
	}
	err := db.DeleteAnimalByName(DB, animalName)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithError(w, http.StatusOK, "Successfuly deleted")
}

func DeleteAnimalByID(DB *sql.DB, animalID string, w http.ResponseWriter) {

	id, _ := strconv.Atoi(animalID)
	exist := db.ExistsID(DB, id)
	log.Println(exist)
	if exist == "0" {
		respondWithError(w, http.StatusBadRequest, "Index ("+animalID+") not present in database")
		return
	}
	err := db.DeleteAnimalByID(DB, id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithError(w, http.StatusOK, "Successfuly deleted")
}

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
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}
func respondWithJSON(w http.ResponseWriter, code int, dataStruct interface{}) {
	log.Println(dataStruct)
	response, err := json.Marshal(dataStruct)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
	} else {
		w.WriteHeader(code)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
func main() {
	cfg, port := processFlags()
	DB := db.ConnectDB(cfg)
	router := mux.NewRouter()
	//a := structures.Animal{}
	a, _ := db.Select(DB, "olo", "name")
	fmt.Println(a)
	//Route Handlers/Endpoints

	router.HandleFunc("/animals", GetAnimals(DB)).Methods("GET")
	router.HandleFunc("/animals/{entry}", SAnimal(DB)).Methods("GET")

	//router.HandleFunc("/animals/{entry}", GetAnimal(DB)).Methods("GET")
	router.HandleFunc("/animals/{entry}", AddAnimal(DB)).Methods("POST")
	router.HandleFunc("/animals/{entry}", DeleteAnimal(DB)).Methods("DELETE")
	router.HandleFunc("/animals/{entry}", UpdateAnimal(DB)).Methods("PUT")

	log.Fatal(http.ListenAndServe(port, router))
}

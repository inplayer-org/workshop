package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/structures"
	"strconv"
)


var animals []structures.Animal


func GetAnimals(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(animals)
}

func GetAnimal(w http.ResponseWriter, req *http.Request){

	parameters := mux.Vars(req)

	for _,animal := range animals{
		w.Header().Set("Content-Type","application/json")
		if strconv.Itoa(animal.AnimalID) == parameters["id"]{
			json.NewEncoder(w).Encode(animal)
			return
		}
	}

	//Returns not found if id isn't present in the database
	w.WriteHeader(http.StatusNotFound)
}

/*func AddAnimal(w http.ResponseWriter, req *http.Request){
	var animal structures.Animal
	param := mux.Vars(req)

	for _,anim := range animals{
		if strconv.Itoa(anim.AnimalID)==param["id"]{
			w.WriteHeader(http.StatusAlreadyReported)
			return
		}
	}
	json.NewDecoder(req.Body).Decode(&animal)
	animals = append(animals, animal)
	json.NewEncoder(w).Encode(animal)
}*/


func main(){

	//Mock objects
	anim1 := structures.Animal{AnimalID:1,Name:"Tigar",Species:"carnivore"}
	anim2 := structures.Animal{AnimalID:2,Name:"Rabbit",Species:"herbivore"}
	animals = append(animals, anim1)
	animals = append(animals, anim2)

	//Initializing a Router
	router := mux.NewRouter()

	//Route Handlers/Endpoints
	router.HandleFunc("/animals",GetAnimals).Methods("GET")
	router.HandleFunc("/animals/{id}",GetAnimal).Methods("GET")
	router.HandleFunc("/animals/{id}",AddAnimal).Methods("POST")
	router.HandleFunc("/animals",GetAnimals).Methods("GET")
	router.HandleFunc("/animals",GetAnimals).Methods("GET")


	log.Fatal(http.ListenAndServe(":8000",router))
}
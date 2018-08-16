package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/structures"
	"encoding/json"
	"fmt"
	"strconv"
)

//PROTOTYPE VERSION
var foods []structures.Food


func GetFoods(w http.ResponseWriter, req *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foods)
}

func GetFood(w http.ResponseWriter, req *http.Request){

	params := mux.Vars(req)

	for _,food := range foods{
		if strconv.Itoa(food.FoodID) == params["id"]{
			w.Header().Set("Content-Type","application/json")
			json.NewEncoder(w).Encode(food)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func AddFood(w http.ResponseWriter,req *http.Request){

	var food structures.Food
	json.NewDecoder(req.Body).Decode(&food)

	//Check if the food id is already present in the database
	for _,checkFood := range foods{
		if food.FoodID == checkFood.FoodID{
			w.WriteHeader(http.StatusAlreadyReported)
			json.NewEncoder(w).Encode(checkFood)
			return
		}
	}

	w.WriteHeader(http.StatusAccepted)
	foods = append(foods,food)
	json.NewEncoder(w).Encode(food)

}

func DeleteFood(w http.ResponseWriter,req *http.Request){

/*	var food structures.Food

	param := mux.Vars(req)

	for()*/
}


func main() {

	router := mux.NewRouter()


	//Mock object
	foods = append(foods,structures.Food{FoodID:1,Type:"Herb",Name:"Grass"})
	foods = append(foods,structures.Food{FoodID:2,Type:"Meat",Name:"Rabbit"})
	foods = append(foods,structures.Food{FoodID:3,Type:"Herb",Name:"Leaves"})
	foods = append(foods,structures.Food{FoodID:4,Type:"Meat",Name:"Deer"})

	fmt.Println(foods)

	//Handlers
	router.HandleFunc("/foods",GetFoods).Methods("GET")
	router.HandleFunc("/foods/{id}",GetFood).Methods("GET")
	router.HandleFunc("/foods/{id}",AddFood).Methods("POST")


	http.ListenAndServe(":8889", router)

}

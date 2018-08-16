package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/structures"
	"encoding/json"
	"fmt"
	"strconv"
	"log"
)

//PROTOTYPE VERSION
var foods []structures.Food

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, dataStruct interface{}) {
	log.Println(dataStruct)
	response, err := json.Marshal(dataStruct)
	if err!=nil{
		w.WriteHeader(http.StatusUnprocessableEntity)
	}else{
		w.WriteHeader(code)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func GetFoods(w http.ResponseWriter, req *http.Request){
	respondWithJSON(w,http.StatusOK,foods)
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
	respondWithError(w,http.StatusBadRequest,"Index not present in database")
}

func AddFood(w http.ResponseWriter,req *http.Request){

	var food structures.Food
	json.NewDecoder(req.Body).Decode(&food)

	//Check if the food id is already present in the database
	for _,checkFood := range foods{
		if food.FoodID == checkFood.FoodID{
			respondWithJSON(w,http.StatusAlreadyReported,checkFood)
			return
		}
	}
	respondWithJSON(w,http.StatusAccepted,food)
	foods = append(foods,food)

}

func DeleteFood(w http.ResponseWriter,req *http.Request){

	param := mux.Vars(req)
	for i,food := range foods {
		if strconv.Itoa(food.FoodID) == param["id"]{
			foods = append(foods[:i],foods[i+1:]...)
			respondWithJSON(w,http.StatusAccepted,food)
			return
		}
	}
	respondWithError(w,http.StatusBadRequest,"Index not present in database")


}

func UpdateFood(w http.ResponseWriter,req *http.Request){

	param := mux.Vars(req)
	var updFood structures.Food
	json.NewDecoder(req.Body).Decode(&updFood)
	if strconv.Itoa(updFood.FoodID) != param["id"]{
		parID,_ := strconv.Atoi(param["id"])
		respondWithJSON(w,http.StatusMultipleChoices,map[string]structures.Food{"struct1":{FoodID:parID},"struct2":updFood})
		return
	}
	for i,food := range foods{
		if strconv.Itoa(food.FoodID) == param["id"]{
			foods[i] = updFood
			respondWithJSON(w,http.StatusAccepted,updFood)
			return
		}
	}
	respondWithError(w,http.StatusBadRequest,"Index not present in database")
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
	router.HandleFunc("/foods/{id}",DeleteFood).Methods("DELETE")
	router.HandleFunc("/foods/{id}",UpdateFood).Methods("PUT")

	http.ListenAndServe(":8889", router)

}



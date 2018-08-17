package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/structures"
	"encoding/json"
	"fmt"
	"strconv"
	"log"
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/controllinput"
	_ "github.com/go-sql-driver/mysql"
	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/db"
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

func GetFoods(DB *sql.DB)func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		respondWithJSON(w, http.StatusOK, db.SelectAllFood(DB))
	}
}

func GetFood(DB *sql.DB)func(w http.ResponseWriter, req *http.Request){
	return func(w http.ResponseWriter, req *http.Request) {
		parameter := mux.Vars(req)
		entry := parameter["entry"]
		if controllinput.IntOnly(entry){
			GetFoodByID(DB,entry,w)
		}else if controllinput.CheckString(&entry){
			GetFoodByName(DB,entry,w)
		}else{
			respondWithError(w,http.StatusBadRequest,"Request is a mix of ints and chars or is shorter than 2 characters or longer than 30 characters. Please provide correct request")
		}
	}

}

func GetFoodByName(DB *sql.DB,name string,w http.ResponseWriter){
	food := db.SelectFoodbyName(DB,name)
	if  food == nil{
		respondWithError(w,http.StatusNotFound,"Food index ("+name+") not present in database")
	}else{
		respondWithJSON(w,http.StatusOK,food)
	}
}

func GetFoodByID(DB *sql.DB,ID string,w http.ResponseWriter){
		id,_ := strconv.Atoi(ID)
		food := db.SelectFoodbyID(DB,id)
		if  food == nil{
		respondWithError(w,http.StatusNotFound,"Food index ("+ID+") not present in database")
		}else{
			respondWithJSON(w,http.StatusOK,food)
		}
}


func AddFood(DB *sql.DB)func(w http.ResponseWriter, req *http.Request){
	return func(w http.ResponseWriter, req *http.Request) {
		var food structures.Food
		parameters := mux.Vars(req)
		entry := parameters["entry"]
		json.NewDecoder(req.Body).Decode(&food)

		if food.FoodID==0 || !controllinput.CheckString(&food.Name) || !controllinput.CheckString(&food.Type){
			respondWithError(w,http.StatusBadRequest,"Invalid structure data, Food name and type has to be between 2 and 30 characters and can't contain numbers. Food id shouldn't be lower than 1")

		}else if controllinput.IntOnly(entry){
			if entry!= strconv.Itoa(food.FoodID){
				respondWithError(w,http.StatusMultipleChoices,"Entry index("+entry+") and JSON index("+strconv.Itoa(food.FoodID)+") has to be equal")
			}else{
				AddFilteredFood(DB,food,w)
			}

		}else if controllinput.CheckString(&entry){
			if entry!=food.Name {
				respondWithError(w,http.StatusMultipleChoices,"Entry name("+entry+") and JSON name("+food.Name+") has to be equal")
			}else{
				AddFilteredFood(DB, food, w)
				}

		}else{
			respondWithError(w,http.StatusBadRequest,"Request is a mix of ints and chars or is shorter than 2 characters or longer than 30 characters. Please provide correct request")
		}
	}
}

func AddFilteredFood(DB *sql.DB,newFood structures.Food ,w http.ResponseWriter){
	//Check if the food id is already present in the database


		if db.SelectFoodbyID(DB,newFood.FoodID)!=nil {
			respondWithJSON(w, http.StatusAlreadyReported, "Food with ID ("+strconv.Itoa(newFood.FoodID)+") already exists in database, If you want to update entry use the PUT method")
			return
		}else if db.SelectFoodbyName(DB,newFood.Name)!=nil{
			respondWithJSON(w, http.StatusAlreadyReported, "Food with name ("+newFood.Name+") already exists in database, If you want to update entry use the PUT method")
			return
		}
	db.InsertFood(DB,newFood)
	respondWithJSON(w, http.StatusCreated, newFood)
}

func DeleteFood(DB *sql.DB)func(w http.ResponseWriter, req *http.Request){
	return func(w http.ResponseWriter, req *http.Request) {
		parameter := mux.Vars(req)
		entry := parameter["entry"]
		if controllinput.IntOnly(entry){
			DeleteFoodByID(DB,entry,w)
		}else if controllinput.CheckString(&entry){
			DeleteFoodByName(DB,entry,w)
		}else{
			respondWithError(w,http.StatusBadRequest,"Request is a mix of ints and chars or is shorter than 2 characters or longer than 30 characters. Please provide correct request")
		}
	}
}

func DeleteFoodByID(DB *sql.DB,entry string,w http.ResponseWriter) {

		for i, checkFood := range foods {
			if strconv.Itoa(checkFood.FoodID) == entry {
				foods = append(foods[:i], foods[i+1:]...)
				respondWithJSON(w, http.StatusAccepted, checkFood)
				return
			}
		}
		respondWithError(w, http.StatusBadRequest, "Index ("+entry+") not present in database")

}

func DeleteFoodByName(DB *sql.DB,entry string,w http.ResponseWriter) {
	for i, checkFood := range foods {
		if checkFood.Name == entry {
			foods = append(foods[:i], foods[i+1:]...)
			respondWithJSON(w, http.StatusAccepted, checkFood)
			return
		}
	}
	respondWithError(w, http.StatusBadRequest, "Name ("+entry+") not present in database")
}


func UpdateFood(DB *sql.DB)func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		var food structures.Food
		parameters := mux.Vars(req)
		entry := parameters["entry"]
		json.NewDecoder(req.Body).Decode(&food)

		if food.FoodID<=0 || !controllinput.CheckString(&food.Name) || !controllinput.CheckString(&food.Type){
			respondWithError(w,http.StatusBadRequest,"Invalid structure data, Food name and type has to be between 2 and 30 characters and can't contain numbers. Food id shouldn't be lower than 1")

		}else if controllinput.IntOnly(entry){
			if entry!= strconv.Itoa(food.FoodID){
				respondWithError(w,http.StatusMultipleChoices,"Entry index("+entry+") and JSON index("+strconv.Itoa(food.FoodID)+") has to be equal")
			}else{
				UpdateFoodByID(DB,food,w)
			}

		}else if controllinput.CheckString(&entry){
			if entry!=food.Name {
				respondWithError(w,http.StatusMultipleChoices,"Entry name("+entry+") and JSON name("+food.Name+") has to be equal")
			}else{
				UpdateFoodByName(DB, food, w)
			}

		}else{
			respondWithError(w,http.StatusBadRequest,"Request is a mix of ints and chars or is shorter than 2 characters or longer than 30 characters. Please provide correct request")
		}
	}

}


func UpdateFoodByID(DB *sql.DB,updateFood structures.Food,w http.ResponseWriter){


		for i, food := range foods {
			if food.FoodID == updateFood.FoodID {
				for _,checkNameExistance := range foods{
					if checkNameExistance.Name == updateFood.Name{
						respondWithError(w,http.StatusAlreadyReported,"Food with name ("+updateFood.Name+") already exist in the database, and there can't be multiple foods with equal name")
						return
				}
				}
				foods[i] = updateFood
				respondWithJSON(w, http.StatusAccepted, updateFood)
				return
			}
		}
	respondWithError(w, http.StatusBadRequest, "Food with ID ("+strconv.Itoa(updateFood.FoodID)+") doesn't exist in database. If you want to add new entry use the POST method")

}

func UpdateFoodByName(DB *sql.DB,updateFood structures.Food,w http.ResponseWriter){


	for i, food := range foods {
		if food.Name == updateFood.Name {
			for _,checkNameExistance := range foods{
				if checkNameExistance.FoodID == updateFood.FoodID{
					respondWithError(w,http.StatusAlreadyReported,"Food with ID ("+strconv.Itoa(updateFood.FoodID)+") already exist in the database, and there can't be multiple foods with equal ID")
					return
				}
			}
			foods[i] = updateFood
			respondWithJSON(w, http.StatusAccepted, updateFood)
			return
		}
	}
	respondWithError(w, http.StatusBadRequest, "Food with name ("+updateFood.Name+") doesn't exist in database. If you want to add new entry use the POST method")
}

func main() {

	router := mux.NewRouter()

	DB := db.ConnectDB("root:1111@tcp(127.0.0.1:3306)/animalKingdom")

	//Mock object
	foods = append(foods,structures.Food{FoodID:1,Type:"herb",Name:"grass"})
	foods = append(foods,structures.Food{FoodID:2,Type:"meat",Name:"rabbit"})
	foods = append(foods,structures.Food{FoodID:3,Type:"herb",Name:"leaves"})
	foods = append(foods,structures.Food{FoodID:4,Type:"meat",Name:"deer"})

	fmt.Println(foods)

	//Handlers
	router.HandleFunc("/foods",GetFoods(DB)).Methods("GET")
	router.HandleFunc("/foods/{entry}",GetFood(DB)).Methods("GET")
	router.HandleFunc("/foods/{entry}",AddFood(DB)).Methods("POST")
	router.HandleFunc("/foods/{entry}",DeleteFood(DB)).Methods("DELETE")
	router.HandleFunc("/foods/{entry}",UpdateFood(DB)).Methods("PUT")


	http.ListenAndServe(":8889", router)

}



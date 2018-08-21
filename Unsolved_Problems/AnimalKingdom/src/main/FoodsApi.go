package main

// import (
// 	"github.com/gorilla/mux"
// 	"net/http"
// 	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/structures"
// 	"encoding/json"
// 	"database/sql"
// 	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/controllinput"
// 	_ "github.com/go-sql-driver/mysql"
// 	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/db"
// 	resp "repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/responses"
// 	. "repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/consts"
// 	handler "repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/foodhandlers"
// )

// //PROTOTYPE VERSION
// //var foods []structures.Food

// func GetFoods(DB *sql.DB)func(w http.ResponseWriter, req *http.Request) {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		resp.RespondWithJSON(w, http.StatusOK, db.SelectAllFood(DB))
// 	}
// }

// func GetFood(DB *sql.DB)func(w http.ResponseWriter, req *http.Request){
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		parameter := mux.Vars(req)
// 		request := parameter["entry"]

// 		filteredRequest,queryCallerType := controllinput.ValidateRequest(request,w)
// 		if queryCallerType==Int{
// 			handler.GetFoodByID(DB,filteredRequest.(int),w)
// 		}else if queryCallerType==String{
// 			handler.GetFoodByName(DB,filteredRequest.(string),w)
// 		}
// 	}

// }

// func AddFood(DB *sql.DB)func(w http.ResponseWriter, req *http.Request) {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		var food structures.Food
// 		parameters := mux.Vars(req)
// 		entry := parameters["entry"]
// 		json.NewDecoder(req.Body).Decode(&food)

// 		if controllinput.ValidateStructureData(food, w) {
// 			queryCallerType := controllinput.ValidateEntry(food,entry,w)

// 			if queryCallerType==Int || queryCallerType==String {
// 				handler.AddFilteredFood(DB, food, w)
// 			}
// 		}
// 	}
// }

// func DeleteFood(DB *sql.DB)func(w http.ResponseWriter, req *http.Request){
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		parameter := mux.Vars(req)
// 		request := parameter["entry"]

// 		filteredRequest,queryCallerType := controllinput.ValidateRequest(request,w)
// 		if queryCallerType==Int{
// 			handler.DeleteFoodByID(DB,filteredRequest.(int),w)
// 		}else if queryCallerType==String{
// 			handler.DeleteFoodByName(DB,filteredRequest.(string),w)
// 		}
// 	}
// }

// func UpdateFood(DB *sql.DB)func(w http.ResponseWriter, req *http.Request) {
// 	return func(w http.ResponseWriter, req *http.Request) {
// 		var food structures.Food
// 		parameters := mux.Vars(req)
// 		entry := parameters["entry"]
// 		json.NewDecoder(req.Body).Decode(&food)

// 		if controllinput.ValidateStructureData(food, w) {
// 			queryCallerType := controllinput.ValidateEntry(food,entry,w)

// 			if queryCallerType==Int{
// 				handler.UpdateFilteredFood(DB,food,w)
// 			}else if queryCallerType==String {
// 				resp.RespondWithError(w,http.StatusBadRequest,resp.NotAllowedToUpdateBy("name"))
// 			}
// 		}
// 	}

// }

// func main() {

// 	router := mux.NewRouter()

// 	DB := db.ConnectDB("root:1111@tcp(127.0.0.1:3306)/animalKingdom")

// 	//Mock object
// 	//foods = append(foods,structures.Food{FoodID:1,Type:"herb",Name:"grass"})
// 	//foods = append(foods,structures.Food{FoodID:2,Type:"meat",Name:"rabbit"})
// 	//foods = append(foods,structures.Food{FoodID:3,Type:"herb",Name:"leaves"})
// 	//foods = append(foods,structures.Food{FoodID:4,Type:"meat",Name:"deer"})

// 	//fmt.Println(foods)

// 	//Handlers
// 	router.HandleFunc("/foods",GetFoods(DB)).Methods("GET")
// 	router.HandleFunc("/foods/{entry}",GetFood(DB)).Methods("GET")
// 	router.HandleFunc("/foods/{entry}",AddFood(DB)).Methods("POST")
// 	router.HandleFunc("/foods/{entry}",DeleteFood(DB)).Methods("DELETE")
// 	router.HandleFunc("/foods/{entry}",UpdateFood(DB)).Methods("PUT")

// 	http.ListenAndServe(":8889", router)

// }

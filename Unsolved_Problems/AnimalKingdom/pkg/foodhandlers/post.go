package foodhandlers

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/structures"
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/db"
	"strconv"
	resp "repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/responses"
)

func AddFilteredFood(DB *sql.DB,newFood structures.Food ,w http.ResponseWriter){
	//Check if the food id is already present in the database


	if db.SelectFoodbyID(DB,newFood.FoodID)!=nil {
		resp.RespondWithJSON(w, http.StatusAlreadyReported, resp.AlreadyExist("ID",strconv.Itoa(newFood.FoodID)))
		return
	}else if db.SelectFoodbyName(DB,newFood.Name)!=nil{
		resp.RespondWithJSON(w, http.StatusAlreadyReported, resp.AlreadyExist("Name",newFood.Name))
		return
	}
	db.InsertFood(DB,newFood)
	resp.RespondWithJSON(w, http.StatusCreated, newFood)
}

package foodhandlers

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/structures"
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/db"
	resp "repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/responses"
)

func UpdateFilteredFood(DB *sql.DB,updateFood structures.Food,w http.ResponseWriter){

	food := db.SelectFoodbyName(DB,updateFood.Name)
	if food!=nil{
		resp.RespondWithError(w, http.StatusAlreadyReported, "Food name("+updateFood.Name+") is already present in the database")
		return
	}
	food,err := db.UpdateFood(DB,updateFood)
	if err!=nil{
		resp.RespondWithError(w,http.StatusInternalServerError,resp.ErrorDuringExec("updating"))
		return
	}
	resp.RespondWithJSON(w,http.StatusOK,updateFood)


}

//func UpdateFoodByName(DB *sql.DB,updateFood structures.Food,w http.ResponseWriter){
//
//
//	for i, food := range foods {
//		if food.Name == updateFood.Name {
//			for _,checkNameExistance := range foods{
//				if checkNameExistance.FoodID == updateFood.FoodID{
//					resp.RespondWithError(w,http.StatusAlreadyReported,"Food with ID ("+strconv.Itoa(updateFood.FoodID)+") already exist in the database, and there can't be multiple foods with equal ID")
//					return
//				}
//			}
//			foods[i] = updateFood
//			resp.RespondWithJSON(w, http.StatusAccepted, updateFood)
//			return
//		}
//	}
//	resp.RespondWithError(w, http.StatusBadRequest, "Food with name ("+updateFood.Name+") doesn't exist in database. If you want to add new entry use the POST method")
//}
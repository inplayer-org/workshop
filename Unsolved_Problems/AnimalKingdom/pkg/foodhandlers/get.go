package foodhandlers

import (
	"database/sql"
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/db"
	"strconv"
	resp "repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/responses"
)

func GetFoodByName(DB *sql.DB,name string,w http.ResponseWriter){
	food := db.SelectFoodbyName(DB,name)
	if  food == nil{
		resp.RespondWithError(w,http.StatusNotFound,"Food index ("+name+") not present in database")
	}else{
		resp.RespondWithJSON(w,http.StatusOK,food)
	}
}

func GetFoodByID(DB *sql.DB,ID int,w http.ResponseWriter){
	food := db.SelectFoodbyID(DB,ID)
	if  food == nil{
		resp.RespondWithError(w,http.StatusNotFound,"Food index ("+strconv.Itoa(ID)+") not present in database")
	}else{
		resp.RespondWithJSON(w,http.StatusOK,food)
	}
}
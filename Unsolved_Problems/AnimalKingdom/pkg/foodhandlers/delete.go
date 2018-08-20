package foodhandlers

import (
	"database/sql"
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/db"
	"strconv"
	resp "repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/responses"
)

func DeleteFoodByID(DB *sql.DB,ID int,w http.ResponseWriter) {

	food := db.SelectFoodbyID(DB,ID)
	if food==nil{
		resp.RespondWithError(w,http.StatusNotFound,resp.NotFound("index",strconv.Itoa(ID)))
		return
	}
	err := db.DeleteFoodbyID(DB,ID)
	if err!=nil{
		resp.RespondWithError(w,http.StatusInternalServerError,resp.ErrorDuringExec("deletion"))
		return
	}
	resp.RespondWithJSON(w,http.StatusOK,food)

}

func DeleteFoodByName(DB *sql.DB,name string,w http.ResponseWriter) {
	food := db.SelectFoodbyName(DB,name)
	if food==nil{
		resp.RespondWithError(w,http.StatusNotFound,resp.NotFound("name",name))
		return
	}
	err := db.DeleteFoodbyName(DB,name)
	if err!=nil{
		resp.RespondWithError(w,http.StatusInternalServerError,resp.ErrorDuringExec("deletion"))
		return
	}
	resp.RespondWithJSON(w,http.StatusOK,food)
}
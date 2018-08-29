package HandlersFunc

import (
	"net/http"
	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errorhandlers"
)

func (a *App) GetLocationByName (w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	name := vars["name"]
	/*if name != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid player name")
		return
	} */

	e := structures.PlayerStats{}
	player,err := e.GetPlayersByLocation(a.DB,name)
/*
		switch err {
		case sql.ErrNoRows:
			errorhandlers.RespondWithError(w, http.StatusNotFound, "Location not found")
		default:
			errorhandlers.RespondWithError(w, http.StatusInternalServerError, "Server Error")
		}
		return*/
		if err!=nil {
			panic(err)
	}
	//fmt.Println(player)
	structures.Tmpl.ExecuteTemplate(w,"TableRanking",player)
}



func (a *App) GetLocations(w http.ResponseWriter, r *http.Request) {


	locations, err := structures.GetAllLocations(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			errorhandlers.RespondWithError(w, http.StatusNotFound, "no locations found")

		default:
			errorhandlers.RespondWithError(w, http.StatusInternalServerError, "Server Error")
		}
		return
	}

	errorhandlers.RespondWithJSON(w, http.StatusOK, locations)
}
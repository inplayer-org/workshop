package HandlersFunc

import (
	"net/http"
	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"strconv"
)

func (a *App) GetLocationByName (w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	name := vars["name"]
	id,err:=strconv.Atoi(name)
	/*if name != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid player name")
		return
	} */
	if err !=nil {
		panic(err)
	}

	player,err := queries.GetPlayersByLocation(a.DB,id)
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
		panic(err)
	}

	structures.Tmpl.ExecuteTemplate(w,"locs.html",locations)
}
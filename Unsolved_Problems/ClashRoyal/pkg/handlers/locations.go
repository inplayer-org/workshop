package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"strconv"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"
)
// Get location by name and return top250 players in that location by wins
func (a *App) GetLocationByName (w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	name := vars["name"]
	id,err:=strconv.Atoi(name)

	if err !=nil {
		panic(err)
	}
// querry from DB to list and sort 250 players from 1 location
	player,err := queries.GetPlayersByLocation(a.DB,id)

		if err!=nil {
			panic(err)
	}

	tmpl.Tmpl.ExecuteTemplate(w,"tableranking.html",player)

}
// Returning all locations from DB
func (a *App) GetLocations(w http.ResponseWriter, r *http.Request) {
// listing(returning) all locations from DB
	locations, err := queries.GetAllLocations(a.DB)

	if err != nil {
		panic(err)
	}

	tmpl.Tmpl.ExecuteTemplate(w,"locs.html",locations)

}
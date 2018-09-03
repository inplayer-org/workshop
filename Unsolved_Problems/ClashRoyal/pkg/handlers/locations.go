package handlers

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

	if err !=nil {
		panic(err)
	}

	player,err := queries.GetPlayersByLocation(a.DB,id)

		if err!=nil {
			panic(err)
	}

	structures.Tmpl.ExecuteTemplate(w,"tableranking.html",player)

}

func (a *App) GetLocations(w http.ResponseWriter, r *http.Request) {

	locations, err := queries.GetAllLocations(a.DB)

	if err != nil {
		panic(err)
	}

	structures.Tmpl.ExecuteTemplate(w,"locs.html",locations)

}
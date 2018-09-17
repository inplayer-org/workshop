package handlers

import (
	"net/http"

	"github.com/gorilla/mux"

	"strconv"

	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/locations"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/rankedPlayer"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"
)

// Get location by name and return top250 rankedPlayer in that location by wins
func (a *App) GetLocationByName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]
	id, err := strconv.Atoi(name)

	if err != nil {
		tmpl.Tmpl.ExecuteTemplate(w, "error.html", errors.NewResponseError("Location ID not number", "ID should contain only numbers", 404))
		return
	}
	// querry from DB to list and sort 250 rankedPlayer from 1 location
	player, _ := rankedPlayer.GetPlayersByLocation(a.DB, id)

	if len(player) == 0 {
		tmpl.Tmpl.ExecuteTemplate(w, "error.html", errors.NewResponseError("Location not found", "Location with "+strconv.Itoa(id)+" doesnot exist", 404))
		return
	}

	tmpl.Tmpl.ExecuteTemplate(w, "tableranking.html", player)

}
/*
// Returning all locations from DB
 func (a *App) GetLocations(w http.ResponseWriter, r *http.Request) {
 // listing(returning) all locations from DB
 	locs, err := locations.GetAllLocations(a.DB)

 	if err != nil {
 		tmpl.Tmpl.ExecuteTemplate(w,"error.html",errors.NewResponseError("Server error","Can't load locations something went wrong",503))
 		return
 	}

 	tmpl.Tmpl.ExecuteTemplate(w,"locs.html",locs)


 } */

// Returning all locations from DB
func (a *App) GetLocations(w http.ResponseWriter, r *http.Request) {
	// listing(returning) all locations from DB
	locs, err := locations.GetAllLocations(a.DB)

	if err != nil {
		tmpl.Tmpl.ExecuteTemplate(w, "error.html", errors.NewResponseError("Server error", "Can't load locations something went wrong", 503))
		return
	}
	m := make(map[string][]string)
	for _, j := range locs {
		if j.IsCountry {
			key, _ := parser.FirstCharFrom(j.Name)
			m[key] = append(m[key], j.Name)
		}
	}
	tmpl.Tmpl.ExecuteTemplate(w, "locationsNew.html", m)

}

package handlers



import (
	"net/http"
	"github.com/gorilla/mux"

	"strconv"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/players"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/locations"
)
// Get location by name and return top250 players in that location by wins
func (a *App) GetLocationByName (w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	name := vars["name"]
	id,err:=strconv.Atoi(name)

	if err !=nil {
		tmpl.Tmpl.ExecuteTemplate(w,"error.html",errors.NewResponseError("Location ID not number","ID should contain only numbers",404))
		return
	}
// querry from DB to list and sort 250 players from 1 location
	player,_ := players.GetPlayersByLocation(a.DB,id)

		if len(player)==0 {
			tmpl.Tmpl.ExecuteTemplate(w,"error.html",errors.NewResponseError("Location not found","Location with "+strconv.Itoa(id)+" doesnot exist",404))
			return
	}

	tmpl.Tmpl.ExecuteTemplate(w,"tableranking.html",player)

}
// Returning all locations from DB
func (a *App) GetLocations(w http.ResponseWriter, r *http.Request) {
// listing(returning) all locations from DB
	locations, err := locations.GetAllLocations(a.DB)

	if err != nil {
		tmpl.Tmpl.ExecuteTemplate(w,"error.html",errors.NewResponseError("Server error","Can't load locations something went wrong",503))
		return
	}

	tmpl.Tmpl.ExecuteTemplate(w,"locs.html",locations)

}
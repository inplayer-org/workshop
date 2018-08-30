package HandlersFunc

import (
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"database/sql"

	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errorhandlers"
)

func (a *App) GetPlayerByName (w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	name := vars["name"]
	/*if name != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid player name")
		return
	} */

	e := structures.PlayerStats{Name: name}
	if err := e.GetNamePlayer(a.DB); err != nil {

		switch err {
		case sql.ErrNoRows:
			errorhandlers.RespondWithError(w, http.StatusNotFound, "Player not found")
		default:
			errorhandlers.RespondWithError(w, http.StatusInternalServerError, "Server Error")
		}
		return
	}

	errorhandlers.RespondWithJSON(w, http.StatusOK, e)
}


func (a *App) Home(w http.ResponseWriter, r *http.Request) {


	players, err := queries.GetSortedRankedPlayers(a.DB,"wins",10)
	if err != nil {
		panic(err)
	}

structures.Tmpl.ExecuteTemplate(w,"home.html",players)
	}
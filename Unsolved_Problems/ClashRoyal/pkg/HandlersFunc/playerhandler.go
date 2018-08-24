package HandlersFunc

import (
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"database/sql"

	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/jsonExample/pkg/errorhandle"
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
			errorhandle.RespondWithError(w, http.StatusNotFound, "Player not found")
		default:
			errorhandle.RespondWithError(w, http.StatusInternalServerError, "Server Error")
		}
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusOK, e)
}

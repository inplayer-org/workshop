package HandlersFunc

import (
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"database/sql"

	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errorhandlers"
	"log"
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


func (a *App) GetPlayers(w http.ResponseWriter, r *http.Request) {


	players, err := structures.GetAllPlayers(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:

			errorhandlers.RespondWithError(w, http.StatusNotFound, "no players found")

		default:
			log.Println(err)
			errorhandlers.RespondWithError(w, http.StatusInternalServerError, "Server error")
		}
		return
	}

	errorhandlers.RespondWithJSON(w, http.StatusOK, players)
}
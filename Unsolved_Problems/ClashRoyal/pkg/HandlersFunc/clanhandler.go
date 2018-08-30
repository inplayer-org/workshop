package HandlersFunc

import (
	"net/http"
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errorhandlers"
)



/*func (a *App) GetClanByTag (w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	tag := (vars["tag"])
	/*if err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid clan tag")
		return
	}

	e := structures.Clan{Tag: tag}
	if err := e.GetTagClan(a.DB); err != nil {

		switch err {
		case sql.ErrNoRows:
			errorhandle.RespondWithError(w, http.StatusNotFound, "Clan not found")
		default:
			errorhandle.RespondWithError(w, http.StatusInternalServerError, "Server Error")
		}
		return
	}

	errorhandle.RespondWithJSON(w, http.StatusOK, e)
}
 */
func (a *App) GetClanByName (w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	name := (vars["name"])
	/*if err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid clan name")
		return
	} */

	e := structures.Clan{Name: name}
	if err := e.GetNameClan(a.DB); err != nil {

		switch err {
		case sql.ErrNoRows:
			errorhandlers.RespondWithError(w, http.StatusNotFound, "Clan not found")
		default:
			errorhandlers.RespondWithError(w, http.StatusInternalServerError, "Server Error")
		}
		return
	}

	errorhandlers.RespondWithJSON(w, http.StatusOK, e)
}


func (a *App) GetClans(w http.ResponseWriter, r *http.Request) {


	clans, err := structures.GetAllClans(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			errorhandlers.RespondWithError(w, http.StatusNotFound, "no clans found")

		default:
			errorhandlers.RespondWithError(w, http.StatusInternalServerError, "Server Error")
		}
		return
	}

	errorhandlers.RespondWithJSON(w, http.StatusOK, clans)
}
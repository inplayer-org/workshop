package HandlersFunc

import (
	"net/http"
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errorhandlers"
)



func (a *App) GetClanByTag (tag string, w http.ResponseWriter, r *http.Request){
	clan := structures.SelectClanByTag(a.DB,tag)
	if clan != nil {
		switch clan {
		case sql.ErrNoRows:
			errorhandlers.RespondWithError(w, http.StatusNotFound, "no clans found")
		default:
			errorhandlers.RespondWithError(w, http.StatusInternalServerError,"Server Error")
		}
		return
	}

	errorhandlers.RespondWithJSON(w, http.StatusOK, clan)
}

func (a *App) GetClanByName (name string,w http.ResponseWriter, r *http.Request){
	clan := structures.SelectClanByName(a.DB,name)
	if clan != nil {
		switch clan {
		case sql.ErrNoRows:
			errorhandlers.RespondWithError(w, http.StatusNotFound, "no clans found")
		default:
			errorhandlers.RespondWithError(w, http.StatusInternalServerError,"Server Error")
		}
		return
	}

	errorhandlers.RespondWithJSON(w, http.StatusOK, clan)
}


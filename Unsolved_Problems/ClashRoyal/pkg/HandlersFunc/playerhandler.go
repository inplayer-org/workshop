package HandlersFunc

import (
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"database/sql"

	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errorhandlers"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"log"
)

func (a *App) GetPlayerByName (w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	name := vars["name"]
	/*if err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid clan name")
		return
	} */

	players,err := queries.GetPlayersLike(a.DB,name)
	if err != nil {

		switch err {
		case sql.ErrNoRows:
			errorhandlers.RespondWithError(w, http.StatusNotFound, "Player not found")
		default:
			errorhandlers.RespondWithError(w, http.StatusInternalServerError, "Server Error")
		}
		return
	}
	log.Println(players)
	structures.Tmpl.ExecuteTemplate(w,"byplayersname.html",players)

}



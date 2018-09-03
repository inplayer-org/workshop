package HandlersFunc

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errorhandlers"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
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
	name := vars["name"]
	/*if err != nil {
		errorhandle.RespondWithError(w, http.StatusBadRequest, "Invalid clan name")
		return
	} */

	clans,err := queries.GetClansLike(a.DB,name)
	if err != nil {

		//Needs to be reworked into error template
		switch err {
		case sql.ErrNoRows:
			errorhandlers.RespondWithError(w, http.StatusNotFound, "Clan not found")
		default:
			errorhandlers.RespondWithError(w, http.StatusInternalServerError, "Server Error")
		}
		return
	}
	log.Println(clans)
	structures.Tmpl.ExecuteTemplate(w,"byclansname.html",clans)

	//errorhandlers.RespondWithJSON(w, http.StatusOK, e)
}


func (a *App) GetClans(w http.ResponseWriter, r *http.Request) {


	clans, err := queries.GetAllClans(a.DB)
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

func (a *App) UpdateClan(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	tag := vars["tag"]
	tag = "#"+tag
	log.Println("clanTag = ",tag)
	i := update.GetRequestForPlayersFromClan(a.DB,tag)
	if i == 404 {
		fmt.Println(http.StatusNotFound)
	} else {
		clanName, err := queries.GetClanName(a.DB,tag)
		if err != nil {
			panic(err)
		}
		log.Println("clanName = ", clanName)
		http.Redirect(w, r, "http://localhost:3303/clans/"+clanName+"/"+tag[1:], http.StatusTemporaryRedirect)
	}
}
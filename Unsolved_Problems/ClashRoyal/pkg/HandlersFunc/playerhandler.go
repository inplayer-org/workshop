package HandlersFunc

import (
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"github.com/gorilla/mux"
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
panic(err)
	}
	log.Println(players)
	structures.Tmpl.ExecuteTemplate(w,"byplayersname.html",players)

}


func (a *App) GetPlayerByTag(w http.ResponseWriter, r *http.Request){

	vars:=mux.Vars(r)
	tag:=vars["tag"]

	player,err:=queries.GetFromTag(a.DB,tag)

	if err!=nil {
		panic(err)
	}

	structures.Tmpl.ExecuteTemplate(w,"player.html",player)

}


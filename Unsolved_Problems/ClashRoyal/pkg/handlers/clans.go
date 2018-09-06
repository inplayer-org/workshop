package handlers

import (
	"fmt"
	"log"
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"
)
// Get clan by name from DB
func (a *App) GetClanByName (w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	name := vars["name"]

	clans,err := queries.GetClansLike(a.DB,name)
	if err != nil {
		//Needs to be reworked into error template
		panic(err)
	}

	tmpl.Tmpl.ExecuteTemplate(w,"byclansname.html",clans)

}

//Request tag to API to refres clan informations with members
func (a *App) UpdateClan(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	tag := vars["tag"]
	tag = "#"+tag
	log.Println("clanTag = ",tag)

	err := update.GetRequestForPlayersFromClan(a.DB,tag)

	if err!=nil {
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
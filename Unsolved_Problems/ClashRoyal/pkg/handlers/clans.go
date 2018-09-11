package handlers

import (
	"fmt"
	"log"
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
)
// Get clan by name from DB
func (a *App) GetClanByName (w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	name := vars["name"]
//querry for get clans by name
	clans,err := queries.GetClansLike(a.DB,name)
	if err != nil {
		//Needs to be reworked into error template
		tmpl.Tmpl.ExecuteTemplate(w,"error.html",err)
		return
	}

	fmt.Println(err)

	tmpl.Tmpl.ExecuteTemplate(w,"byclansname.html",clans)

}

//Request tag to API to refres clan informations with members
func (a *App) UpdateClan(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	tag := vars["tag"]

	tag = parser.ToHashTag(tag)
	log.Println("clanTag = ",tag)
// request for range over players in 1 clan
	err := update.GetRequestForPlayersFromClan(a.DB,tag)

	if err!=nil {
		fmt.Println(http.StatusNotFound)
	} else {
		// querry for clan name in DB
		clanName, err := queries.GetClanName(a.DB,tag)

		if err != nil {
			log.Println(err)
		}

		log.Println("clanName = ", clanName)
		http.Redirect(w, r, "http://localhost:3303/clans/"+clanName+"/"+tag[1:], http.StatusTemporaryRedirect)

	}
}
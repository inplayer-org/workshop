package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"
)

// Get clan by name from DB
func (a *App) GetClanByName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]
	//querry for get clans by name
	clans, err := queries.GetClansLike(a.DB, name)
	if err != nil {
		//Needs to be reworked into error template
		panic(err)
	}

	tmpl.Tmpl.ExecuteTemplate(w, "byclansname.html", clans)

}

//Request tag to API to refres clan informations with members
func (a *App) UpdateClan(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tag := vars["tag"]
	fmt.Println(tag)
	tag = "#" + tag
	log.Println("clanTag = ", tag)
	// request for range over players in 1 clan
	err := update.GetRequestForPlayersFromClan(a.DB, tag)

	if err != nil {
		fmt.Println(http.StatusNotFound)
	} else {
		// querry for clan name in DB
		clanName, err := queries.GetClanName(a.DB, tag)

		if err != nil {
			log.Println(err)
		}

		log.Println("clanName = ", clanName)
		http.Redirect(w, r, "http://localhost:3303/clans/"+clanName+"/"+tag[1:], http.StatusTemporaryRedirect)

	}
}

package handlers

import (
	"log"
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"
)
// Get clan by name from DB
func (a *App) GetClanByName (w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	name := vars["name"]

//querry for get clans by name
	clans,err := queries.GetClansLike(a.DB,name)

	if err != nil {
		tmpl.Tmpl.ExecuteTemplate(w,"error.html",errors.NewResponseError("Clan Name doesn't exist","There is no clan name like "+name,404))
		return
	}

	tmpl.Tmpl.ExecuteTemplate(w,"byclansname.html",clans)

}
// Sending string clan tag and response from DB clan informations
func (a *App)GetClanByTag(w http.ResponseWriter, r *http.Request){

	vars:=mux.Vars(r)
	tag:=vars["tag"]

	tag = parser.ToHashTag(tag)

	players,err:=queries.GetPlayersByClanTag(a.DB,tag)

	if err != nil {
		panic(err)
	}

	tmpl.Tmpl.ExecuteTemplate(w,"clan.html",players)

}
//Request tag to API to refres clan informations with members
func (a *App) UpdateClan(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	tag := vars["tag"]

	tag = parser.ToHashTag(tag)
	log.Println("clanTag = ",tag)
// request for range over players in 1 clan
	e := update.GetRequestForPlayersFromClan(a.DB,a.Client,tag)

	if e!=nil {
			tmpl.Tmpl.ExecuteTemplate(w,"error.html",e)
			return
	} else {
		// querry for clan name in DB
		clanName, err := queries.GetClanName(a.DB,tag)

		if err != nil {
			tmpl.Tmpl.ExecuteTemplate(w,"error.html",e)
			return
		}

		log.Println("clanName = ", clanName)
		http.Redirect(w, r, "http://localhost:3303/clans/"+clanName+"/"+tag[1:], http.StatusTemporaryRedirect)

	}
}
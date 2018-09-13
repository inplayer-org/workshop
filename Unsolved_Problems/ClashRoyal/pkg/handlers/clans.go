package handlers

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/clans"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/players"
)
// Get clan by name from DB
func (a *App) GetClanByName (w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	name := vars["name"]

//querry for get clans by name
	clans,err := clans.GetClansLike(a.DB,name)

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

	players,err:=findClan(a,tag)

	if len(players)==0 {
		tmpl.Tmpl.ExecuteTemplate(w,"error.html",err)
		return
	}

	tmpl.Tmpl.ExecuteTemplate(w,"clan.html",players)

}
func findClan(a *App,tag string) ([]players.RankedPlayer, error) {

	player,_:=players.GetPlayersByClanTag(a.DB,tag)

	if len(player)==0{

		clan,err:=a.Client.GetClan(tag)

		if err!=nil{
			return nil,err
		}

		err=clans.UpdateClans(a.DB,clan)

		if err!=nil{
			return nil,errors.NewResponseError(err.Error(), "Failed to insert/update the clan into database, please try again later", 500)
		}

		err=update.GetRequestForPlayersFromClan(a.DB,a.Client,clan.Tag)

		if err!=nil{
			return nil,errors.NewResponseError(err.Error(), "Can't find the clan", 404)
		}

		player,_=players.GetPlayersByClanTag(a.DB,tag)

	}

	return player,nil

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
		clanName, err := clans.GetClanName(a.DB,tag)

		if err != nil {
			tmpl.Tmpl.ExecuteTemplate(w,"error.html",e)
			return
		}

		log.Println("clanName = ", clanName)
		http.Redirect(w, r, "http://localhost:3303/clans/"+clanName+"/"+tag[1:], http.StatusTemporaryRedirect)

	}
}
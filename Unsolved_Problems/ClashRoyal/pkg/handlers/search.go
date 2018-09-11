package handlers

import (
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
)

func (a *App) Search(w http.ResponseWriter,r *http.Request){

	option := r.FormValue("searchby")
	text := r.FormValue("text")

	//log.Println("text = ", text, "option = ", option)

	if option=="playerName"{
		http.Redirect(w,r,"http://localhost:3303/players/"+text,http.StatusTemporaryRedirect)
	}
	if option=="playerTag" {

		name, err := queries.GetPlayerName(a.DB, text)

		if err == nil {
			http.Redirect(w, r, "http://localhost:3303/players/"+name+"/"+parser.ToRawTag(text), http.StatusTemporaryRedirect)
		}else{
			http.Redirect(w, r, "http://localhost:3303/players/new/"+parser.ToRawTag(text), http.StatusTemporaryRedirect)
		}
	}
	if option=="clanName"{
		http.Redirect(w,r,"http://localhost:3303/clans/"+text,http.StatusTemporaryRedirect)
	}
	if option=="clanTag" {
		clanName,err := queries.GetClanName(a.DB,text)
		if err!=nil{
				http.Redirect(w, r, "http://localhost:3303/clans/"+clanName+"/"+parser.ToRawTag(text), http.StatusTemporaryRedirect)
		}else{
			http.Redirect(w, r, "http://localhost:3303/clans/new/"+parser.ToRawTag(text), http.StatusTemporaryRedirect)
		}
	}
}
package handlers

import (
	"log"
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
	"fmt"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
)

func (a *App) Search(w http.ResponseWriter,r *http.Request){
	option := r.FormValue("searchby")
	text := r.FormValue("text")

	log.Println("text = ", text, "option = ", option)

	if option=="playerName"{
		http.Redirect(w,r,"http://localhost:3303/players/"+text,http.StatusTemporaryRedirect)
	}
	if option=="playerTag" {
		var name string
		name, err := queries.GetPlayerName(a.DB, text)

		if err == sql.ErrNoRows {
			i := update.GetRequestForPlayer(a.DB, parser.ToUrlTag(text))
			if i == 404 {
				fmt.Println(http.StatusNotFound)
			} else {
				name, err = queries.GetPlayerName(a.DB, text)
				if err != nil {
					panic(err)
				}
				log.Println("name = ", name)
				http.Redirect(w, r, "http://localhost:3303/players/"+name+"/"+text[1:], http.StatusTemporaryRedirect)
			}
		} else if err != nil {
			log.Println("Error = ", err)
			//Error
		} else {
			log.Println("name = ", name)
			http.Redirect(w, r, "http://localhost:3303/players/"+name+"/"+text[1:], http.StatusTemporaryRedirect)
		}
	}
	if option=="clanName"{
		http.Redirect(w,r,"http://localhost:3303/clans/"+text,http.StatusTemporaryRedirect)
	}
	if option=="clanTag" {
		clanName,err := queries.GetClanName(a.DB,text)
		if err==sql.ErrNoRows{
			i := update.GetRequestForPlayersFromClan(a.DB,text)     // koga kje sredis za clans so interface od client isto tuka moras da pormenis
			if i == 404 {
				fmt.Println(http.StatusNotFound)
			} else {
				clanName, err = queries.GetClanName(a.DB,text)
				if err != nil {
					panic(err)
				}
				log.Println("clanName = ", clanName)
				http.Redirect(w, r, "http://localhost:3303/clans/"+clanName+"/"+text[1:], http.StatusTemporaryRedirect)
			}
		} else if err != nil {
			log.Println("Error = ", err)
			//Error
		}else {
			log.Println("clanName = ", clanName)
			http.Redirect(w, r, "http://localhost:3303/clans/"+clanName+"/"+text[1:], http.StatusTemporaryRedirect)
		}
	}
}
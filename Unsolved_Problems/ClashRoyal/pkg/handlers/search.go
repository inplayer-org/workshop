package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/interface"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
)

func (a *App) Search(w http.ResponseWriter, r *http.Request) {
	option := r.FormValue("searchby")
	text := r.FormValue("text")

	//	tag:=parser.ToHashTag(text)

	log.Println("text = ", text, "option = ", option)

	client := _interface.NewClient()

	if option == "playerName" {
		http.Redirect(w, r, "http://localhost:3303/players/"+text[1:], http.StatusTemporaryRedirect)
	}
	if option == "playerTag" {
		var name string
		name, err := queries.GetPlayerName(a.DB, text)

		if err == sql.ErrNoRows {
			player, err := client.GetRequestForPlayer(text)
			if err != nil {
				fmt.Println(http.StatusNotFound)
			} else {

				var i int
				if player.LocationID != nil {
					i = player.LocationID.(int)
				} else {
					i = 0
				}

				err := queries.UpdatePlayer(a.DB, player, i)

				if err != nil {
					panic(err)
				} else {
					name, err = queries.GetPlayerName(a.DB, text)
					if err != nil {
						panic(err)
					}
					log.Println("name = ", name)
					http.Redirect(w, r, "http://localhost:3303/players/"+name+"/"+text[1:], http.StatusTemporaryRedirect)
				}
			}
		} else if err != nil {
			log.Println("Error = ", err)
			//Error
		} else {
			log.Println("name = ", name)
			http.Redirect(w, r, "http://localhost:3303/players/"+name+"/"+text[1:], http.StatusTemporaryRedirect)
		}
	}
	if option == "clanName" {
		http.Redirect(w, r, "http://localhost:3303/clans/"+text, http.StatusTemporaryRedirect)
	}
	if option == "clanTag" {
		clanName, err := queries.GetClanName(a.DB, text)
		if err == sql.ErrNoRows {
			e := update.GetRequestForPlayersFromClan(a.DB, text) // koga kje sredis za clans so interface od client isto tuka moras da pormenis
			if e != nil {
				fmt.Println(http.StatusNotFound)
			} else {
				clanName, err = queries.GetClanName(a.DB, text)
				if err != nil {
					panic(err)
				}
				log.Println("clanName = ", clanName)
				http.Redirect(w, r, "http://localhost:3303/clans/"+clanName+"/"+text[1:], http.StatusTemporaryRedirect)
			}
		} else if err != nil {
			log.Println("Error = ", err)
			//Error
		} else {
			log.Println("clanName = ", clanName)
			http.Redirect(w, r, "http://localhost:3303/clans/"+clanName+"/"+text[1:], http.StatusTemporaryRedirect)
		}
	}
}

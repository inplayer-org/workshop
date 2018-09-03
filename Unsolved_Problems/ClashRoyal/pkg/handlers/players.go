package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"log"
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
	"fmt"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"
)

func (a *App) GetPlayerByName (w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	name := vars["name"]

	players,err := queries.GetPlayersLike(a.DB,name)

	if err != nil {
		panic(err)
	}

	tmpl.Tmpl.ExecuteTemplate(w,"byplayersname.html",players)

}

func (a *App) GetPlayerByTag(w http.ResponseWriter, r *http.Request){

	vars:=mux.Vars(r)
	tag:=vars["tag"]

	player,err:=queries.GetFromTag(a.DB,tag)

	if err!=nil{
		if err==sql.ErrNoRows {
			t := "#" + tag
			i := update.GetRequestForPlayer(a.DB, parser.ToUrlTag(t))
			if i == 404 {
				fmt.Println(http.StatusNotFound)
				panic(err)
			} else {
				player, err := queries.GetFromTag(a.DB, tag)

				if err != nil {
					if err==sql.ErrNoRows{
						player,err:=queries.ClanNotFoundByTag(a.DB,tag)

						if err!=nil{
							panic(err)
						}

						fmt.Println(player)
						tmpl.Tmpl.ExecuteTemplate(w, "player.html", player)
						return
					}else {
					panic(err)
					}
				}

				tmpl.Tmpl.ExecuteTemplate(w, "player.html", player)
				return
			}
		}else{
			panic(err)
			}
	}else {
		tmpl.Tmpl.ExecuteTemplate(w, "player.html", player)
	}
}

func (a *App)GetPlayersByClanTag(w http.ResponseWriter, r *http.Request){

	vars:=mux.Vars(r)
	tag:=vars["tag"]

	players,err:=queries.GetPlayersByClanTag(a.DB,tag)

	if err != nil {
		panic(err)
	}

	tmpl.Tmpl.ExecuteTemplate(w,"clan.html",players)

}

func (a *App) UpdatePlayer(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	tag:=vars["tag"]

	t:="#"+tag
	i:=update.GetRequestForPlayer(a.DB, parser.ToUrlTag(t))

	if i==404{
		fmt.Println(http.StatusNotFound)
	}else{
		name, err := queries.GetPlayerName(a.DB, t)

		if err != nil {
			panic(err)
		}

		log.Println("name = ", name)
		http.Redirect(w, r, "http://localhost:3303/players/"+name+"/"+t[1:], http.StatusTemporaryRedirect)
	}
}
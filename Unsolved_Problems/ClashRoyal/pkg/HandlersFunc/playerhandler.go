package HandlersFunc

import (
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"log"
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
	"fmt"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
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

	//fmt.Println(err)

	if err!=nil{
		if err==sql.ErrNoRows {
			t := "#" + tag
			i := update.GetRequestForPlayer(a.DB, parser.ToUrlTag(t))
			if i == 404 {
				fmt.Println(http.StatusNotFound)
			} else {
				p, err := queries.GetFromTag(a.DB, tag)
				if err != nil {
					panic(err)
				}
				structures.Tmpl.ExecuteTemplate(w, "player.html", p)
				return
			}
		}else{
			fmt.Println(err)
			}
	}else {

		structures.Tmpl.ExecuteTemplate(w, "player.html", player)
	}
}

func (a *App)GetPlayersByClanTag(w http.ResponseWriter, r *http.Request){

	vars:=mux.Vars(r)
	tag:=vars["tag"]

	players,err:=queries.GetPlayersByClanTag(a.DB,tag)

	if err != nil {
		panic(err)
	}

	structures.Tmpl.ExecuteTemplate(w,"clan.html",players)

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
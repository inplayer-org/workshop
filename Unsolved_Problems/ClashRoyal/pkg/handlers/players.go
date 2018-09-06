package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"log"
	"database/sql"
	"fmt"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/interface"
)
// Sending Name as string to DB response Player by Name with all stats from PlayerStats
func (a *App) GetPlayerByName (w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	name := vars["name"]
//get players by name from DB
	players,err := queries.GetPlayersLike(a.DB,name)

	if err != nil {
		panic(err)
	}

	tmpl.Tmpl.ExecuteTemplate(w,"byplayersname.html",players)

}
//RequestTag -> sending to API response generated Player with playerstats struct and updating in DB
func (a *App) GetPlayerByTag(w http.ResponseWriter, r *http.Request){

	vars:=mux.Vars(r)

	tag:=vars["tag"]

	client := _interface.NewClient()

	player,err:=queries.GetFromTag(a.DB,tag)

	if err!=nil{
		if err==sql.ErrNoRows {
			t := "#" + tag
			player,err:= client.GetRequestForPlayer(t)

			//player cant be updated
			//moze da se staj od bazata so ima tova da dade ako nemoze da napraj req
			//poposle da sesredi
			if err!=nil {
				panic(err)
			}

			var i int

			if player.LocationID==nil{
				i=0
			}else{
				i=player.LocationID.(int)
			}

			err=queries.UpdatePlayer(a.DB,player,i)

			//nemoze da napraj insert ili update
			if err!=nil{
				log.Println(err)
			}

			player,err=queries.GetFromTag(a.DB,tag)

			//nemoze da go zapisha u databaza
			if err!=nil {
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


		}else{
			panic(err)
			}
	}else {
		tmpl.Tmpl.ExecuteTemplate(w, "player.html", player)
		return
	}
}
// Sending string clan tag and response from DB clan informations
func (a *App)GetPlayersByClanTag(w http.ResponseWriter, r *http.Request){

	vars:=mux.Vars(r)
	tag:=vars["tag"]

	players,err:=queries.GetPlayersByClanTag(a.DB,tag)

	if err != nil {
		panic(err)
	}

	tmpl.Tmpl.ExecuteTemplate(w,"clan.html",players)

}
// Sending RequestTag response hashtag .. cheking for player in db if not exist req to API and updating db
func (a *App) UpdatePlayer(w http.ResponseWriter, r *http.Request){
	vars:=mux.Vars(r)
	tag:=vars["tag"]
	client := _interface.NewClient()
	t:="#"+tag
	//sending request to API For 1 player if doesent exist in DB to update it
	player,err:=client.GetRequestForPlayer(t)

	if err !=nil {
		log.Println(err)
	}
	var i int

	if player.LocationID==nil{
		i=0
	}else{
		i=player.LocationID.(int)
	}

// querry to updateplayer from API To DB
	err=queries.UpdatePlayer(a.DB,player,i)

	if err!=nil{
		panic(err)
	}else{
		//querry to get PLayer name from DB
		name, err := queries.GetPlayerName(a.DB, t)

		if err != nil {
			panic(err)
		}

		log.Println("name = ", name)
		http.Redirect(w, r, "http://localhost:3303/players/"+name+"/"+t[1:], http.StatusTemporaryRedirect)
	}
}
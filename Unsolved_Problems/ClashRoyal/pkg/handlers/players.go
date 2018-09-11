package handlers

import (
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"
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

	t := parser.ToHashTag(tag)

	player,err:=findPlayer(a,t)

	if err != nil {
		tmpl.Tmpl.ExecuteTemplate(w,"error.html",err)
		return
	}

	tmpl.Tmpl.ExecuteTemplate(w, "player.html", player)

}
// Sending RequestTag response hashtag .. cheking for player in db if not exist req to API and updating db
func (a *App) UpdatePlayer(w http.ResponseWriter, r *http.Request){

	vars:=mux.Vars(r)
	tag:=vars["tag"]
	t:=parser.ToHashTag(tag)
	//sending request to API For 1 player if doesent exist in DB to update it
	player,err:=a.Client.GetRequestForPlayer(t)

	if err !=nil {
		log.Println(err)
	}

// querry to updateplayer from API To DB
	err=queries.UpdatePlayer(a.DB,player,nil)

	if err!=nil{
		panic(err)
	}else{
		//querry to get PLayer name from DB
		name, err := queries.GetPlayerName(a.DB, t)

		if err != nil {
			panic(err)
		}

		log.Println("name = ", name)
		http.Redirect(w, r, "http://localhost:3303/players/"+name+"/"+tag, http.StatusTemporaryRedirect)
	}
}

func findPlayer(a *App,tag string)(structures.PlayerStats,error) {

	player, err := queries.GetFromTag(a.DB, tag)

	if err != nil {
		//fmt.Println(err)
		if err == sql.ErrNoRows {

			player, err := a.Client.GetRequestForPlayer(tag)

			//poposle da sesredi
			if err != nil {
				return player, err
			}

			err = queries.UpdatePlayer(a.DB, player, nil)

			//nemoze da napraj insert ili update
			if err != nil {
				return player, errors.NewResponseError(err.Error(), "Failed to insert/update the player into database, please try again later", 500)
			}

			player, err = queries.GetFromTag(a.DB, tag)

			//nemoze da go zapisha u databaza
			if err != nil {
				if err == sql.ErrNoRows {
					player, err := queries.ClanNotFoundByTag(a.DB, tag)
					if err != nil {
						return player, errors.NewResponseError(err.Error(), "Can't find the player", 404)
					}
					return player, err
				} else {
					return player, errors.NewResponseError(err.Error(), "Unexpected error with the database", 500)
				}
			}

			return player, err

		} else {
			return player, errors.NewResponseError(err.Error(), "Unexpected error with the database", 500)
		}
	} else {
		return player, err
	}
}
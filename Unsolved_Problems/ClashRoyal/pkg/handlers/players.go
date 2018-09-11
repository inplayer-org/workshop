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

//Tries to find player in multiple steps, first in local database, then through clash royale api and returns error if it doesn't exist
func findPlayer(a *App,tag string)(structures.PlayerStats,error) {

	//Search for the player in local database
	player, err := queries.GetFromTag(a.DB, tag)

	if err != nil {

		//In case it doesn't exists in the database
		if err == sql.ErrNoRows {

			//Search for the player through the clash royale api
			player, err := a.Client.GetRequestForPlayer(tag)

			//Returns error if there was an error with the request to the api
			if err != nil {
				return player, err
			}

			//Updates the player in the local database
			err = queries.UpdatePlayer(a.DB, player, nil)

			//Error during the updating of the player in the local database
			if err != nil {
				return player, errors.NewResponseError(err.Error(), "Failed to insert/update the player into database, please try again later", 500)
			}

			//Reads the newly inserted player from the local database
			player, err = queries.GetFromTag(a.DB, tag)

			//Error during reading the newly inserted player from the database
			if err != nil {
				if err == sql.ErrNoRows {
					//Tries to read the same player without a clan
					player, err := queries.ClanNotFoundByTag(a.DB, tag)

					//In case the backup read without clan fails, player doesn't exists or isn't reachable at the moment
					if err != nil {
						return player, errors.NewResponseError(err.Error(), "Can't find the player", 404)
					}
					//Returns the structure of the player and a nil error
					return player, nil
				} else {
					//In case the database error isn't ErrNoRows, returns unexpected error
					return player, errors.NewResponseError(err.Error(), "Unexpected error with the database", 500)
				}
			}

			return player, nil

		} else {
			//In case the database error isn't ErrNoRows, returns unexpected error
			return player, errors.NewResponseError(err.Error(), "Unexpected error with the database", 500)
		}
	}

	//Returns the player struct and nil error if the player is present in our local database
		return player, nil

}
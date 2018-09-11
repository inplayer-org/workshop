package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/interface"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"
)

// Sending Name as string to DB response Player by Name with all stats from PlayerStats
func (a *App) GetPlayerByName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]
	//get players by name from DB
	players, err := queries.GetPlayersLike(a.DB, name)

	if err != nil {
		panic(err)
	}

	tmpl.Tmpl.ExecuteTemplate(w, "byplayersname.html", players)

}

//RequestTag -> sending to API response generated Player with playerstats struct and updating in DB
func (a *App) GetPlayerByTag(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	tag := vars["tag"]

	t := parser.ToHashTag(tag)
	log.Println("++++++++++++++++++++=")
	fmt.Println(t)

	client := _interface.NewClient()

	player, err := queries.GetFromTag(a.DB, t)

	if err != nil {
		//fmt.Println(err)
		if err == sql.ErrNoRows {

			player, err := client.GetRequestForPlayer(t)

			//player cant be updated
			//moze da se staj od bazata so ima tova da dade ako nemoze da napraj req
			//poposle da sesredi
			if err != nil {
				panic(err)

			}

			var i int

			if player.LocationID == nil {
				i = 0
			} else {
				i = player.LocationID.(int)
			}

			err = queries.UpdatePlayer(a.DB, player, i)

			//nemoze da napraj insert ili update
			if err != nil {
				log.Println(err)
			}

			player, err = queries.GetFromTag(a.DB, t)

			//nemoze da go zapisha u databaza
			if err != nil {
				if err == sql.ErrNoRows {
					player, err := queries.ClanNotFoundByTag(a.DB, t)

					if err != nil {
						panic(err)
					}

					fmt.Println(player)
					tmpl.Tmpl.ExecuteTemplate(w, "player.html", player)
					return
				} else {

					panic(err)
				}
			}

			tmpl.Tmpl.ExecuteTemplate(w, "player.html", player)
			return

		} else {
			panic(err)
		}
	} else {
		tmpl.Tmpl.ExecuteTemplate(w, "player.html", player)
		return
	}
}

// Sending string clan tag and response from DB clan informations
func (a *App) GetPlayersByClanTag(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	tag := vars["tag"]

	tag = parser.ToHashTag(tag)

	players, err := queries.GetPlayersByClanTag(a.DB, tag)

	if err != nil {
		panic(err)
	}

	tmpl.Tmpl.ExecuteTemplate(w, "clan.html", players)

}

// Sending RequestTag response hashtag .. cheking for player in db if not exist req to API and updating db
func (a *App) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tag := vars["tag"]
	t := parser.ToHashTag(tag)
	//sending request to API For 1 player if doesent exist in DB to update it
	player, err := a.Client.GetRequestForPlayer(t)

	if err != nil {
		log.Println(err)
	}

	// querry to updateplayer from API To DB
	err = queries.UpdatePlayer(a.DB, player, nil)

	if err != nil {
		panic(err)
	} else {
		//querry to get PLayer name from DB
		name, err := queries.GetPlayerName(a.DB, t)

		if err != nil {
			panic(err)
		}

		log.Println("name = ", name)
		http.Redirect(w, r, "http://localhost:3303/players/"+name+"/"+t[1:], http.StatusTemporaryRedirect)
	}
}

func (a *App) ComaprePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	player1 := vars["entry"]
	player2 := parser.ToRawTag(r.FormValue("player"))
	log.Println("bla", player1, player2)
	//p := r.PostFormValue("text")
	http.Redirect(w, r, "http://localhost:3303/compare/"+player1+"/"+player2, http.StatusTemporaryRedirect)

}

func (a *App) Compare2Players(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	player1 := parser.ToHashTag(vars["tag1"])
	player2 := parser.ToHashTag(vars["tag2"])
	log.Println("++++++++++++", player1, player2)
	var p structures.TwoPlayers
	p1, _ := queries.GetFromTag(a.DB, player1)
	p2, err := queries.GetFromTag(a.DB, player2)
	if err == sql.ErrNoRows {
		fmt.Println("ne postoi")
		http.Redirect(w, r, "http://localhost:3303/players/"+p1.Name+"/"+p1.Tag, http.StatusTemporaryRedirect)

	} else {
		p.Player2 = p2
		p.Player1 = p1
		tmpl.Tmpl.ExecuteTemplate(w, "compare.html", p)

	}

}

func (a *App) Compare(w http.ResponseWriter, r *http.Request) {
	var p structures.TwoPlayers
	var responseErr errors.ResponseError
	player1 := r.FormValue("player1")
	player2 := r.FormValue("player2")
	//log.Println("+++++++++++", player1, player2)
	p1, err1 := queries.GetFromTag(a.DB, player1)

	p2, err2 := queries.GetFromTag(a.DB, player2)
	if err1 == sql.ErrNoRows && err2 == sql.ErrNoRows {
		responseErr.Message = "404"
		responseErr.Reason = "Players NOT FOUND"
		tmpl.Tmpl.ExecuteTemplate(w, "error.html", responseErr)
	} else if err1 == sql.ErrNoRows {
		log.Println("err1")
		http.Redirect(w, r, "http://localhost:3303/players/"+p2.Name+"/"+p2.Tag, http.StatusTemporaryRedirect)
	} else if err2 == sql.ErrNoRows {
		log.Println("err2")
		http.Redirect(w, r, "http://localhost:3303/players/"+p1.Name+"/"+p1.Tag, http.StatusTemporaryRedirect)
	} else {
		log.Println("ok e")
		p.Player1 = p1
		p.Player2 = p2
		tmpl.Tmpl.ExecuteTemplate(w, "compare.html", p)
	}

}

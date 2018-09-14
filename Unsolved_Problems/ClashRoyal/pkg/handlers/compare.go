package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"
)

func (a *App) ComparePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	player1 := vars["entry"]
	player2 := parser.ToRawTag(r.FormValue("player"))
	log.Println("vo compare")
	http.Redirect(w, r, "http://localhost:3303/compare/"+player1+"/"+player2, http.StatusTemporaryRedirect)

}

//Compare the player with another player
func (a *App) Compare2Players(w http.ResponseWriter, r *http.Request) {
	var p structures.TwoPlayers

	vars := mux.Vars(r)
	player1 := parser.ToHashTag(vars["tag1"])
	player2 := parser.ToHashTag(vars["tag2"])
	p1, err := findPlayer(a, player1)
	p2, err := findPlayer(a, player2)
	if err != nil {
		http.Redirect(w, r, "http://localhost:3303/players/"+p1.Name+"/"+p1.Tag, http.StatusTemporaryRedirect)

	} else if player1 == player2 {
		http.Redirect(w, r, "http://localhost:3303/players/"+p1.Name+"/"+p1.Tag, http.StatusTemporaryRedirect)
	} else {
		p.Player2 = p2
		p.Player1 = p1
		tmpl.Tmpl.ExecuteTemplate(w, "compare.html", p)

	}
}

//Compare 2 Players
func (a *App) Compare(w http.ResponseWriter, r *http.Request) {
	var p structures.TwoPlayers

	player1 := r.FormValue("player1")
	player2 := r.FormValue("player2")
	/*Tries to find tle two players in multiple steps, first in local database, then through clash royale api
	and returns error if it doesn't exist*/
	p1, err1 := findPlayer(a, player1)
	p2, err2 := findPlayer(a, player2)
	//If the 2 Players does not exist returns that the PLayers does not exis
	if err1 != nil && err2 != nil {
		tmpl.Tmpl.ExecuteTemplate(w, "playerNotExist.html", errors.NewResponseError("Incorrect player tags", "Players NOT FOUND", 404))
	} else if err1 != nil {
		http.Redirect(w, r, "http://localhost:3303/players/"+p2.Name+"/"+p2.Tag, http.StatusTemporaryRedirect)
	} else if err2 != nil {
		http.Redirect(w, r, "http://localhost:3303/players/"+p1.Name+"/"+p1.Tag, http.StatusTemporaryRedirect)
	} else if player1 == player2 {
		http.Redirect(w, r, "http://localhost:3303/players/"+p1.Name+"/"+p1.Tag, http.StatusTemporaryRedirect)
	} else {
		p.Player1 = p1
		p.Player2 = p2
		tmpl.Tmpl.ExecuteTemplate(w, "compare.html", p)
	}

}

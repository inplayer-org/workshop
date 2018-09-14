package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/playerStats"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/twoPlayers"
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
	var p twoPlayers.TwoPlayers

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
	var p twoPlayers.TwoPlayers
	player1 := r.FormValue("player1")
	player2 := r.FormValue("player2")
	//log.Println("+++++++++++", player1, player2)
	p1, err1 := playerStats.GetFromTag(a.DB, player1)

	p2, err2 := playerStats.GetFromTag(a.DB, player2)
	if err1 == sql.ErrNoRows && err2 == sql.ErrNoRows {
		tmpl.Tmpl.ExecuteTemplate(w, "error.html", errors.NewResponseError("Incorrect player tags", "Players NOT FOUND", 404))
	} else if err1 == sql.ErrNoRows {
		log.Println("err1")
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

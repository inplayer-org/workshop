package handlers

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/players"
)

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
	var p players.TwoPlayers
	p1, _ := players.GetFromTag(a.DB, player1)
	p2, err := players.GetFromTag(a.DB, player2)
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
	var p players.TwoPlayers
	player1 := r.FormValue("player1")
	player2 := r.FormValue("player2")
	//log.Println("+++++++++++", player1, player2)
	p1, err1 := players.GetFromTag(a.DB, player1)

	p2, err2 := players.GetFromTag(a.DB, player2)
	if err1 == sql.ErrNoRows && err2 == sql.ErrNoRows {
		tmpl.Tmpl.ExecuteTemplate(w, "error.html", errors.NewResponseError("Incorrect player tags","Players NOT FOUND",404))
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
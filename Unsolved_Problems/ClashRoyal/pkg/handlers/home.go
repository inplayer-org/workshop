package handlers

import (
	"log"
	"net/http"

	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/rankedPlayer"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"
)

func (a *App) Home(w http.ResponseWriter, r *http.Request) {

	wins := "wins"
	players, err := rankedPlayer.GetSortedRankedPlayers(a.DB, wins, 10)

	if err != nil {
		tmpl.Tmpl.ExecuteTemplate(w, "error.html", errors.NewResponseError("Server error", "Can't load rankedPlayer something went wrong", 503))
		return
	}
	log.Println(players)
	tmpl.Tmpl.ExecuteTemplate(w, "homeNew.html", players)

}

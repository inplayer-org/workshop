package handlers

import (
	"net/http"


	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/players"
)

func (a *App) Home(w http.ResponseWriter, r *http.Request) {

	wins := "wins"
	players, err := players.GetSortedRankedPlayers(a.DB, wins, 10)

	if err != nil {
		tmpl.Tmpl.ExecuteTemplate(w,"error.html",errors.NewResponseError("Server error","Can't load players something went wrong",503))
		return
	}

	tmpl.Tmpl.ExecuteTemplate(w, "home.html", players)

}

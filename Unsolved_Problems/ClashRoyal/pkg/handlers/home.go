package handlers

import (
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"
)

func (a *App) Home(w http.ResponseWriter, r *http.Request) {

	wins := "wins"
	players, err := queries.GetSortedRankedPlayers(a.DB,wins,10)

	if err != nil {
		panic(err)
	}

	tmpl.Tmpl.ExecuteTemplate(w,"home.html",players)

}
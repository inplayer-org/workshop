package HandlersFunc

import (
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
)

func (a *App) Home(w http.ResponseWriter, r *http.Request) {

wins := "wins"
players, err := queries.GetSortedRankedPlayers(a.DB,wins,10)
if err != nil {
log.Println(err)
}
log.Println(players)
structures.Tmpl.ExecuteTemplate(w,"home.html",players)
}
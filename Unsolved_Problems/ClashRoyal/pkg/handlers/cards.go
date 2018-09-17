package handlers

import (
	"net/http"

	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/cards"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"
)

func (a *App) GetCards(w http.ResponseWriter, r *http.Request) {

	card, err := cards.GetAllCards(a.DB)

	if err != nil {
		tmpl.Tmpl.ExecuteTemplate(w, "error.html", errors.NewResponseError("Server error", "Can't load cards something went wrong", 503))
		return
	}

	tmpl.Tmpl.ExecuteTemplate(w, "cards.html", card)

}

package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
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

// Get card by name from DB
func (a *App) GetCardByName(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	name := vars["name"]

	//querry for get card by name
	cardsLike, err := cards.GetCardLike(a.DB, name)

	if err != nil {
		tmpl.Tmpl.ExecuteTemplate(w, "error.html", errors.NewResponseError("Card Name doesn't exist", "There is no card name like "+name, 404))
		return
	}

	tmpl.Tmpl.ExecuteTemplate(w, "bynamecards.html", cardsLike)

}

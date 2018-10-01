package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

)

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	//tmpl.Tmpl.ExecuteTemplate(w, "home.html", nil)
}

func (a *App) Search(w http.ResponseWriter, r *http.Request) {
	option := r.FormValue("searchby")
	text := r.FormValue("text")
	if option == "id" {
		http.Redirect(w, r, "http://localhost:3303", http.StatusTemporaryRedirect)
	} else {
		http.Redirect(w, r, "http://localhost:3303/members/"+text, http.StatusTemporaryRedirect)
	}

}

func (a *App) GetMemberByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tag := vars["tag"]
	log.Println(tag)
}



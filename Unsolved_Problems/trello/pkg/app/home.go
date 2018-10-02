package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"
)

func (a *App) Home(w http.ResponseWriter, r *http.Request) {
	tmpl.Tmpl.ExecuteTemplate(w, "login.html", nil)

}

func (a *App) Search(w http.ResponseWriter, r *http.Request) {
	// board := r.FormValue("board")
	// member := r.FormValue("member")
	// if len(board) > 0 && len(member) > 0 {
	// 	log.Println("you have to enter eather board or member ")
	// }
	// log.Println("member", member, board)

}

func (a *App) GetMemberByUsername(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	tag := vars["tag"]
	log.Println(tag)
}

func (a *App) LoginTry(w http.ResponseWriter, r *http.Request) {
	log.Println("tukasi")
	username := r.FormValue("username")
	password := r.FormValue("password")

	log.Println("print", username, password)

}

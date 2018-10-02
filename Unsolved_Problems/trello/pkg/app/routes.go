package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/",a.Home).Methods("GET")
	a.Router.HandleFunc("/search", a.Search).Methods("GET")

	a.Router.HandleFunc("/css/{cssName}",a.serveCSS)

	a.Router.HandleFunc("/members/{tag}", a.GetMemberByUsername).Methods("GET")


	a.Router.HandleFunc("/board",a.GetBoards).Methods("GET")
	a.Router.HandleFunc("/board/{id}",a.GetBoardByID).Methods("GET")

}


func (a *App) serveCSS(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	cssFileName := vars["cssName"]

	log.Println(cssFileName)

	filePath := "../tmpl/css/" + cssFileName + ".css"

	http.ServeFile(w, r, filePath)

}
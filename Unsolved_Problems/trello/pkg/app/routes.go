package app

import "net/http"

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/",a.Home).Methods("GET")
	a.Router.HandleFunc("/search", a.Search).Methods("GET")

	a.Router.HandleFunc("/css/",serveCSS)

	a.Router.HandleFunc("/members/{tag}", a.GetMemberByUsername).Methods("GET")


	a.Router.HandleFunc("/board",a.GetBoards).Methods("GET")
	a.Router.HandleFunc("/board/{id}",a.GetBoardByID).Methods("GET")

}


func serveCSS(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../tmpl/css/*.css")

}
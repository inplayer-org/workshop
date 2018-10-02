package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/",a.Home).Methods("GET")
	a.Router.HandleFunc("/search", a.Search).Methods("GET")

	a.Router.HandleFunc("/css/{cssName}.css",a.serveCSS)
	a.Router.HandleFunc("/img/{imageName}.png",a.serveImagePng)
	a.Router.HandleFunc("/img/{imageName}.jpg",a.serveImageJpg)

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

func (a *App) serveImagePng(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)

	imageFileName := vars["imageName"]

	log.Println(imageFileName,"PNG")

	filePath := "../tmpl/img/" + imageFileName + ".png"

	http.ServeFile(w, r, filePath)

}

func (a *App) serveImageJpg(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)

	imageFileName := vars["imageName"]

	log.Println(imageFileName,"JPG")

	filePath := "../tmpl/img/" + imageFileName + ".jpg"

	http.ServeFile(w, r, filePath)

}
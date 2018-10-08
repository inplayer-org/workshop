package app

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func (a *App) initializeRoutes() {


	a.Router.HandleFunc("/loginform",a.LoginForm).Methods("GET")
	a.Router.HandleFunc("/loginform",a.LogingIn).Methods("POST")

	//a.Router.HandleFunc("/registerform",a.RegisterForm).Methods("GET")
	a.Router.HandleFunc("/registerform" ,a.Registering).Methods("POST")

	a.Router.HandleFunc("/logoutform",a.LogOut).Methods("GET")

	a.Router.HandleFunc("/search", a.Search).Methods("GET")

	a.Router.HandleFunc("/boards",a.Boards).Methods("GET")

	a.Router.HandleFunc("/css/{cssName}.css",a.serveCSS)
	a.Router.HandleFunc("/img/{imageName}.png",a.serveImagePng)
	a.Router.HandleFunc("/img/{imageName}.jpg",a.serveImageJpg)
	a.Router.HandleFunc("/js/{jsName}.js",a.serveJS)

	a.Router.HandleFunc("/members/{tag}", a.GetMemberByUsername).Methods("GET")


	a.Router.HandleFunc("/board",a.GetBoards).Methods("GET")
	a.Router.HandleFunc("/board/{id}",a.GetBoardByID).Methods("GET")

	a.Router.HandleFunc("/",a.Home)

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

func (a *App) serveJS(w http.ResponseWriter, req *http.Request){

	vars:=mux.Vars(req)

	jsFileName:=vars["jsName"]

	filePath:="../tmpl/js/"+jsFileName+".js"

	http.ServeFile(w,req,filePath)

}
package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/user"
	"fmt"
)

func (a *App) Home(w http.ResponseWriter, req *http.Request) {
	c,err:=req.Cookie("sessions")
	if err!=nil{
		http.Redirect(w,req,"/loginform",303)
	}else {
		_, err := user.WhoAmI(c)
		u:=user.User{ID:12,Username:"Asd",Password:"asd",Token:""}
		if err != nil {
			http.Redirect(w, req, "/loginform", 303)
		}
		tmpl.ExecuteTemplate(w, "home.html", u.Username)
	}
}

func (a *App) Search(w http.ResponseWriter, r *http.Request) {
	// board := r.FormValue("board")
	// member := r.FormValue("member")
	// if len(board) > 0 && len(member) > 0 {
	// 	log.Println("you have to enter eather board or member ")
	// }
	// log.Println("member", member, board)

}

func (a *App) GetMemberByUsername(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	tag := vars["tag"]
	log.Println(tag)
}

func (a *App)LoginForm(w http.ResponseWriter,req *http.Request){

	tmpl.ExecuteTemplate(w,"login.html",nil)

}

func (a *App) LogingIn(w http.ResponseWriter, req *http.Request) {
	c,err:=req.Cookie("sessions")
	if err!=nil{
		c=&http.Cookie{
			Name:"sessions",
			Value:"test",
		}
		http.SetCookie(w,c)
	}
	fmt.Println(c)
	var p string



	username:=req.FormValue("username")
	password:=req.FormValue("password")

	//	dbSessions[c.Value]=username
	//	dbUsers[username]=password


	fmt.Println(username)
	p=password


	tmpl.ExecuteTemplate(w,"home.html",p)

}

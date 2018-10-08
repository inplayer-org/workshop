package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/user"
	"fmt"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/validators"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/session"
	"github.com/nu7hatch/gouuid"
)

func (a *App) Home(w http.ResponseWriter, req *http.Request) {
	fmt.Println("home" )
	c,err:=req.Cookie("sessions")
	if err!=nil{
		http.Redirect(w,req,"/loginform",303)
		return
	}else {
		u, err := user.WhoAmI(a.DB,c)

		if err != nil {
			http.Redirect(w, req, "/loginform", 303)
			return
		}
		tmpl.ExecuteTemplate(w, "home.html", u)
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
	c,err:=req.Cookie("sessions")

	if err!=nil {
		tmpl.ExecuteTemplate(w, "loginForm.html", nil)
	}else{
		_,err:=user.WhoAmI(a.DB,c)
		if err!=nil {
			tmpl.ExecuteTemplate(w, "loginForm.html", nil)
		}
		http.Redirect(w,req,"/",303)
	}

}

/*func (a *App)RegisterForm(w http.ResponseWriter,req *http.Request){
	c,err:=req.Cookie("sessions")

	if err!=nil {
		tmpl.ExecuteTemplate(w, "loginForm.html", nil)
	}else{
		_,err:=user.WhoAmI(a.DB,c)
		if err!=nil {
			tmpl.ExecuteTemplate(w, "loginForm.html", nil)
		}
		http.Redirect(w,req,"/",303)
	}

}*/

func (a *App)Registering(w http.ResponseWriter,req *http.Request){

	c,err:=req.Cookie("sessions")

	if err!=nil{
		cookieValue,_:=uuid.NewV4()
		c=&http.Cookie{
			Name:"sessions",
			Value:cookieValue.String(),
		}
		http.SetCookie(w,c)
	}
	fmt.Println(c)

	u,err:=user.WhoAmI(a.DB,c)

	if err!=nil{
		log.Println(err.Error())
	}
	exist,err:=validators.ExistsElementInColumn(a.DB,"Users",u.Username,"username")

	if err!=nil{
		log.Println(err.Error())
	}

	if exist{
		http.Redirect(w,req,"/",303)
	}
	username:=req.FormValue("username")
	password:=req.FormValue("password")

	u=user.User{Username:username,Password:password,Token:""}

	err=u.Insert(a.DB)
	if err!=nil{
		log.Println("inserting user error")
	}

	id,err:=user.GetUserID(a.DB,u.Username)

	if err!=nil{
		log.Println("error kaj get user ID")
	}

	s:=session.Session{UID:c.Value,IDuser:id}

	err=s.Insert(a.DB)
	if err!=nil{
		log.Println("inserting session error")
	}

	http.Redirect(w,req,"/",303)

}

func (a *App) LogingIn(w http.ResponseWriter, req *http.Request) {
	c,err:=req.Cookie("sessions")

	if err!=nil {
		cookieValue,_:=uuid.NewV4()
		c=&http.Cookie{
			Name:"sessions",
			Value:cookieValue.String(),
		}
		http.SetCookie(w,c)
	}

		u,err:=user.WhoAmI(a.DB,c)

		exist,_:=validators.ExistsElementInColumn(a.DB,"Users",u.Username,"username")

		if err!=nil || !exist{
			username:=req.FormValue("username")
			password:=req.FormValue("password")

			id,err:=user.GetUserID(a.DB,username)

			if err!=nil{
				http.Redirect(w,req,"/loginform",303)
				return
			}

			s:=session.Session{UID:c.Value,IDuser:id}
			err=s.Insert(a.DB)

			if err!=nil{
				http.Redirect(w,req,"/loginform",303)
				return
			}

			p,err:=user.GetUserPassword(a.DB,username)
			if p!=password{

				http.Redirect(w,req,"/loginform",303)
				return
			}
			http.Redirect(w,req,"/",303)
			return
		}
		http.Redirect(w,req,"/",303)
		return


}

func (a *App)LogOut(w http.ResponseWriter,req *http.Request){
fmt.Println("vleguva" )
	c,_:=req.Cookie("sessions")

	fmt.Println(c)

	u,_:=user.WhoAmI(a.DB,c)

	id,_:=user.GetUserID(a.DB,u.Username)

	s:=session.Session{UID:c.Value,IDuser:id}

	err:=s.Delete(a.DB)

	fmt.Println(err)

tmpl.ExecuteTemplate(w,"loginForm.html",nil)
	}
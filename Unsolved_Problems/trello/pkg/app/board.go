package app

import (
	"net/http"
	"fmt"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/boards"

	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/user"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/members"
)


func (a *App) GetBoards(w http.ResponseWriter, r *http.Request) {

	board, err := boards.GetAllBoards(a.DB)

	//error handler

	if err != nil {
		//tmpl.Tmpl.ExecuteTemplate(w, "error.html")
		panic(err)
	}

//	tmpl.Tmpl.ExecuteTemplate(w, "board.html", board)

fmt.Println(board)
}






func (a *App) GetBoardByID (w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	id := vars["id"]

	//querry for get board by ID
	board,err := boards.GetBoardbyID(a.DB,id)

	if err != nil {

		fmt.Println(err)
	}

//	tmpl.Tmpl.ExecuteTemplate(w,"board.html",board)
	fmt.Println(board)

}

type MemberBoards struct{
	m members.Member
	b []boards.Board
}

func (a *App) Boards (w http.ResponseWriter,req *http.Request){

	c,err:=req.Cookie("session")

	if err!=nil{
		http.Redirect(w,req,"/loginform",302)
	}

	var mb MemberBoards

	u,err:=user.WhoAmI(a.DB,c)

	if err!=nil{
		http.Redirect(w,req,"/loginform",302)
	}

	membber,err:=members.GetMemberWithUsername(a.DB,u.Username)

	mb.m=membber

}
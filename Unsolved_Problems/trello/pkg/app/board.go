package app

import (
	"net/http"
	"fmt"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/boards"

	"github.com/gorilla/mux"
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
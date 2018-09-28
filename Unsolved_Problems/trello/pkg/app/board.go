package app

import (

	"net/http"
	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/tmpl"

	"net/url"
)

//RequestTag -> sending to API response generated Player with playerstats struct and updating in DB
func (a *App) GetBoardByID(w http.ResponseWriter, r *http.Request){

	vars:=mux.Vars(r)

	id:=vars["id"]


	err:=RegisterFromBoard(a,id)

	if err != nil {
		tmpl.Tmpl.ExecuteTemplate(w,"error.html",err)
		return
	}

	tmpl.Tmpl.ExecuteTemplate(w, "board.html", err)

}


func RegisterFromBoard(a *App,id string)error{


	board,err := a.Client.GetBoard(id)

	if err!=nil{
		return err
	}

	err = board.Update(a.DB)


	if err!=nil{
		return err
	}

	var q *url.URL

	q.Query().Add("key","9ecdc5f04a4ccb643b83d4fd2b920416")
	q.Add("token","125a712b04063b34d2c22392704bb38a5fc88bb48f665c0f6bdf2d516f473c9d")

	cards,err := a.Client.GetBoard()



}

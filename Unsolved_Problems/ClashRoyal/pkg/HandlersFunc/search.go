package HandlersFunc

import (
	"log"
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/queries"
)

func (a *App) Search(w http.ResponseWriter,r *http.Request){
	option := r.FormValue("searchby")
	text := r.FormValue("text")

	if option=="clans"{
		http.Redirect(w,r,"http://localhost:3303/clans/"+text,http.StatusTemporaryRedirect)
	}
	if option=="playerName"{
		http.Redirect(w,r,"http://localhost:3303/players/"+text,http.StatusTemporaryRedirect)
	}
	if option=="playerTag"{
		name,err := queries.GetPlayerName(a.DB,text)
		if err!=nil{
			log.Println("Error = ",err)
			//Error
		}
		log.Println("name = ",name)
		http.Redirect(w,r,"http://localhost:3303/players/"+name+"/"+text,http.StatusTemporaryRedirect)
	}
	log.Println("text = ",text,"option = ",option)
}
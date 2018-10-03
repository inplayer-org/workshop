package user

import (
	"net/http"
	"repo.inplayer.com/workshop/Unsolved_Problems/trello/pkg/session"
	"database/sql"
)

type User struct{
	ID int
	Username string
	Password string
	Token string
}

func WhoAmI(DB *sql.DB,cookie *http.Cookie)(User,error){

	var u User

	IDuser,err:=session.GetFromSeassions(DB,cookie.Value)

	if err!=nil{
		return u,err
	}

	u,err=GetFromUser(DB,IDuser.IDuser)

	return u,nil
}

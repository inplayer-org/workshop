package user

import "net/http"

type User struct{
	ID int
	Username string
	Password string
	Token string
}

func WhoAmI(cookie *http.Cookie)(User,error){

}

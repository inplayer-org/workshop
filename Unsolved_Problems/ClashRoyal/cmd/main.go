package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/handlers"
)


func handleErr(err error){
	if(err!=nil){
		log.Println(err)
	}
}


//enterFlags flags for DbName UserName and Password
func enterFlags() (string,string,string) {

	DbName := flag.String("database", "Clash_Royale", "the name of you database")

	UserName := flag.String("username", "root", "the username to make a connection to the database")

	Password := flag.String("password", "12345", "the password for your username to make a conection to the database")

	flag.Parse()

	return *DbName,*UserName,*Password
}

func main (){

	dbName,userName,password:=enterFlags()

	connectionString := fmt.Sprintf("%s:%s@/%s", userName, password, dbName)
	fmt.Println("Connection string =",connectionString)
	db,err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)    }

	router := mux.NewRouter()

	var app handlers.App

	app.Initialize(db,router)


}
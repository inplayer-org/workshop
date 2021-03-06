package main

import (
	"database/sql"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/PrevozMk/pkg/errors"
	"repo.inplayer.com/workshop/Unsolved_Problems/PrevozMk/pkg/handlers"
)

//enterFlags flags for DbName UserName and Password
func enterFlags() (string,string,string) {

	DbName := flag.String("database", "Prevoz_Mk", "the name of you database")

	UserName := flag.String("username", "root", "the username to make a connection to the database")

	Password := flag.String("password", "12345", "the password for your username to make a conection to the database")

	flag.Parse()

	return *DbName,*UserName,*Password
}


//Starting the database,opening router and performing listen and serve
func main (){

	//Database information through flags
	dbName,userName,password:=enterFlags()

	//Opening connection to the database
	connectionString := fmt.Sprintf("%s:%s@/%s", userName, password, dbName)
	fmt.Println("Connection string =",connectionString)
	db,err := sql.Open("mysql", connectionString)

	//Panic if there is a problem with the database since whole web app isn't functional and is dependent on a connection to the database
	err = errors.Database(err)
	if err != nil {
		panic(err)
	}

	//Creating a router
	router := mux.NewRouter()

	//App structure for connected database and router
	var app handlers.App

	//Open the routes and perform listen and serve
	app.Initialize(db,router)


}
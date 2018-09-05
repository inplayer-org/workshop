package main

import (
	"fmt"
	"database/sql"
	"flag"
	"repo.inplayer.com/workshop/Unsolved_Problems/testing/pkg/interface"
	"log"
)

func enterFlags() (string,string,string) {

	DbName := flag.String("database", "Clash_Royale", "the name of you database")

	UserName := flag.String("username", "root", "the username to make a connection to the database")

	Password := flag.String("password", "12345", "the password for your username to make a conection to the database")

	flag.Parse()

	return *DbName,*UserName,*Password
}

func main(){

	dbName,userName,password:=enterFlags()

	connectionString := fmt.Sprintf("%s:%s@/%s", userName, password, dbName)
	fmt.Println("Connection string =",connectionString)
	db,err := sql.Open("mysql", connectionString)

	if err != nil {
		log.Println(err)
	}

	client:=_interface.NewClient(db)

	fmt.Println("sadfg")

	playerTags,err:=client.GetPlayerTagFromClans("#8YL8802Y")

	fmt.Println(playerTags)

	if err!=nil{
		log.Println(err)
	}

	for i,elem:=range playerTags.Player{
		fmt.Println(i,elem.Tag)
	}
}

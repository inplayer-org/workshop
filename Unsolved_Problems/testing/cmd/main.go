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

	locs,err:=client.GetLocations()

	fmt.Println(locs)

	if err != nil {
		log.Println(err)
	}

	for _,elem:=range locs.Location {
		fmt.Println(elem.Name, elem.ID)
		tags,err:=client.GetPlayerTagsFromLocation(elem.ID)

		if err != nil {
			log.Println(err)
		}

		for _,t:=range tags.Player{
			fmt.Println(t.Tag)
		}
	}



}

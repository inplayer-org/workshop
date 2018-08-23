package main

import (
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/locations"
	"fmt"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/gettagbyclans"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/sortplayers"
	"flag"
	"database/sql"
	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/routeranddb"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/get"
)


//enterFlags flags for DbName UserName and Password
func enterFlags() (string,string,string) {

	DbName := flag.String("database", "demodb", "the name of you database")

	UserName := flag.String("username", "root", "the username to make a conection to the database")

	Password := flag.String("password", "12345", "the password for your username to make a conection to the database")
	flag.Parse()


	return *DbName,*UserName,*Password
}

func main (){

	dbName,userName,password:=enterFlags()

	connectionString := fmt.Sprintf("%s:%s@/%s", userName, password, dbName)

	db,err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)    }

	router := mux.NewRouter()

	var app routeranddb.App

	app.Initialize(db,router)

	//ushte da se koristi bazata

	loc,err:=locations.GetLocations()

	if err!=nil {
		panic(err)
	}

	locationsMap:=locations.LocationMap(loc)

	mkdID:=locationsMap["Albania"]

	fmt.Println(mkdID)

	playerTags,err:=locations.GetPlayerTagsPerLocation(mkdID)

	if err!=nil{
		panic(err)
	}

	tagsFromLoc:=parser.ToUrlTags(playerTags.GetTags())

	//fmt.Println(tags)

	tagsFromClan := get.GetTagByClans(parser.ToUrlTag("#2LQGYRV"))

	sortplayers.ByWins(tagsFromClan)

	sortplayers.ByWins(tagsFromLoc)

}
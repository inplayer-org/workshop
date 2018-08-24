package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/get"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/locations"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
)

func handleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

//enterFlags flags for DbName UserName and Password
func enterFlags() (string, string, string) {

	DbName := flag.String("database", "Clash_Royale", "the name of you database")

	UserName := flag.String("username", "root", "the username to make a connection to the database")

	Password := flag.String("password", "12345", "the password for your username to make a conection to the database")

	flag.Parse()

	return *DbName, *UserName, *Password
}

func dailyUpdate(db *sql.DB) {
	log.Println("Updating all locations data")
	allLocations, err := locations.DailyUpdateLocations(db)
	log.Println("Finished updating locations data")
	handleErr(err)
	for _, elem := range allLocations.Location {
		playerTags, err := locations.GetPlayerTagsPerLocation(elem.ID)

		handleErr(err)
		if elem.IsCountry {
			log.Println("Updating players for country -> ", elem.Name)
			update.Players(db, parser.ToUrlTags(playerTags.GetTags()), elem.ID)
		}
	}
	allClans, err := update.GetAllClans(db)
	handleErr(err)
	log.Println("Refreshing data for all clans present in the database")
	for _, elem := range allClans {
		clan := get.GetTagByClans(elem.Tag)
		log.Println("Updating players for clan ->", elem.Name)
		update.Players(db, parser.ToUrlTags(clan), 0)
	}
}

func main() {

	dbName, userName, password := enterFlags()

	connectionString := fmt.Sprintf("%s:%s@/%s", userName, password, dbName)
	fmt.Println("Connection string =", connectionString)
	db, err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	var app routeranddb.App

	app.Initialize(db, router)

	dailyUpdate(db)

	//ushte da se koristi bazata

	//loc,err:=locations.GetLocations()
	//
	//if err!=nil {
	//	panic(err)
	//}
	//
	//locationsMap:=locations.LocationMap(loc)
	//
	//mkdID:=locationsMap["Albania"]
	//
	//fmt.Println(mkdID)
	//
	//playerTags,err:=locations.GetPlayerTagsPerLocation(mkdID)
	//
	//if err!=nil{
	//	panic(err)
	//}
	//
	//tagsFromLoc:=parser.ToUrlTags(playerTags.GetTags())
	//
	////fmt.Println(tags)
	//
	//tagsFromClan := get.GetTagByClans(parser.ToUrlTag("#2LQGYRV"))
	//
	//sortplayers.ByWins(tagsFromClan)
	//
	//sortplayers.ByWins(tagsFromLoc)

}

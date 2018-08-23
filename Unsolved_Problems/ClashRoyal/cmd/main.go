package main

import (
	"fmt"
	"flag"
	"database/sql"
	"github.com/gorilla/mux"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/routeranddb"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/locations"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
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


func dailyUpdate(db *sql.DB){
	allLocations,err := locations.DailyUpdateLocations(db)
	handleErr(err)
	for _,elem := range allLocations.Location{
		playerTags,err := locations.GetPlayerTagsPerLocation(57000007)

		handleErr(err)
		if elem.IsCountry {
			log.Println("Updating for country -> ",elem.Name)
			update.Players(db, parser.ToUrlTags(playerTags.GetTags()), 57000007)
		}
	}
}


func main (){

	dbName,userName,password:=enterFlags()

	connectionString := fmt.Sprintf("%s:%s@/%s", userName, password, dbName)
	fmt.Println("Connection string =",connectionString)
	db,err := sql.Open("mysql", connectionString)

	if err != nil {
		panic(err)    }

	router := mux.NewRouter()

	var app routeranddb.App

	app.Initialize(db,router)

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
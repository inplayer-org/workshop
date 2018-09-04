package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/structures"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
)

func handleErr(err error){
	if err!=nil{
		log.Println(err)
	}
}

var  locationsCounter int

func enterFlags() (string,string,string) {

	DbName := flag.String("database", "Clash_Royale", "the name of you database")

	UserName := flag.String("username", "root", "the username to make a connection to the database")

	Password := flag.String("password", "12345", "the password for your username to make a conection to the database")

	flag.Parse()

	return *DbName,*UserName,*Password
}

func main() {

	dbName,userName,password:=enterFlags()

	connectionString := fmt.Sprintf("%s:%s@/%s", userName, password, dbName)
	fmt.Println("Connection string =",connectionString)
	db,err := sql.Open("mysql", connectionString)

	locationInfoChan := make(chan structures.Locationsinfo,300)
	defer close(locationInfoChan)

	done := make(chan string,300)
	defer close(done)

	locationsCounter = 0


	//Section 1 - Update for locations table
	log.Println("Updating all locations data")
	allLocations,err := update.DailyUpdate(db)
	log.Println("Finished updating locations data")
	handleErr(err)

	//Section 2 - Update players from locations table

	//Starting Workers
	for i:=0;i<40;i++{
		go PlayerWorker(db,locationInfoChan,done)

	}

	//Sending Locations to workers
		for _, location := range allLocations.Location {

			if location.IsCountry {

				locationInfoChan <- location
				locationsCounter++

			}
		}


	log.Println("Ready for information through done for locations ...")

		//Waiting the responses from the workers
	for i:=0;i<locationsCounter;i++{

		log.Println("Finished Updating for location ",<-done )

	}

	log.Println("\n finished with the locations update")
	}





func PlayerWorker(db *sql.DB,locationInfoChan <- chan structures.Locationsinfo,done chan <- string){


	for location :=  range locationInfoChan {

		playerTags, err := update.GetPlayerTagsPerLocation(location.ID)

		handleErr(err)

		log.Println("Sending players for country -> ", location.Name)

		allErrors := update.Players(db, parser.ToUrlTags(playerTags.GetTags()), location.ID)

		for _, err := range allErrors {
			log.Println(err)
		}
		done <- location.Name
	}
}
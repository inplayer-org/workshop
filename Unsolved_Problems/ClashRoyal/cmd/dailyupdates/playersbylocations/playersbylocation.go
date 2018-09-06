package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/cmd/dailyupdates/pkg/workers"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
)

func handleErr(err error){
	if err!=nil{
		log.Println(err)
	}
}


func enterFlags() (string,string,string) {

	DbName := flag.String("database", "Clash_Royale", "the name of you database")

	UserName := flag.String("username", "root", "the username to make a connection to the database")

	Password := flag.String("password", "darko123", "the password for your username to make a conection to the database")

	flag.Parse()

	return *DbName,*UserName,*Password
}

func main() {

	dbName,userName,password:=enterFlags()

	connectionString := fmt.Sprintf("%s:%s@/%s", userName, password, dbName)
	fmt.Println("Connection string =",connectionString)
	db,err := sql.Open("mysql", connectionString)

	locationInfoChan := make(chan workers.Worker)
	defer close(locationInfoChan)

	done := make(chan string)
	defer close(done)


	//Section 1 - Update for locations table
	log.Println("Updating all locations data")
	allLocations,err := update.DailyUpdate(db)
	log.Println("Finished updating locations data")
	handleErr(err)

	//Section 2 - Update players from locations table

	//Starting Workers
	for i:=0;i<40;i++{
		go workers.StartWorker(db,locationInfoChan,done)

	}

	//Sending Locations to workers
		go func(){
			for _, location := range allLocations.Location {

				locationInfoChan <- workers.NewLocationWorker(location)

			}
		}()


	log.Println("Ready for information through done for locations ...")

		//Waiting the responses from the workers
	for i:=0;i<len(allLocations.Location);i++{

		log.Println("Finished Updating for location ",<-done,",",len(allLocations.Location)-i,"locations left to update" )

	}

	log.Println("Finished with the locations update")
	}

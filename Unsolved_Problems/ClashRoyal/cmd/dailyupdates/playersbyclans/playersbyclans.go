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

	Password := flag.String("password", "12345", "the password for your username to make a conection to the database")

	flag.Parse()

	return *DbName,*UserName,*Password
}


/*Daily update using clans present in our database

It updates all of the players from all of clans that are already present in our database that have been registered
either though the web app or the the location daily update

-Should be started after finishing the Daily update using the locations since the players gotten from
location api might have new clans that aren't already present in our database.

*/
func main() {

	//Database information through flags
	dbName,userName,password:=enterFlags()

	//Opening connection to the databse
	connectionString := fmt.Sprintf("%s:%s@/%s", userName, password, dbName)
	fmt.Println("Connection string =",connectionString)
	db,err := sql.Open("mysql", connectionString)

	//Channel for sending Clans to the Workers
	clanInfoChan := make(chan workers.Worker)
	defer close(clanInfoChan)

	//Channel for counting finished jobs
	done := make(chan string)
	defer close(done)

	//Starting Workers
	for i:=0;i<40;i++{
		go workers.StartWorker(db,clanInfoChan,done)

	}

	//Getting all clans present in the database
	allClans,err := update.GetAllClans(db)
	handleErr(err)
	log.Println("Refreshing data for all clans present in the database")

	//Sending Clans to active workers
	go func(){
		for _,clan := range allClans{

			clanInfoChan <- workers.NewClanWorker(clan)

		}
	}()

	log.Println("Ready for information through done for clans ...")

	//Waiting the responses from the workers
	for i:=0;i<len(allClans);i++{

		log.Println("Finished Updating for clan ",<-done,",",len(allClans)-i,"clans left to update" )

	}

	log.Println("Finished with the clans update")
}


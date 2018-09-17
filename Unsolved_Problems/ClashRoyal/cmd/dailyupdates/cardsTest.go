package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/errors"

	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
)

func handleErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func enterFlags() (string, string, string) {

	DbName := flag.String("database", "Clash_Royale", "the name of you database")

	UserName := flag.String("username", "root", "the username to make a connection to the database")

	Password := flag.String("password", "12345", "the password for your username to make a conection to the database")

	flag.Parse()

	return *DbName, *UserName, *Password
}

/*Daily update using all locations present in clash royale api

It updates information in our database about all locations available through the clash royale api and after that
updates information in our database about all top rankedPlayer from every location that is a country and has ranking present for it

-Should be started first and finish before update playersbyclans is started (ideally)

*/
func main() {

	//Database information through flags
	dbName, userName, password := enterFlags()

	//Opening connection to the database
	connectionString := fmt.Sprintf("%s:%s@/%s", userName, password, dbName)
	fmt.Println("Connection string =", connectionString)
	db, err := sql.Open("mysql", connectionString)

	err = errors.Database(err)
	if err != nil {
		panic(err)
	}

	cards, _ := update.CardsUpdate(db)

	log.Println("Finished with the locations update", cards)
}

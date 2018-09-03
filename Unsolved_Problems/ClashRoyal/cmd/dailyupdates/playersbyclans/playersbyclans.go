package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/parser"
	"repo.inplayer.com/workshop/Unsolved_Problems/ClashRoyal/pkg/update"
	"time"
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

func main() {

	dbName,userName,password:=enterFlags()

	connectionString := fmt.Sprintf("%s:%s@/%s", userName, password, dbName)
	fmt.Println("Connection string =",connectionString)
	db,err := sql.Open("mysql", connectionString)


	done := make(chan interface{})
	defer close(done)

	start := make(chan bool)
	defer close(start)

	isStarted := false

	countFinished := 0

	//Section 3 - Update players from clans table
	allClans,err := update.GetAllClans(db)
	handleErr(err)
	log.Println("Refreshing data for all clans present in the database")
	go func() {
		for _, elem := range allClans {
			for ; countFinished >= 40; {
				time.Sleep(time.Second * 5)
			}
			fmt.Println(elem.Tag)
			clan := update.GetTagByClans(parser.ToUrlTag(elem.Tag))
			log.Println("Updating players for clan ->", elem.Name)
			countFinished++
			go update.Players(db, clan, 0, done)
			if isStarted == false {
				isStarted = true
				start <- true
			}
		}
	}()
	<-start
	log.Println("Ready for information through done for clans ...")
	for ;countFinished>0;countFinished--{
		log.Println("Finished Updating for clan ",<-done )
	}
}
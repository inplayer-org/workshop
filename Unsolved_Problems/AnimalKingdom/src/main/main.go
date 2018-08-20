package main

import (
	"flag"

	_ "github.com/go-sql-driver/mysql"
)

func processFlags() (string, string) {
	var connectString, port string
	flag.StringVar(&connectString, "db-connect", "root:1111@tcp(127.0.0.1:3306)/animalKingdom", "DB Connect String")
	flag.StringVar(&port, "port", ":8888", "HTTP listen spec port")

	flag.Parse()
	return connectString, port
}

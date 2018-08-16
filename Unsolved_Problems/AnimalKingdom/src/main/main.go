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

// func main() {
// 	cfg := processFlags()
// 	d := db.ConnectDB(cfg)
// 	fmt.Println(d)
// 	a := db.SelectAnimal(d, 2)
// 	fmt.Println(a)
// 	//b := db.SelectAllFood(d)
// 	animal := structures.Animal{Name: "an1", Species: "jjjjj", Height: 12}
// 	db.InsertAnimal(d, animal)
// 	//fmt.Println(animal)
// 	s := db.SelectFoodAnimalEat(d, "hdjsds")
// 	fmt.Println(s)
// 	ss := db.SelectAnimalsEatCertainFood(d, "mesojadno")
// 	fmt.Println(ss)
// }

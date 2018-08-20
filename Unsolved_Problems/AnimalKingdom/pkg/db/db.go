package db

import (
	"database/sql"
	"fmt"
	"strconv"

	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/controllinput"
	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/structures"
)

func errorHandler(err error) {
	if err != nil {
		fmt.Println("ERROR :", err)
	}
}

func errorExistant(err error) bool{
	if err != nil{
		fmt.Println("ERROR :", err)
		return true
	}
	return false
}

//Func that return if the entry is name or id
func NameOrID(entry string) (interface{}, string) {
	if controllinput.IntOnly(entry) {
		i, _ := strconv.Atoi(entry)
		return i, "animalID"
	}
	return entry, "name"
}

//ConnectDB creates a database object
func ConnectDB(connectString string) *sql.DB {
	dbConn, err := sql.Open("mysql", connectString)
	errorHandler(err)

	return dbConn

}

//SelectAllAnimals return slice of all animals
func SelectAllAnimals(db *sql.DB) []structures.Animal {
	rows, err := db.Query("SELECT * FROM Animal")
	errorHandler(err)
	animals := []structures.Animal{}
	for rows.Next() {
		var name, species string
		var id, height int
		err := rows.Scan(&id, &name, &species, &height)
		errorHandler(err)
		animal := structures.Animal{AnimalID: id, Species: species, Height: height, Name: name}
		animals = append(animals, animal)
	}
	return animals
}

//Select Animal by name or id
func SelectAnimal(db *sql.DB, entry string) (structures.Animal, error) {
	var name, species string
	var id, height int
	value, what := NameOrID(entry)
	err := db.QueryRow("SELECT * FROM Animal WHERE "+what+"=(?)", value).Scan(&id, &name, &species, &height)
	animal := structures.Animal{AnimalID: id, Species: species, Height: height, Name: name}
	return animal, err
}

//Delete animal by name or id
func DeleteAnimal(db *sql.DB, entry string) error {
	value, what := NameOrID(entry)
	delAnimal, err := db.Prepare("DELETE FROM Animal WHERE " + what + "=(?)")
	delAnimal.Exec(value)
	return err
}




//Insert Animal with structure Animal
func InsertAnimal(db *sql.DB, animal structures.Animal) error {
	a := animal
	_, err := db.Exec("INSERT INTO Animal(name,species,height) VALUES (?,?,?)", a.Name, a.Species, a.Height)
	if err != nil {
		panic(err.Error)
	}
	return err
}

//Update Animal by name
func UpdateAnimal(db *sql.DB, animal structures.Animal) (structures.Animal, error) {
	a := animal
	update, err := db.Prepare("UPDATE Animal set species=(?),height=(?) WHERE name=(?)")

	update.Exec(a.Species, a.Height, a.Name)
	return a, err
}



//Select all Food that an Animal Eat
func SelectFoodAnimalEat(db *sql.DB, animalName string) []string {
	rows, err := db.Query("SELECT Food.foodName from Eat inner join Food on Eat.foodID=Food.foodID inner join Animal on Animal.animalID=Eat.animalID WHERE Animal.name=(?)", animalName)
	errorHandler(err)
	var listOfFood []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		errorHandler(err)
		listOfFood = append(listOfFood, name)
	}
	return listOfFood

}




//Insert eat with Structue Animal and Structure Food
func InsertEat(db *sql.DB, animal structures.Animal, food structures.Food) {
	f := food
	a := animal
	_, err := db.Exec("INSERT INTO Eat(animalID,foodID)  VALUES (?,?)", a.AnimalID, f.FoodID)
	errorHandler(err)
	//return err
}




func Exists(db *sql.DB, name string) string {
	var exists string
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM Animal WHERE name=(?))", name).Scan(&exists)
	errorHandler(err)
	return exists
}
func ExistsID(db *sql.DB, id int) string {
	var exists string
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM Animal WHERE animalID=(?))", id).Scan(&exists)
	errorHandler(err)
	return exists
}

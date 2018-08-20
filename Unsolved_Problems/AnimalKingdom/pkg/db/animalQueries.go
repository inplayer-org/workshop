package db

import (
	"database/sql"
	"fmt"
	"regexp"
	"strconv"

	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/structures"
)

func errorHandler(err error) {
	if err != nil {
		fmt.Println("ERROR :", err)
	}
}

var IsInt = regexp.MustCompile(`^[0-9+]+$`).MatchString

//Func that return if the entry is name or id
func NameOrID(entry string) (interface{}, string) {
	if IsInt(entry) {
		i, _ := strconv.Atoi(entry)
		return i, "animalID"
	}
	return entry, "name"
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
	//errorHandler(err)
	return animal, err
}

//Delete animal by name or id
func DeleteAnimal(db *sql.DB, entry string) error {
	value, what := NameOrID(entry)
	delAnimal, _ := db.Prepare("DELETE FROM Animal WHERE " + what + "=(?)")
	_, err := delAnimal.Exec(value)
	return err
}

//Insert Animal with structure Animal
func InsertAnimal(db *sql.DB, animal structures.Animal) error {
	a := animal
	_, err := db.Exec("INSERT INTO Animal(name,species,height) VALUES (?,?,?)", a.Name, a.Species, a.Height)

	return err
}

//Update Animal by name
func UpdateAnimal(db *sql.DB, animal structures.Animal) (structures.Animal, error) {
	a := animal
	update, err := db.Prepare("UPDATE Animal set species=(?),height=(?) WHERE name=(?)")

	update.Exec(a.Species, a.Height, a.Name)
	return a, err
}

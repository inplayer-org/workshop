package db

import (
	"database/sql"
	"fmt"

	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/structures"
)

func errorHandler(err error) {
	if err != nil {
		fmt.Println("ERROR :", err)
	}
}

func errorExistant(err error)bool{
	if err != nil {
		fmt.Println("ERROR :", err)
		return true
	}
	return false
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
		animal := structures.Animal{}
		err := rows.Scan(&id, &name, &species, &height)
		errorHandler(err)
		animal.Name = name
		animal.Species = species
		animal.Height = height
		animal.AnimalID = id
		animals = append(animals, animal)
	}
	return animals
}

//Select Animal by name
func SelectAnimal(db *sql.DB, animalName string) (structures.Animal, error) {
	var name, species string
	var id, height int
	animal := structures.Animal{}
	err := db.QueryRow("SELECT * FROM Animal WHERE name=(?)", animalName).Scan(&id, &name, &species, &height)
	animal.AnimalID = id
	animal.Name = name
	animal.Species = species
	animal.Height = height
	return animal, err
}

//SelectAllFood return slice of all food
func SelectAllFood(db *sql.DB) []structures.Food {
	rows, err := db.Query("SELECT * FROM Food")
	errorHandler(err)
	listOfFood := []structures.Food{}
	for rows.Next() {
		var typeFood, name string
		var ID int
		food := structures.Food{}
		err := rows.Scan(&ID, &typeFood,&name)
		errorHandler(err)
		food.FoodID = ID
		food.Name = name
		food.Type = typeFood
		listOfFood = append(listOfFood, food)
	}
	return listOfFood
}

//Select Food by ID return structures.Food
func SelectFoodbyID(db *sql.DB, foodID int) interface{} {
	var typeFood, name string
	var ID int
	food := structures.Food{}
	err := db.QueryRow("SELECT * FROM Food WHERE foodID=(?)", foodID).Scan(&ID, &typeFood,&name)
	if errorExistant(err){
		return nil
	}
	food.FoodID = foodID
	food.Name = name
	food.Type = typeFood
	return food
}

func SelectFoodbyName(db *sql.DB, foodName string) interface{} {
	var typeFood, name string
	var ID int
	food := structures.Food{}
	err := db.QueryRow("SELECT * FROM Food WHERE foodName=(?)", foodName).Scan(&ID, &typeFood,&name)
	if errorExistant(err){
		return nil
	}
	food.FoodID = ID
	food.Name = name
	food.Type = typeFood
	return food
}

//InsertAnimal with structure Animal
func InsertAnimal(db *sql.DB, animal structures.Animal) {
	a := animal
	_, err := db.Exec("INSERT INTO Animal(name,species,height) VALUES (?,?,?)", a.Name, a.Species, a.Height)
	errorHandler(err)
}

//DeleteAnimal
func DeleteAnimal(db *sql.DB, animalName string) {
	delAnimal, err := db.Prepare("DELETE FROM Animal WHERE name=(?)")
	errorHandler(err)
	delAnimal.Exec(animalName)
}

//UpdateAnimal
func UpdateAnimal(db *sql.DB, animal structures.Animal) {
	a := animal
	update, err := db.Prepare("UPDATE Animal set species=(?),height=(?) WHERE name=(?)")
	errorHandler(err)
	update.Exec(a.Species, a.Height, a.Name)
}

//InsertFood with structure Food
func InsertFood(db *sql.DB, food structures.Food) {
	f := food
	_, err := db.Exec("INSERT INTO Food(foodID,foodName,type)  VALUES (?,?,?)",f.FoodID, f.Name, f.Type)
	errorHandler(err)
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

//Select all Animals that eat certain food
func SelectAnimalsEatCertainFood(db *sql.DB, foodName string) []string {
	rows, err := db.Query("SELECT Animal.name from Eat inner join Food on Eat.foodID=Food.foodID inner join Animal on Animal.animalID=Eat.animalID WHERE Food.foodName=(?)", foodName)
	errorHandler(err)
	var animals []string
	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		errorHandler(err)
		animals = append(animals, name)
	}
	return animals

}

//Select all food that have certain type
func SelectAllFoodByType(db *sql.DB, typeFood string) []string {
	rows, err := db.Query("SELECT foodName, count(foodName) FROM Food WHERE type=(?) GROUP BY foodName", typeFood)
	errorHandler(err)
	var food []string
	for rows.Next() {
		var name string
		var count int
		err := rows.Scan(&name, &count)
		errorHandler(err)
		food = append(food, name)
	}
	return food

}

//Insert eat with Structue Animal and Structure Food
func InsertEat(db *sql.DB, animal structures.Animal, food structures.Food) {
	f := food
	a := animal
	_, err := db.Exec("INSERT INTO Eat(animalID,foodID)  VALUES (?,?)", a.AnimalID, f.FoodID)
	errorHandler(err)
}

package db

import (
	"database/sql"

	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/structures"
)

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

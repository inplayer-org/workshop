package db

import (
	"database/sql"
	"repo.inplayer.com/workshop/Unsolved_Problems/AnimalKingdom/pkg/structures"
	"log"
)



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

func SelectFoodbyID(db *sql.DB, foodID int) interface{} {
	var typeFood, name string
	var ID int
	err := db.QueryRow("SELECT * FROM Food WHERE foodID=(?)", foodID).Scan(&ID, &typeFood,&name)
	if errorExistant(err){
		return nil
	}
	return structures.Food{FoodID:ID,Name:name,Type:typeFood}
}

func SelectFoodbyName(db *sql.DB, foodName string) interface{} {
	var typeFood, name string
	var ID int
	err := db.QueryRow("SELECT * FROM Food WHERE foodName=(?)", foodName).Scan(&ID, &typeFood,&name)
	if errorExistant(err){
		return nil
	}
	return structures.Food{FoodID:ID,Name:name,Type:typeFood}
}

func DeleteFoodbyName(db *sql.DB, foodName string) error{
	delFood, err := db.Prepare("DELETE FROM Food WHERE foodName=(?)")
	errorHandler(err)
	_,err = delFood.Exec(foodName)
	log.Println(err)
	return err
}

func DeleteFoodbyID(db *sql.DB, foodID int) error{
	delFood, err := db.Prepare("DELETE FROM Food WHERE foodID=(?)")
	errorHandler(err)
	_,err = delFood.Exec(foodID)
	log.Println(err)
	return err
}

func InsertFood(db *sql.DB, food structures.Food) error{
	f := food
	_, err := db.Exec("INSERT INTO Food(foodID,foodName,type)  VALUES (?,?,?)",f.FoodID, f.Name, f.Type)
	log.Println(err)
	return err
}

func UpdateFood(db *sql.DB, food structures.Food) (structures.Food, error) {
	update, err := db.Prepare("UPDATE Food set foodName=(?),type=(?) WHERE foodID=(?)")
	update.Exec(food.Name,food.Type,food.FoodID)
	if err!=nil{
		return structures.Food{},err
	}
	return food,nil
}

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
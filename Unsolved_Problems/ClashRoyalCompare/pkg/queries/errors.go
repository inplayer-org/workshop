package queries

import "fmt"

func UpdateError(err error){
	if(err!=nil){
		fmt.Println("Error while updating the table -> ",err)
	}
}

func InsertError(err error){
	if(err!=nil){
		fmt.Println("Error while inserting in the table -> ",err)
	}
}

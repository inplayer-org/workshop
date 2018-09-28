package validators

import (
	"database/sql"
	"fmt"
)

func ExistsElementInColumn(db *sql.DB,tableName string,element interface{},column string)(bool,error){

	var exists int

	query := fmt.Sprintf("SELECT Exists (SELECT * FROM %s WHERE %s=?);",tableName,column)

	err := db.QueryRow(query,element).Scan(&exists)

	if err!=nil{
		return false,err
	}

	if exists == 1{
		return true,nil
	}

	return false,nil

}


func ExistsRelation(db *sql.DB,tableName string,ForeignKey1 interface{},ForeignKey2 interface{},columnName1 string,columnName2 string)(bool,error){
	var exists int

	query := fmt.Sprintf("SELECT Exists (SELECT * FROM %s WHERE %s=? && %s=?);",tableName,columnName1,columnName2)

	err := db.QueryRow(query,ForeignKey1,ForeignKey2).Scan(&exists)

	if err!=nil{
		return false,err
	}

	if exists == 1{
		return true,nil
	}

	return false,nil
}
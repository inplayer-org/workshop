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

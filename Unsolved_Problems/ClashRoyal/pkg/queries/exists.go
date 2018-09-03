package queries

import (
	"database/sql"
	"fmt"
)

func Exists(DB *sql.DB,table string,column string,value string) bool{

	var result int

	query:=fmt.Sprintf(`SELECT COUNT(%s) FROM %s WHERE %s="%s"`,column,table,column,value)

	DB.QueryRow(query).Scan(&result)

	if result==0{
		return false
	}

	return true

}
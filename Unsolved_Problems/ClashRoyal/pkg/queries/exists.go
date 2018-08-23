package queries

import "database/sql"

func Exists(DB *sql.DB,table string,column string,value string) bool{

	var result int

	DB.QueryRow(`SELECT * FROM (?) where (?)="(?)"`,table,column,value).Scan(&result)

	if result==0{
		return false
	}

	return true

}

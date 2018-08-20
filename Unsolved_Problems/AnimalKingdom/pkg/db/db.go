package db

import (
	"database/sql"
)

//ConnectDB creates a database object
func ConnectDB(connectString string) *sql.DB {
	dbConn, err := sql.Open("mysql", connectString)
	errorHandler(err)

	return dbConn

}

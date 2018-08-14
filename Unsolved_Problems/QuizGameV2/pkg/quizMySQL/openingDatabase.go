package quizmysql

import (
	"database/sql"
	"fmt"

	qInput "repo.inplayer.com/workshop/Unsolved_Problems/QuizGameV2/pkg/quizInput"
)

func errorHandler(err error) {
	if err != nil {
		fmt.Println("ERROR :", err)
	}
}

func parseDataSourceName(password string) string {
	dataSourceName := "root:" + password + "@tcp(127.0.0.1:3306)/QuizGame"
	return dataSourceName
}

func createDataSourceName() string {
	fmt.Printf("Enter your user for the root account of the database (Dino or Elena):")
	password := qInput.DetermineDataBase()

	return parseDataSourceName(string(password))
}

//DataBaseConnect - asks for password on root and opens a connection
func DataBaseConnect() (db *sql.DB) {

	//Creating data source name with entering password for accessing the database
	dataSourceName := createDataSourceName()

	db, err := sql.Open("mysql", dataSourceName)
	errorHandler(err)
	return db
}

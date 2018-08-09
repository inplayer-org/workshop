package insertIntoHighScores

import (
	"database/sql"
	"fmt"
)

func errorHandler(err error) {
	if err != nil {
		fmt.Println("ERROR :", err)
	}
}

//Insert current name and score of the PLayer in table HighScores
func InsertIntoHighScores(db *sql.DB, name string, score int) {
	_, err := db.Exec("INSERT INTO HighScores(score,name) VALUES (?,?)", score, name)
	errorHandler(err)
}

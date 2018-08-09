package quizmysql

import (
	"database/sql"
	"fmt"
)

type player struct {
	Name  string
	Score int
}

//GetTop10 - Return a slice of the top10 Players sorted by score
func getTop10(db *sql.DB) []player {
	top10, err := db.Query("SELECT score,name FROM HighScores ORDER BY score DESC LIMIT 10")
	errorHandler(err)
	players := []player{}
	for top10.Next() {
		var name string
		var score int
		player := player{}
		err := top10.Scan(&score, &name)
		errorHandler(err)
		player.Score = score
		player.Name = name
		players = append(players, player)
	}
	return players
}

//InsertIntoHighScores - Insert current name and score of the Player in table HighScores
func InsertIntoHighScores(db *sql.DB, name string, score int) {
	_, err := db.Exec("INSERT INTO HighScores(score,name) VALUES (?,?)", score, name)
	errorHandler(err)
}

//PrintTop10 - Print the top 10 players present in the database in a sorted order
func PrintTop10(db *sql.DB) {
	players := getTop10(db)
	fmt.Println("\nTOP 10 PLAYERS :")
	fmt.Println("----------------------------------------")
	fmt.Printf("|%-32s|%-5s|\n", "Name", "Score")
	fmt.Println("------+---------------------------------")
	for _, j := range players {
		fmt.Printf("|%-32s|%-5d|\n", j.Name, j.Score)
	}
	fmt.Println("----------------------------------------")
}

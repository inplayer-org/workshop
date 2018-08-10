package quizmysql

import (
	"database/sql"
	"fmt"
)

//Player - structure that hold the player name, his score and his rank
type Player struct {
	Name  string
	Score int
	Rank  int
}

//GetTop10 - Return a slice of the top10 Players sorted by score
func GetTop10(db *sql.DB) []Player {
	top10, err := db.Query("SELECT score,name FROM HighScores ORDER BY score DESC LIMIT 10")
	errorHandler(err)
	playerRank := 1
	players := []Player{}
	for top10.Next() {
		var name string
		var score int
		player := Player{}
		err := top10.Scan(&score, &name)
		errorHandler(err)
		player.Score = score
		player.Name = name
		player.Rank = playerRank
		players = append(players, player)
		playerRank++
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
	players := GetTop10(db)
	fmt.Println("\nTOP 10 PLAYERS :")
	fmt.Println("----------------------------------------")
	fmt.Printf("|%-32s|%-5s|\n", "Name", "Score")
	fmt.Println("------+---------------------------------")
	for _, j := range players {
		fmt.Printf("|%-32s|%-5d|\n", j.Name, j.Score)
	}
	fmt.Println("----------------------------------------")
}

//GetTop10PlaysOfUser - return the top 10 scores by a user in a sorted and ranked order
func GetTop10PlaysOfUser(db *sql.DB, user string) []Player {
	top10, err := db.Query("SELECT score,name FROM HighScores WHERE name=(?) ORDER BY score DESC LIMIT 10", user)
	errorHandler(err)
	playerRank := 1
	players := []Player{}
	for top10.Next() {
		var name string
		var score int
		player := Player{}
		err := top10.Scan(&score, &name)
		errorHandler(err)
		if name != user {
			continue
		}
		player.Score = score
		player.Name = name
		player.Rank = playerRank
		players = append(players, player)
		playerRank++
	}
	return players
}

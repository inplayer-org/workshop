package getTop10

import (
	"database/sql"
	"fmt"
)

type Player struct {
	Name  string
	Score int
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println("ERROR :", err)
	}
}

// Get top 10 Players and returning slice of top10 Players used for func PrintTop10
func GetTop10(db *sql.DB) []Player {
	top10, err := db.Query("SELECT score,name FROM HighScores ORDER BY score DESC LIMIT 10")
	errorHandler(err)
	players := []Player{}
	for top10.Next() {
		var name string
		var score int
		player := Player{}
		err := top10.Scan(&score, &name)
		errorHandler(err)
		player.Score = score
		player.Name = name
		players = append(players, player)
	}
	return players
}

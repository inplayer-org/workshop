package printTop10

import (
	"database/sql"
	"fmt"

	qGetTop10 "repo.inplayer.com/workshop/Unsolved_Problems/QuizGameV2/mySql/getTop10"
)

//Printing top10 Players in terminal

func PrintTop10(db *sql.DB) {
	players := qGetTop10.GetTop10(db)
	fmt.Println("\nTOP 10 PLAYERS :")
	fmt.Println("----------------------------------------")
	fmt.Printf("|%-32s|%-5s|\n", "Name", "Score")
	fmt.Println("------+---------------------------------")
	for _, j := range players {
		fmt.Printf("|%-32s|%-5d|\n", j.Name, j.Score)
	}
	fmt.Println("----------------------------------------")
}

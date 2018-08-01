package quizPrinting

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type publicUser struct {
	userId      int
	userName    string
	fullName    string
	highScore   int
	currentRank int
}

func PrintErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func openBase(client string) *sql.DB {
	dataBase, err := sql.Open("mysql", client)
	if err != nil {
		fmt.Println("ERROR WITH DATABASE !!!", err)
		os.Exit(3)
	}
	return dataBase

}

func PrintPublicUser(user_id int) {

	user := createUserStruct(user_id)
	fmt.Printf("\n\tAccount Details:\nUser ID : %d\nUsername : %s\nFull Name : %s\nHigh Score : %d\nCurrent Rank : %d\n\n", user.userId, user.userName, user.fullName, user.highScore, user.currentRank)
}

func PrintTop10() {
	connectionAccount := "root:112234@/quiz_game_base"
	dataBase := openBase(connectionAccount)
	top10UsersRows, err := dataBase.Query("SELECT user_id FROM high_scores ORDER BY high_score DESC LIMIT 10")
	PrintErr(err)
	for top10UsersRows.Next() {
		var nextId int
		err = top10UsersRows.Scan(&nextId)
		PrintErr(err)
		PrintPublicUser(nextId)
	}

}

func createUserStruct(user_id int) publicUser {
	connectionAccount := "root:112234@/quiz_game_base"
	dataBase := openBase(connectionAccount)
	userIdRow := dataBase.QueryRow("SELECT `username`,`full_name`,high_scores.`high_score`,(SELECT ranking FROM (SELECT (@row_number := @row_number +1)AS ranking, high_score,user_id FROM (SELECT @row_number :=0) AS temp, high_scores ORDER BY high_score DESC) AS T WHERE user_id=(?)) AS ranking FROM users JOIN high_scores ON high_scores.`user_id`=users.`user_id` WHERE users.`user_id`=(?);", user_id, user_id)
	var highS, currentR int
	var userN, fullN string
	err := userIdRow.Scan(&userN, &fullN, &highS, &currentR)
	PrintErr(err)
	userCreation := publicUser{user_id, userN, fullN, highS, currentR}
	return userCreation
}

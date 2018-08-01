package quizPrinting

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"

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

func PrintRankedUser(user_id int) {
	user := createUserStruct(user_id)
	rank := strconv.Itoa(user.currentRank)
	fmt.Println()
	print := "|RANK : " + rank + " |"
	for range print {
		fmt.Printf("-")
	}
	fmt.Println()
	fmt.Println(print)
	PrintPublicUser(user_id)
}

func PrintPublicUser(user_id int) {

	user := createUserStruct(user_id)
	fmt.Println("|-----------------------------|")
	fmt.Println("|       Account Details       |")
	fmt.Printf("|User ID | %-11d        |\n", user.userId)
	fmt.Printf("|Username | %-16s  |\n", user.userName)
	fmt.Printf("|Full Name | %-16s |\n", user.fullName)
	fmt.Printf("|High Score | %-16d|\n", user.highScore)
	fmt.Println("|-----------------------------|")

}

func PrintTop10() {
	connectionAccount := "root:112234@/quiz_game_base"
	dataBase := openBase(connectionAccount)
	fmt.Println("\nRANKINGS :")
	top10UsersRows, err := dataBase.Query("SELECT user_id FROM high_scores ORDER BY high_score DESC LIMIT 10")
	PrintErr(err)
	fmt.Println("----------------------------------------")
	fmt.Printf("|%-5s|%-32s|\n", "Rank", "Full Name")
	fmt.Println("------+---------------------------------")
	for top10UsersRows.Next() {
		var nextId int
		err = top10UsersRows.Scan(&nextId)
		PrintErr(err)
		user := createUserStruct(nextId)
		fmt.Printf("|%-5d|%-32s|\n", user.currentRank, user.fullName)
	}
	fmt.Println("----------------------------------------")
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

func ListAllUsers() {
	connectionAccount := "root:112234@/quiz_game_base"
	dataBase := openBase(connectionAccount)
	listOfUsersRows, err := dataBase.Query("SELECT user_id FROM users")
	PrintErr(err)
	fmt.Printf("---------------------------------------------------\n")
	fmt.Printf("|%-16s|%-32s|\n", "Username", "Full Name")
	fmt.Printf("|----------------+--------------------------------|\n")
	for listOfUsersRows.Next() {
		var nextId int
		err = listOfUsersRows.Scan(&nextId)
		PrintErr(err)
		user := createUserStruct(nextId)
		fmt.Printf("|%-16s|%-32s|\n", user.userName, user.fullName)
	}
	fmt.Printf("---------------------------------------------------\n")
}

func GetBestScoreHistory(user_id int) {
	connectionAccount := "root:112234@/quiz_game_base"
	dataBase := openBase(connectionAccount)
	scoreHistoryRow := dataBase.QueryRow("SELECT score_id,date_played FROM scores_history JOIN high_scores ON high_scores.history_id=scores_history.score_id WHERE high_scores.user_id = (?)", user_id)
	var scoreId int
	var datePlayed mysql.NullTime
	err := scoreHistoryRow.Scan(&scoreId, &datePlayed)
	PrintErr(err)
	timY, timM, timD := datePlayed.Time.Date()
	fmt.Println("Best score id :", scoreId)
	fmt.Printf("Best score date : %d-%s-%d\n", timY, timM, timD)

}

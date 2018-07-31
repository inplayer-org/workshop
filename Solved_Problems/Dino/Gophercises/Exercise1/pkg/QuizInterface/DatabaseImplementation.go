package quizInterface

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func printErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func FindHighScore(hiScore <-chan int, end chan<- bool, contin chan<- bool, user_id int) {
	connectionAccount := "root:112234@/quiz_game_base"
	dataBase := openBase(connectionAccount)
	highScoreRow := dataBase.QueryRow("SELECT `high_score` FROM high_scores WHERE `user_id`=(?)", user_id)
	var highScore int
	err := highScoreRow.Scan(&highScore)
	printErr(err)
	for score := range hiScore {
		time.Sleep(time.Second * 2)
		_, err = dataBase.Exec("INSERT INTO scores_history (`score`,`date_played`,`user_id`) VALUES ((?),now(),(?))", score, user_id)
		fmt.Println()
		if highScore < score {
			highScore = score
			fmt.Println("Congratulations !! NEW HIGHEST SCORE ", highScore)
			scoreIdRow := dataBase.QueryRow("SELECT `score_id` FROM scores_history ORDER BY score_id DESC LIMIT 1")
			var score_id int
			err = scoreIdRow.Scan(&score_id)
			printErr(err)
			_, err = dataBase.Exec("UPDATE high_scores SET `high_score`=(?),`history_id`=(?) WHERE `user_id`=(?)", highScore, score_id, user_id)
			printErr(err)
		} else {
			fmt.Println("Your highest score remains", highScore)
		}
		contin <- true
	}
	fmt.Println("Your highest score was", highScore)
	end <- true
}

func openBase(client string) *sql.DB {
	dataBase, err := sql.Open("mysql", client)
	if err != nil {
		fmt.Println("ERROR WITH DATABASE !!!", err)
		os.Exit(3)
	}
	return dataBase

}

func LoginSystem() int {
	connectionAccount := "root:112234@/quiz_game_base"
	userInputReader := bufio.NewReader(os.Stdin)
	dataBase := openBase(connectionAccount)
	successfulLogin := false
	var user_id int
	for !successfulLogin {
		username := handleUsername()

		usernameRow := dataBase.QueryRow("SELECT `pass` FROM users WHERE `username` = (?)", username)

		var userPassword string
		err := usernameRow.Scan(&userPassword)
		if err != nil {
			fmt.Println("User doesn't exists in the database, do you want to create a new User ? (y/n)..")
			if !Repeat() {
				continue
			} else {
				user_id = createAccount(dataBase)
				fmt.Println("\nLoging in...\n")
				time.Sleep(time.Second * 2)
				fmt.Println("Login successful !")
				time.Sleep(time.Second * 1)
				break
			}
		}
		for {
			fmt.Print("Enter your password : ")
			passEntry, _, err := userInputReader.ReadLine()
			printErr(err)
			fmt.Println("\nLoging in...\n")
			time.Sleep(time.Second * 2)
			if string(passEntry) == userPassword {
				fmt.Println("Login successful !")
				getIdRow := dataBase.QueryRow("SELECT `user_id` FROM users WHERE `username`=(?)", username)
				err = getIdRow.Scan(&user_id)
				printErr(err)
				time.Sleep(time.Second * 1)
				successfulLogin = !successfulLogin
				break
			} else {
				fmt.Println("Invalid password")
				fmt.Println("Do you want to try again ? (y/n)..")
				if !Repeat() {
					break
				} else {
					continue
				}
			}
		}
	}
	return user_id
}

func validateInput(input string, maxCharacters int) bool {
	input = strings.Trim(input, " ")
	if len(input) < maxCharacters && len(input) > 3 {
		return false
	}
	return true
}

func createAccount(dataBase *sql.DB) int {
	accountCredentials := []string{} // [0] = username, [1] = pass, [2] = full_name
	fmt.Println("\nACCOUNT CREATION MENU : ")
	accountCredentials = append(accountCredentials, handleUsername())
	accountCredentials = append(accountCredentials, handlePassword())
	accountCredentials = append(accountCredentials, handleFullName())
	_, err := dataBase.Exec("INSERT INTO users (`username`,`pass`,`full_name`) VALUES ((?),(?),(?))",
		accountCredentials[0], accountCredentials[1], accountCredentials[2])
	if err != nil {
		panic("SOMETHING WENT WRONG WHEN CREATING THE ACCOUNT")
	}
	userIdRow := dataBase.QueryRow("SELECT `user_id` FROM users WHERE `username`=(?)", accountCredentials[0])
	var userId int
	err = userIdRow.Scan(&userId)
	printErr(err)
	_, err = dataBase.Exec("INSERT INTO high_scores (`user_id`) VALUE ((?))", userId)
	if err != nil {
		panic("SOMETHING WENT WRONG WHEN CREATING THE ACCOUNT")
	}
	fmt.Printf("\nAccount Created !\n\n Your info :\n\n User_Id : %d\n Username : %s\n Password : %s\n Full Name: : %s\n High Score: 0\n", userId, accountCredentials[0], accountCredentials[1], accountCredentials[2])

	return userId
}

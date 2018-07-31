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

func FindHighScore(hiScore <-chan int, end chan<- bool, contin chan<- bool) {
	highScore := -1
	for score := range hiScore {
		time.Sleep(time.Second * 2)
		fmt.Println()
		if highScore < score {
			highScore = score
			fmt.Println("Congratulations !! NEW HIGHEST SCORE ", highScore)
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

func LoginSystem() {
	connectionAccount := "root:112234@/quiz_game_base"
	userInputReader := bufio.NewReader(os.Stdin)
	dataBase := openBase(connectionAccount)
	successfulLogin := false

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
				username = createAccount(dataBase)
				break
			}
		}

		if len(userPassword) == 0 {
			fmt.Println("user", username, "doesn't exist in our database")
			fmt.Println(userPassword)
			os.Exit(1)
		} else {
			passEntry, _, err := userInputReader.ReadLine()
			printErr(err)
			if string(passEntry) == userPassword {
				fmt.Println("Login successful")
			} else {
				fmt.Println("Invalid password")
			}
		}

		os.Exit(4)
	}
}

func validateInput(input string, maxCharacters int) bool {
	input = strings.Trim(input, " ")
	if len(input) < maxCharacters && len(input) > 3 {
		return false
	}
	return true
}

func handleUsername() string {
	userInputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username : ")
	input, _, err := userInputReader.ReadLine()
	printErr(err)
	username := string(input)

	for validateInput(username, 16) {
		fmt.Println("Invalid input, username has to be less than 16 characters and more than 3 characters")
		fmt.Print("Enter your username : ")
		input, _, err := userInputReader.ReadLine()
		printErr(err)
		username = string(input)
	}
	return username
}

func handlePassword() string {
	userInputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your password : ")
	input, _, err := userInputReader.ReadLine()
	printErr(err)
	pass := string(input)

	for validateInput(pass, 16) {
		fmt.Println("Invalid input, password has to be less than 16 characters and more than 3 characters")
		fmt.Print("Enter your password : ")
		input, _, err := userInputReader.ReadLine()
		printErr(err)
		pass = string(input)
	}
	return pass
}

func handleFullName() string {
	userInputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Full Name : ")
	input, _, err := userInputReader.ReadLine()
	printErr(err)
	full_name := string(input)

	for validateInput(full_name, 32) {
		fmt.Println("Invalid input, Full Name has to be less than 32 characters and more than 3 characters")
		fmt.Print("Enter your Full Name : ")
		input, _, err := userInputReader.ReadLine()
		printErr(err)
		full_name = string(input)
	}
	return full_name
}

func createAccount(dataBase *sql.DB) string {
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
	fmt.Printf("\nAccount Created !\n\n Your info :\n\n Username : %s\n Password : %s\n Full Name: : %s\n", accountCredentials[0], accountCredentials[1], accountCredentials[2])

	return accountCredentials[0]
}

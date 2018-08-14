package quizinput

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func errorHandler(err error) {
	if err != nil {
		fmt.Println("ERROR :", err)
	}
}

//ChooseBetweenTwo - Forces the user to chose between 2 options (returns true if option1 is entered, returns false if option2 is entered)
func ChooseBetweenTwo(option1 string, option2 string) bool {

	retake := UserInputReader()

	//Making all the inputs to upper case in order to evade case sensitive testing
	retake = strings.ToUpper(retake)
	option1 = strings.ToUpper(option1)
	option2 = strings.ToUpper(option2)

	for retake != option1 && retake != option2 {
		fmt.Printf("Invalid value, please enter (%s/%s)..\n ", option1, option2)
		fmt.Scanln(&retake)
		retake = strings.ToUpper(retake)
	}
	if retake == option2 {
		return false
	}
	return true
}

//DetermineDataBase - Chose Dino or Elena for accesing database
func DetermineDataBase() string {
	password := ChooseBetweenTwo("Dino", "Elena")
	if password { //Dino
		return "112234"
	} //ChooseBetweenTwo forces the user to chose either Dino or Elena so default is Elena
	return "1111"
}

//CheckEqualStrings - check if 2 strings are equal regardless of Case
func CheckEqualStrings(entry1 string, entry2 string) bool {
	entry1 = strings.ToUpper(entry1)
	entry2 = strings.ToUpper(entry2)
	if entry1 == entry2 {
		return true
	}
	return false
}

//EntryForWebOrTerminal - Choose whether to play the game in terminal(1) or web(0)
func EntryForWebOrTerminal() string {
	fmt.Println("Enter 1 for playing in terminal or Enter 0 for web")
	userInput := UserInputReader()
	return userInput
}

func UserInputReader() string {
	reader := bufio.NewReader(os.Stdin)
	userInput, _, err := reader.ReadLine()
	errorHandler(err)
	return string(userInput)
}

//EnterPlayerName - Reading user input through terminal to use for inserting the score into the database
func EnterPlayerName() string {
	fmt.Println("Enter your name:")
	playerName := UserInputReader()
	return playerName
}

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

//ControlYesNo - Checks if the entered input is Y or N and isn't case sensitive
func ControlYesNo() bool {
	var retake string
	fmt.Scanln(&retake)
	retake = strings.ToUpper(retake)
	for retake != "Y" && retake != "N" {
		fmt.Print("Invalid value, please enter (y/n).. ")
		fmt.Scanln(&retake)
		retake = strings.ToUpper(retake)
	}
	if retake == "N" {
		return false
	}
	return true
}

//
func DetermineDataBase() string {
	reader := bufio.NewReader(os.Stdin)
	for {
		passwordTry, _, err := reader.ReadLine()
		errorHandler(err)
		if CheckEqualStrings(string(passwordTry), "Dino") {
			return "112234"
		} else if CheckEqualStrings(string(passwordTry), "Elena") {
			return "1111"
		}
		fmt.Println("Please enter Dino or Elena to access your database")
	}
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

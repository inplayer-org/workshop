package quizinput

import (
	"fmt"
	"strings"
)

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

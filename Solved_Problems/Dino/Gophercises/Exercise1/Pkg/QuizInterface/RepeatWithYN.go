package quizInterface

import (
	"fmt"
	"strings"
)

func RepeatQuiz() bool {
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

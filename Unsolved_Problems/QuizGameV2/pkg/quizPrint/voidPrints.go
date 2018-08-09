package quizprint

import "fmt"

// PrintQuestion - print the current question asked
func PrintQuestion(currentQuestion string) {
	fmt.Printf("%s = ", currentQuestion)
}

//PrintCurrentSettings - prints the current settings set by the flash and waits user to press enter to begin the execution
func PrintCurrentSettings(fileName string, quizTimerDuration int) {
	fmt.Println()
	fmt.Println("To get info about setting flag values initialize the program with -f from command line")
	fmt.Println("Questions base:", fileName)
	fmt.Println("Timer set to:", quizTimerDuration, "seconds")
	fmt.Println()
}

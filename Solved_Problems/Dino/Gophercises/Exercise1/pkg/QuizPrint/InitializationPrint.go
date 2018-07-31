package quizPrinting

import "fmt"

func InitializationPrint(fileName string, randomQuestions bool, quizTimer int) {
	fmt.Println()
	fmt.Println("To get info about setting flag values initialize the program with -f from command line")
	fmt.Println("Questions base:", fileName)
	fmt.Println("Random questions:", randomQuestions)
	fmt.Println("Timer set to:", quizTimer)
}

package quizPrinting

import "fmt"

func Results(correctAnswers, incorrectAnswers, numberOfQuestionsInBase int) {
	fmt.Printf("\n\nYour Results:\n")
	fmt.Printf("%31s %4d\n", "Correct answers :", correctAnswers)
	fmt.Printf("%31s %4d\n", "Incorrect answers :", incorrectAnswers)
	fmt.Printf("%31s %4d\n", "Questions present in the base :", numberOfQuestionsInBase)
}

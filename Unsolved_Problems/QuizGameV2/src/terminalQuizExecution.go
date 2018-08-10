package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	qInput "repo.inplayer.com/workshop/Unsolved_Problems/QuizGameV2/pkg/quizInput"
	qMySql "repo.inplayer.com/workshop/Unsolved_Problems/QuizGameV2/pkg/quizMySQL"
	qPrint "repo.inplayer.com/workshop/Unsolved_Problems/QuizGameV2/pkg/quizPrint"
)

func createTimer(timeDuration int) *time.Timer {
	newTimer := time.NewTimer(time.Second * time.Duration(timeDuration))
	return newTimer
}

func registerUserAnswer() string {

	answerAttempt := qInput.UserInputReader()
	answerAttempt = strings.Replace(answerAttempt, " ", "", -1) // Remove all blank spaces and convert []bytes to string

	return answerAttempt
}

func quizLifeCycle(questionBase []questionStructure, endQuizByQuestionsChannel chan<- bool, hasTimerFinished <-chan bool, waitingTheScore chan<- int) {

	currentScore := 0

	for questionNumber := range questionBase {

		qPrint.PrintQuestion(questionBase[questionNumber].Question)

		attemptedAnswer := registerUserAnswer()
		correctAnswer := questionBase[questionNumber].Answer

		select {
		case <-hasTimerFinished:
			waitingTheScore <- currentScore
			return

		default:
			if checkAnswerCorrectness(attemptedAnswer, correctAnswer) {
				currentScore++
			}
		}
	}
	endQuizByQuestionsChannel <- true
	waitingTheScore <- currentScore
}

func controlQuizEnding(endQuizByQuestionsChannel <-chan bool, endQuizByTimeChannel *time.Timer, hasTimerFinished chan<- bool) {

	select {
	case <-endQuizByQuestionsChannel:
		fmt.Printf("\n\tYou gave answer to all possible questions\n")
		endQuizByTimeChannel.Stop()

	case <-endQuizByTimeChannel.C:
		fmt.Printf("\n\tYour time ran out, press enter to continue\n")
		hasTimerFinished <- true
	}
}

func quizProgramController(questionBase []questionStructure, quizTimerDuration int) int {

	//Controling whether the quiz is ended by running out of time or by answering all possible questions in the database
	endQuizByTimeChannel := createTimer(quizTimerDuration)
	endQuizByQuestionsChannel := make(chan bool)

	//Sending signal that the time has ran out so it can finish the questioning function
	hasTimerFinished := make(chan bool)

	//Channel to make sure the program gets back the current score before continuing and making sure quizLifeCycle is already ended before starting new game
	waitingTheScore := make(chan int)

	//Starting a game
	go quizLifeCycle(questionBase, endQuizByQuestionsChannel, hasTimerFinished, waitingTheScore)

	//Control how the quiz will end and return the score
	controlQuizEnding(endQuizByQuestionsChannel, endQuizByTimeChannel, hasTimerFinished)

	newestPlayScore := <-waitingTheScore

	fmt.Println("Current score =", newestPlayScore)
	//Closing all channels involved in the quiz execution (timer is closed with .Stop() or automatically if it has expired)
	close(endQuizByQuestionsChannel)
	close(hasTimerFinished)
	close(waitingTheScore)
	return newestPlayScore

}
func executeGameInTerminal(currentQuestionsData []questionStructure, quizTimerDuration int, db *sql.DB) {

	//Make sure the user is ready to start a game
	fmt.Print("Press Enter when you are ready to start ")
	fmt.Scanln()

	//Receving the user name and score from current play and inserting it into the database
	score := quizProgramController(currentQuestionsData, quizTimerDuration)
	name := qInput.EnterPlayerName()
	qMySql.InsertIntoHighScores(db, name, score)

	//Print top 10 PLayers
	qMySql.PrintTop10(db)
}

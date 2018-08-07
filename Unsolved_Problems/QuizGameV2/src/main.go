package main

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type questionStructure struct { //Structure for parsing questions from CSV file
	question string
	answer   string
}

type Player struct {
	Name  string
	Score int
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println("ERROR :", err)
	}
}
func createQuestionStructure(fileName string) []questionStructure { //Creates slice of QuestionStructure from a CSV file with (filename)
	file, err := os.Open(fileName)
	errorHandler(err)
	csvReader := csv.NewReader(bufio.NewReader(file))
	sliceOfQuestionStructures := []questionStructure{}
	for {
		readRow, err := csvReader.Read() //readRow is a row from CSV file where readRow[0] is the unfiltered question and readRow[1] is the answer
		if err == io.EOF {
			break //End the endless loop when csvReader tries to read EOF
		} else {
			errorHandler(err)
		}
		sliceOfQuestionStructures = append(sliceOfQuestionStructures, createQuestionStructureEntry(readRow))
	}
	return sliceOfQuestionStructures
}

func createQuestionStructureEntry(readRow []string) questionStructure { //Creates QuestionStructure entry that is appended to the slice of all questions present
	return questionStructure{question: filterQuestion(readRow[0]), answer: readRow[1]}
}
func filterQuestion(unfilteredQuestion string) string { //Remove all unnecessary characters from the question
	regexpFilter := regexp.MustCompile("([0-9]+)|(\\+|\\*|-|/|\\^)|([0-9]+)")
	return strings.Join(regexpFilter.FindAllString(unfilteredQuestion, -1), "")
}

func createTimer(timeDuration int) *time.Timer {
	newTimer := time.NewTimer(time.Second * time.Duration(timeDuration))
	return newTimer
}

func printQuestion(currentQuestion questionStructure) {
	fmt.Printf("%s = ", currentQuestion.question)
}

func registerUserAnswer() string {
	reader := bufio.NewReader(os.Stdin)

	answerEntry, _, err := reader.ReadLine()
	errorHandler(err)

	trimmedAnswer := strings.Replace(string(answerEntry), " ", "", -1) // Remove all blank spaces and convert []bytes to string

	return trimmedAnswer
}

func checkAnswerCorrectness(attemptedAnswer string, correctAnswer string) bool {
	if attemptedAnswer == correctAnswer {
		return true
	}
	return false
}

func quizLifeCycle(questionBase []questionStructure, endQuizByQuestionsChannel chan<- bool, hasTimerFinished <-chan bool, waitingTheScore chan<- int) {
	currentScore := 0

	for questionNumber := range questionBase {

		printQuestion(questionBase[questionNumber])

		attemptedAnswer := registerUserAnswer()
		correctAnswer := questionBase[questionNumber].answer

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
func dbConnect() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:1111@tcp(127.0.0.1:3306)/QuizGame")
	if err != nil {
		panic(err.Error())
	}
	return db
}

//Insert name and score of the PLayer in table HighScores
func insertIntoHighScores(db *sql.DB, name string, score int) {
	_, err := db.Exec("INSERT INTO HighScores(score,name) VALUES (?,?)", score, name)
	errorHandler(err)
}

func EntryName() string {
	fmt.Println("Entry name:")
	reader := bufio.NewReader(os.Stdin)
	name, _, err := reader.ReadLine()
	errorHandler(err)
	return string(name)
}

func main() {
	db := dbConnect()
	defer db.Close()
	fileName := "../csv/"
	flagFile := flag.String("Questions", "Problems1", "Problems1,Problems2")
	quizTimerDuration := flag.Int("Timer", 30, "Set the duration of the timer")
	flag.Parse()
	fileName += *flagFile + ".csv"
	name := EntryName()
	score := quizProgramController(createQuestionStructure(fileName), *quizTimerDuration)
	insertIntoHighScores(db, name, score)

}

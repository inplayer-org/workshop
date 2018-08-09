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

	_ "github.com/go-sql-driver/mysql"
	qInput "repo.inplayer.com/workshop/Unsolved_Problems/QuizGameV2/pkg/quizInput"
	qMySql "repo.inplayer.com/workshop/Unsolved_Problems/QuizGameV2/pkg/quizMySQL"
	qPrint "repo.inplayer.com/workshop/Unsolved_Problems/QuizGameV2/pkg/quizPrint"
)

type questionStructure struct { //Structure for parsing questions from CSV file
	Question string
	Answer   string
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
	return questionStructure{Question: filterQuestion(readRow[0]), Answer: readRow[1]}
}
func filterQuestion(unfilteredQuestion string) string { //Remove all unnecessary characters from the question
	regexpFilter := regexp.MustCompile("([0-9]+)|(\\+|\\*|-|/|\\^)|([0-9]+)")
	return strings.Join(regexpFilter.FindAllString(unfilteredQuestion, -1), "")
}

func checkAnswerCorrectness(attemptedAnswer string, correctAnswer string) bool {
	if attemptedAnswer == correctAnswer {
		return true
	}
	return false
}

func executeGame(chosePlatform bool, currentQuestionsData []questionStructure, quizTimerDuration int, db *sql.DB, fileName string) {
	if chosePlatform {
		executeGameInTerminal(currentQuestionsData, quizTimerDuration, db)
	} else {
		executeGameOnWeb(fileName)
	}
}

func choseGamePlatform() bool {
	fmt.Println("Type either \"Terminal\" or \"Web\" to chose the platform you are going to play on")
	choosePlatform := qInput.ChooseBetweenTwo("Terminal", "Web") //Terminal returns true, Web returns false
	fmt.Println()
	fmt.Print("Press Enter when you are ready to start ")
	fmt.Scanln()
	return choosePlatform
}

func main() {

	//Opening connection to the database
	db := qMySql.DataBaseConnect()
	defer db.Close()

	//Initializing flags
	flagFile := flag.String("Questions", "Problems1", "Problems1,Problems2")
	quizTimerDuration := flag.Int("Timer", 30, "Set the duration of the timer")
	flag.Parse()

	//Parsing the csv file name
	fileName := "../csv/"
	fileName += *flagFile + ".csv"

	//Priting current quiz settings (Printing the flags) and waits user to press Enter to continue
	qPrint.PrintCurrentSettings(*flagFile, *quizTimerDuration)

	//Create the question structure for the quiz
	currentQuestionsData := createQuestionStructure(fileName)

	//Choose to play in terminal or on web
	gamePlatform := choseGamePlatform()
	executeGame(gamePlatform, currentQuestionsData, *quizTimerDuration, db, fileName)
}

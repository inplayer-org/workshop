package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

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

var tmpl = template.Must(template.ParseGlob("../tmpl/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "Index", nil)
}

func Timee() time.Time {
	return time.Now()
}
func Start(fileName string) func(w http.ResponseWriter, r *http.Request) {
	sliceOfQuestionStructures := []questionStructure{}
	sliceOfQuestionStructures = createQuestionStructure(fileName)
	j := 0
	correct := 0
	start := Timee()
	quizDuration := start.Add(time.Second * 30)
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		if r.Method == "GET" {
			tmpl.ExecuteTemplate(w, "Start", sliceOfQuestionStructures[0])
		} else {
			j = j + 1
			ans := sliceOfQuestionStructures[j-1].Answer
			if r.FormValue("Answer") == ans {
				correct++
			}

			if (r.FormValue("Next") == "Next") && (j < len(sliceOfQuestionStructures)) && (t.Sub(start) < quizDuration.Sub(start)) {
				tmpl.ExecuteTemplate(w, "Start", sliceOfQuestionStructures[j])

			} else if t.Sub(start) > quizDuration.Sub(start) {
				tmpl.ExecuteTemplate(w, "TimeRanOut", nil)

			} else {

				fmt.Fprintf(w, "resi")
			}

		}
		log.Println("CORRECT:", correct)
	}

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

	//Choose to play in terminal or on web
	entry1or0 := qInput.EntryForWebOrTerminal()
	for {
		if entry1or0 == "1" {
			//Priting current quiz settings (Printing the flags) and waits user to press Enter to continue
			qPrint.PrintCurrentSettings(*flagFile, *quizTimerDuration)

			//Receving the user name and score from current play and inserting it into the database
			score := QuizProgramController(createQuestionStructure(fileName), *quizTimerDuration)
			name := qInput.EnterPlayerName()
			qMySql.InsertIntoHighScores(db, name, score)

			//Print top 10 PLayers
			qMySql.PrintTop10(db)
		} else if entry1or0 == "0" {
			log.Println("SERVER started on localhost:3010")
			http.HandleFunc("/", Index)
			http.HandleFunc("/start", Start(fileName))
			http.ListenAndServe(":3010", nil)
		} else {
			fmt.Println("You have to enter 0 or 1")
			entry1or0 = qInput.EntryForWebOrTerminal()
		}

	}
}

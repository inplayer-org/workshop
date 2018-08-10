package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

var start time.Time

type questionStructure struct { //Structure for parsing questions from CSV file
	Question string
	Answer   string
}

type webStructure struct {
	question string
	timeLeft float64
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

var control bool

type questions struct {
	prashanja []string
	timer     *time.Timer
}

func (p *questions) questionsHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	p.prashanja = p.prashanja[1:]
	fmt.Println("1")
	pominatoVreme := time.Since(start).Seconds()
	fmt.Println(pominatoVreme)
	t, _ := template.ParseFiles("FrontEnd/html/questions.html")
	send := webStructure{question: p.prashanja[0], timeLeft: 15.00 - pominatoVreme}
	t.Execute(w, send.timeLeft)
	fmt.Println(r.Form)
}

func handler(w http.ResponseWriter, r *http.Request) {
	control = true
	t, _ := template.ParseFiles("FrontEnd/html/index.html")
	t.Execute(w, nil)
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	control = true
	t, _ := template.ParseFiles("FrontEnd/html/index.html")
	t.Execute(w, nil)
}

func actionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Vleguva vo handlerot")

}

func serveCss(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "FrontEnd/css/default.css")

}

func endHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("FrontEnd/html/end.html")
	t.Execute(w, nil)
}

// func questionsHandler(fileName string) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.Response, r *http.Request) {
// 		r.ParseForm()
// 		p.prashanja = p.prashanja[1:]
// 		fmt.Println("1")
// 		pominatoVreme := time.Since(start).Seconds()
// 		fmt.Println(pominatoVreme)
// 		t, _ := template.ParseFiles("FrontEnd/html/questions.html")
// 		send := webStructure{question: p.prashanja[0], timeLeft: 15.00 - pominatoVreme}
// 		t.Execute(w, send.timeLeft)
// 		fmt.Println(r.Form)
// 	}
// }
func main() {
	start = time.Now()
	var quest questions
	file := "../csv/" + "Problems1.csv"
	questions := createQuestionStructure(file)
	fmt.Println(questions)
	control = true
	http.HandleFunc("/", handler)
	http.HandleFunc("/action/", actionHandler)
	http.HandleFunc("/css/", serveCss)
	http.HandleFunc("/questions/", quest.questionsHandler)
	http.HandleFunc("/end/", endHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

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
func executeGameOnWeb(fileName string) {
	log.Println("SERVER started on localhost:3010")
	http.HandleFunc("/", Index)
	http.HandleFunc("/start", Start(fileName))
	http.ListenAndServe(":3010", nil)
}

package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	qMySql "repo.inplayer.com/workshop/Unsolved_Problems/QuizGameV2/pkg/quizMySQL"
)

var tmpl = template.Must(template.ParseGlob("../tmpl/*"))
var htmls = template.Must(template.ParseGlob("../html/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	htmls.ExecuteTemplate(w, "menu.html", nil)
}

func Timee() time.Time {
	return time.Now()
}

func highScoreHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		test := qMySql.GetTop10(db)

		htmls.ExecuteTemplate(w, "rankings.html", test)
	}
}

func showUsersTop10Plays(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Form)
		test := qMySql.GetTop10PlaysOfUser(db, r.FormValue("userSearch"))
		htmls.ExecuteTemplate(w, "rankings.html", test)
	}
}

func findUsersTop10Plays(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		htmls.ExecuteTemplate(w, "findPlayer.html", nil)
	}
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

func serveCss(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../css/default.css")

}

func executeGameOnWeb(fileName string, db *sql.DB) {
	log.Println("SERVER started on localhost:3010")
	http.HandleFunc("/", Index)
	http.HandleFunc("/start", Start(fileName))
	http.HandleFunc("/css/", serveCss)
	http.HandleFunc("/rankings", highScoreHandler(db))
	http.HandleFunc("/findPlayer", findUsersTop10Plays(db))
	http.HandleFunc("/showPlayer", showUsersTop10Plays(db))
	http.ListenAndServe(":3010", nil)
}

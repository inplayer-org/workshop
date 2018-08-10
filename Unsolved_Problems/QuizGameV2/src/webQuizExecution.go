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

var htmls = template.Must(template.ParseGlob("../html/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	htmls.ExecuteTemplate(w, "menu.html", nil)
}

type questions struct {
	Prasanja []string
	Timer    *time.Timer
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

type WebStructure struct {
	Question string
	TimeLeft float64
}

func Questions(sliceOfQuestionStructures []questionStructure, timerDuration int) func(w http.ResponseWriter, r *http.Request) {
	j := 0
	i := len(sliceOfQuestionStructures)
	start := time.Now()
	return func(w http.ResponseWriter, r *http.Request) {
		if i != 0 {
			r.ParseForm()
			fmt.Println(len(sliceOfQuestionStructures), j)
			pominatoVreme := time.Since(start).Seconds()
			fmt.Println(pominatoVreme)
			fmt.Println("len", i)
			t, _ := template.ParseFiles("FrontEnd/html/questions.html")
			send := WebStructure{Question: sliceOfQuestionStructures[j].Question, TimeLeft: float64(timerDuration) - pominatoVreme}
			j = j + 1
			i = i - 1
			t.Execute(w, send)
			fmt.Println(r.Form["answer"])
		} else {
			http.Redirect(w, r, "/action/", 301)
		}
	}
}

func serveCss(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../css/default.css")

}

func executeGameOnWeb(currentQuestionsData []questionStructure, quizTimerDuration int, db *sql.DB) {
	log.Println("SERVER started on localhost:3010")
	http.HandleFunc("/", Index)
	http.HandleFunc("/css/", serveCss)
	http.HandleFunc("/rankings", highScoreHandler(db))
	http.HandleFunc("/findPlayer", findUsersTop10Plays(db))
	http.HandleFunc("/showPlayer", showUsersTop10Plays(db))
	http.HandleFunc("/questions/", Questions(currentQuestionsData, quizTimerDuration))
	http.ListenAndServe(":3010", nil)
}

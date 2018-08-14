package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"

	qMySql "repo.inplayer.com/workshop/Unsolved_Problems/QuizGameV2/pkg/quizMySQL"
)

var htmls = template.Must(template.ParseGlob("../html/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	htmls.ExecuteTemplate(w, "menu.html", nil)
}

func submitScoreHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println("submit score Handler Form =", r.Form)
		score, err := strconv.Atoi(r.FormValue("finalScore"))
		errorHandler(err)
		qMySql.InsertIntoHighScores(db, r.FormValue("name"), score)
		http.Redirect(w, r, "/rankings", 301)
	}
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
	Question     string
	TimeLeft     float64
	CurrentScore int
}

func Questions(sliceOfQuestionStructures []questionStructure, timerDuration int) func(w http.ResponseWriter, r *http.Request) {
	j := 0
	i := len(sliceOfQuestionStructures)
	startingLen := len(sliceOfQuestionStructures)
	score := 0
	start := time.Now()
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if i != startingLen {
			fmt.Println("r.Form.Val(answer) = ", r.FormValue("answer"), "sliceOfQuestionsStructures[", j-1, "].Answer =", sliceOfQuestionStructures[j-1].Answer)
			if r.FormValue("answer") == sliceOfQuestionStructures[j-1].Answer {
				score++
			}
		}
		if i != 0 {
			fmt.Println(len(sliceOfQuestionStructures), j)
			pominatoVreme := time.Since(start).Seconds()
			//fmt.Println(pominatoVreme)
			fmt.Println("len = ", i)
			send := WebStructure{Question: sliceOfQuestionStructures[j].Question, TimeLeft: float64(timerDuration) - pominatoVreme, CurrentScore: score}
			j = j + 1
			i--
			//log.Println("timer =", send.TimeLeft)
			htmls.ExecuteTemplate(w, "questions.html", send)
			fmt.Println(r.Form["answer"])
			fmt.Println(send)
		} else {
			send := WebStructure{Question: "ended", TimeLeft: 0.00, CurrentScore: score}
			htmls.ExecuteTemplate(w, "questions.html", send)
		}
	}
}

func serveCSS(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../css/default.css")

}

func executeGameOnWeb(currentQuestionsData []questionStructure, quizTimerDuration int, db *sql.DB) {
	log.Println("SERVER started on localhost:3010")
	log.Println(currentQuestionsData)
	http.HandleFunc("/", Index)
	http.HandleFunc("/css/", serveCSS)
	http.HandleFunc("/rankings/", highScoreHandler(db))
	http.HandleFunc("/findPlayer/", findUsersTop10Plays(db))
	http.HandleFunc("/showPlayer/", showUsersTop10Plays(db))
	http.HandleFunc("/question/", Questions(currentQuestionsData, quizTimerDuration))
	http.HandleFunc("/submitScore/", submitScoreHandler(db))
	http.ListenAndServe(":3000", nil)
}

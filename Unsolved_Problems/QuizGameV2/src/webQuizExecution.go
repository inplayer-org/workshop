package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type webStructure struct {
	question string
	timeLeft float64
}

//var tmpl = template.Must(template.ParseGlob("../tmpl/*"))
func serveCss(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "FrontEnd/css/default.css")

}
func Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("FrontEnd/html/index.html")
	t.Execute(w, nil)
}

type questions struct {
	Prasanja []string
	Timer    *time.Timer
}
type WebStructure struct {
	Question string
	TimeLeft float64
}

func ActionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Vleguva vo handlerot")

}

func Questions(fileName string) func(w http.ResponseWriter, r *http.Request) {
	sliceOfQuestionStructures := []questionStructure{}
	sliceOfQuestionStructures = createQuestionStructure(fileName)
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
			send := WebStructure{Question: sliceOfQuestionStructures[j].Question, TimeLeft: 30.00 - pominatoVreme}
			j = j + 1
			i = i - 1
			t.Execute(w, send)
			fmt.Println(r.Form["answer"])
		} else {
			http.Redirect(w, r, "/action/", 301)
		}
	}
}

func executeGameOnWeb(fileName string) {
	log.Println("SERVER started on localhost:3010")
	http.HandleFunc("/", Index)
	http.HandleFunc("/action/", ActionHandler)
	http.HandleFunc("/css/", serveCss)
	http.HandleFunc("/questions/", Questions(fileName))
	// http.HandleFunc("/end/", endHandler)
	http.ListenAndServe(":3010", nil)
}

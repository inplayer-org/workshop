package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

var control bool

type questions struct {
	prashanja []string
	timer     *time.Timer
}

func (p *questions) questionsHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	p.prashanja = p.prashanja[1:]
	fmt.Println("1")
	t, _ := template.ParseFiles("FrontEnd/html/questions.html")

	t.Execute(w, p.prashanja[0])
	fmt.Println(r.Form)
}

func handler(w http.ResponseWriter, r *http.Request) {
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

func main() {
	var quest questions
	control = true
	quest.prashanja = append(quest.prashanja, "prvo")
	quest.prashanja = append(quest.prashanja, "vtoro")
	quest.prashanja = append(quest.prashanja, "treto")
	quest.prashanja = append(quest.prashanja, "cetvrto")
	http.HandleFunc("/", handler)
	http.HandleFunc("/action/", actionHandler)
	http.HandleFunc("/css/", serveCss)
	http.HandleFunc("/questions/", quest.questionsHandler)
	http.HandleFunc("/end/", endHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

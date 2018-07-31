package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"
)

func readFile(name string) [][]string {
	file, err := os.Open(name)
	if err != nil {
		log.Fatalf("Can not open %s file", name, err.Error())
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = ';'
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatalf("can't read CSV data", err.Error())
	}
	return rows

}

type problem struct {
	question, answer string
}

func parse(rows [][]string) []problem {
	problems := make([]problem, len(rows))
	for i, f := range rows {
		problems[i] = problem{
			question: f[0],
			answer:   f[1],
		}
	}
	return problems
}

func integers(s string) []string {
	re := regexp.MustCompile("([0-9]+)[\\+-/*]([0-9]+)")
	s1 := re.FindAllString(s, -1)
	return s1
}

func main() {
	rows := readFile("problems.csv")
	problems := parse(rows)
	correct := 0
	c3 := make(chan string)
	//c4 := make(chan string)
	for _, i := range problems {
		q := integers(i.question)
		for _, j := range q {
			fmt.Printf("%s=\n", j)
		}

		go func() {
			reader := bufio.NewReader(os.Stdin)

			a, _, _ := reader.ReadLine()
			words := string(a)

			l := strings.Split(words, " ")
			if len(l) < 2 {
				w := l[0]
				c3 <- w
			} else {

				c3 <- "you entered more then 1 number"
			}

		}()

		select {
		case msg3 := <-c3:
			fmt.Println(msg3)
			//correct += check(msg3, i.answer)
			if msg3 == i.answer {
				correct++
				fmt.Println("CORRECT")
			}

		case <-time.After(4 * time.Second):
			fmt.Println("You have 3sec to enter num \n")

		}
	}
	fmt.Println("correct ans: ", correct)

}

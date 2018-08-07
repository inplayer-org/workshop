package main

import (
	"flag"
	"fmt"
)

type questionStructure struct { //Structure for parsing questions from CSV file
	question string
	answer   string
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println("ERROR :", err)
	}
}

//regexp.Compile("([0-9]+)|(\\+|\\*|-|/|\\^)|([0-9]+)")  Filter for all unnecessary characters in the question

func main() {
	fileName := "../csv/"
	flagFile := flag.String("Questions", "Problems1", "Problems1,Problems2")
	quizTimerDuration := flag.Int("Timer", 30, "Set the duration of the timer")
	flag.Parse()
	fileName += *flagFile + ".csv"
	//startProgram(CSV(*flagFile), quizTimerDuration)
}

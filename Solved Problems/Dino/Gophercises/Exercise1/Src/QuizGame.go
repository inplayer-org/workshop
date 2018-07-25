package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type exercises struct {
	question string
	answer   string
}

func check(e error) {
	if e != nil {
		fmt.Println("ERROR ! >>>>", e)
		os.Exit(2)
	}
}

func result(correctness <-chan bool, done chan<- bool, maxQuestions int) {
	answersData := [3]int{0, 0, maxQuestions} // [0] = Correct, [1] = Incorrect, [2] = Number of all questions possible
	go func() {
		for n := range correctness {
			if n {
				answersData[0]++
			} else {
				answersData[1]++
			}
		}
		fmt.Printf("\n\nYour Results:\n")
		fmt.Printf("%31s %4d\n", "Correct answers :", answersData[0])
		fmt.Printf("%31s %4d\n", "Incorrect answers :", answersData[1])
		fmt.Printf("%31s %4d\n", "Questions present in the base :", answersData[2])
		done <- true
	}()
}

func questions(correctness chan<- bool, questChan chan<- bool, quizData []exercises) {
	reader := bufio.NewReader(os.Stdin)
	for questionNumber := range quizData {
		fmt.Printf("%s = ", quizData[questionNumber].question)
		attemptedAnswer, _, err := reader.ReadLine()
		check(err)
		trimmedAnswer := strings.Replace(string(attemptedAnswer), " ", "", -1)
		if trimmedAnswer == quizData[questionNumber].answer {
			correctness <- true
		} else {
			correctness <- false
		}
	}
	questChan <- true
}

func quizExecution(quizData []exercises) {
	correctness := make(chan bool)
	questChan := make(chan bool)
	timerChan := time.NewTimer(time.Second * 10)
	done := make(chan bool)
	go result(correctness, done, len(quizData))
	go questions(correctness, questChan, quizData)
	select {
	case <-questChan:
		fmt.Printf("\n\tYou gave answer to all possible questions\n")
		close(correctness)
	case <-timerChan.C:
		fmt.Printf("\n\tYour time ran out\n")
		close(correctness)
	}
	<-done
	close(done)

}

func dataReader(procitaj string) []exercises {
	returnData := []exercises{}
	csvFile, err := os.Open(procitaj)
	check(err)
	csvReader := csv.NewReader(bufio.NewReader(csvFile))
	for {
		readRow, err := csvReader.Read()
		if err == io.EOF {
			break
		} else {
			check(err)
		}
		exercisesCreator := exercises{question: readRow[0], answer: readRow[1]}
		returnData = append(returnData, exercisesCreator)
	}
	check(csvFile.Close())
	return returnData
}

func main() {
	dataBase := dataReader("../Csv/Problems1.csv")
	fmt.Print("Press Enter when you are ready to start ")
	fmt.Scanln()
	quizExecution(dataBase)

}

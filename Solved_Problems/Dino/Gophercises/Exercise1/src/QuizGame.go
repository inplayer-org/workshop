package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"

	quizI "repo.inplayer.com/workshop/Solved_Problems/Dino/Gophercises/Exercise1/pkg/QuizInterface"
	qPrint "repo.inplayer.com/workshop/Solved_Problems/Dino/Gophercises/Exercise1/pkg/QuizPrint"
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

func result(correctness <-chan bool, done chan<- int, maxQuestions int) {
	answersData := [3]int{0, 0, maxQuestions} // [0] = Correct, [1] = Incorrect, [2] = Number of all questions possible
	go func() {
		for n := range correctness {
			if n {
				answersData[0]++
			} else {
				answersData[1]++
			}
		}
		qPrint.Results(answersData[0], answersData[1], answersData[2])
		done <- answersData[0]
	}()
}

func questions(correctness chan<- bool, questChan chan<- bool, quizData []exercises, nextCase <-chan bool, random bool) {
	reader := bufio.NewReader(os.Stdin)
	var indexRange []int
	if random {
		indexRange = randomizeQuestions(len(quizData))
	} else {
		for i := range quizData {
			indexRange = append(indexRange, i)

		}
	}
	for _, questionNumber := range indexRange {
		fmt.Printf("%s = ", quizData[questionNumber].question)
		attemptedAnswer, _, err := reader.ReadLine()
		check(err)
		trimmedAnswer := strings.Replace(string(attemptedAnswer), " ", "", -1)
		select {
		case <-nextCase:
			return
		default:
			if trimmedAnswer == quizData[questionNumber].answer {
				correctness <- true
			} else {
				correctness <- false
			}
		}
	}
	questChan <- true
}

func quizExecution(quizData []exercises, hiScore chan<- int, random bool, getTimer int) {
	fmt.Print("\nPress Enter when you are ready to start ")
	fmt.Scanln()
	fmt.Println()
	correctness := make(chan bool)
	questChan := make(chan bool)
	timerChan := time.NewTimer(time.Second * time.Duration(getTimer))
	nextCase := make(chan bool)
	done := make(chan int)
	go result(correctness, done, len(quizData))
	go questions(correctness, questChan, quizData, nextCase, random)
	select {
	case <-questChan:
		fmt.Printf("\n\tYou gave answer to all possible questions\n")
		close(correctness)
	case <-timerChan.C:
		fmt.Printf("\n\tYour time ran out, press enter to continue\n")
		nextCase <- false
		close(correctness)
	}
	correctAnswers := <-done
	hiScore <- correctAnswers
	close(done)

}

func randomizeQuestions(permRange int) []int {
	rand.Seed(time.Now().UnixNano())
	indexes := rand.Perm(permRange)
	return indexes
}

func dataReader(procitaj string) []exercises {
	returnData := []exercises{}
	csvFile, err := os.Open(procitaj)
	check(err)
	csvReader := csv.NewReader(bufio.NewReader(csvFile))
	filter, err := regexp.Compile("([0-9]+)|(\\+|\\*|-|/|\\^)|([0-9]+)") // Filter for all unnecessary characters in the question
	check(err)
	for {
		readRow, err := csvReader.Read()
		if err == io.EOF {
			break
		} else {
			check(err)
		}
		questionParser := strings.Join(filter.FindAllString(readRow[0], -1), "")
		//	fmt.Println("Question parser =", questionParser, "Read row =", readRow[0])
		exercisesCreator := exercises{question: questionParser, answer: readRow[1]}
		returnData = append(returnData, exercisesCreator)
	}
	check(csvFile.Close())
	return returnData
}

func quizLifeCycle(user_id int) {
	fileName := "../csv/"
	flagFile := flag.String("Questions", "Problems1", "Problems1,Problems2")
	random := flag.Bool("Random", false, "true or false")
	getTimer := flag.Int("Timer", 30, "Set the duration of the timer")
	flag.Parse()
	qPrint.InitializationPrint(*flagFile, *random, *getTimer)
	fileName += *flagFile + ".csv"
	dataBase := dataReader(fileName)
	start := true
	hiScore := make(chan int)
	end := make(chan bool)
	contin := make(chan bool)
	go quizI.FindHighScore(hiScore, end, contin, user_id)
	for start {
		quizExecution(dataBase, hiScore, *random, *getTimer)
		<-contin
		time.Sleep(time.Second * 2)
		fmt.Print("\nDo you want to retake the quiz ? (y/n) ")
		start = quizI.Repeat()
	}
	close(hiScore)
	<-end
	fmt.Println("Thanks for playing !")
	close(end)
}

func main() {

	user_id := quizI.LoginSystem()
	quizLifeCycle(user_id)
	qPrint.PrintPublicUser(user_id)
	qPrint.PrintTop10()
	qPrint.ListAllUsers()
	qPrint.GetBestScoreHistory(user_id)
}

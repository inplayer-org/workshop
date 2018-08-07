package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
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
func CreateQuestionStructure(fileName string) []questionStructure {
	file, err := os.Open(fileName)
	errorHandler(err)
	csvReader := csv.NewReader(bufio.NewReader(file))
	sliceOfQuestionStructures := []questionStructure{}
	for {
		readRow, err := csvReader.Read()
		if err == io.EOF {
			break
		} else {
			errorHandler(err)
		}
		fmt.Println("row", Filter(readRow[0]))
		sliceOfQuestionStructures = append(sliceOfQuestionStructures, QuestionsSlice(readRow))
	}
	return sliceOfQuestionStructures
}

func QuestionsSlice(readRow []string) questionStructure {
	return questionStructure{question: Filter(readRow[0]), answer: readRow[1]}
}
func Filter(readRow string) string {
	re := regexp.MustCompile("([0-9]+)|(\\+|\\*|-|/|\\^)|([0-9]+)")
	return strings.Join(re.FindAllString(readRow, -1), "")
}
func main() {
	fileName := "../csv/"
	fmt.Println("start")
	flagFile := flag.String("Questions", "Problems1", "Problems1,Problems2")
	//quizTimerDuration := flag.Int("Timer", 30, "Set the duration of the timer")
	flag.Parse()
	fmt.Println("word:", *flagFile)
	fileName += *flagFile + ".csv"
	//startProgram(CSV(*flagFile), quizTimerDuration)

}

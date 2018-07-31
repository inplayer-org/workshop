package quizInterface

import (
	"fmt"
	"time"
)

func FindHighScore(hiScore <-chan int, end chan<- bool, contin chan<- bool) {
	highScore := -1
	for score := range hiScore {
		time.Sleep(time.Second * 2)
		fmt.Println()
		if highScore < score {
			highScore = score
			fmt.Println("Congratulations !! NEW HIGHEST SCORE ", highScore)
		} else {
			fmt.Println("Your highest score remains", highScore)
		}
		contin <- true
	}
	fmt.Println("Your highest score was", highScore)
	end <- true
}

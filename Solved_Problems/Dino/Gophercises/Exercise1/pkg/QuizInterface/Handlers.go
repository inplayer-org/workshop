package quizInterface

import (
	"bufio"
	"fmt"
	"os"
)

func handleUsername() string {
	userInputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your username : ")
	input, _, err := userInputReader.ReadLine()
	printErr(err)
	username := string(input)

	for validateInput(username, 16) {
		fmt.Println("Invalid input, username has to be less than 16 characters and more than 3 characters")
		fmt.Print("Enter your username : ")
		input, _, err := userInputReader.ReadLine()
		printErr(err)
		username = string(input)
	}
	return username
}

func handlePassword() string {
	userInputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your password : ")
	input, _, err := userInputReader.ReadLine()
	printErr(err)
	pass := string(input)

	for validateInput(pass, 16) {
		fmt.Println("Invalid input, password has to be less than 16 characters and more than 3 characters")
		fmt.Print("Enter your password : ")
		input, _, err := userInputReader.ReadLine()
		printErr(err)
		pass = string(input)
	}
	return pass
}

func handleFullName() string {
	userInputReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your Full Name : ")
	input, _, err := userInputReader.ReadLine()
	printErr(err)
	full_name := string(input)

	for validateInput(full_name, 32) {
		fmt.Println("Invalid input, Full Name has to be less than 32 characters and more than 3 characters")
		fmt.Print("Enter your Full Name : ")
		input, _, err := userInputReader.ReadLine()
		printErr(err)
		full_name = string(input)
	}
	return full_name
}

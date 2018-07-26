package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		citaj := scanner.Text()
		citaj = strings.ToUpper(citaj)
		fmt.Println(citaj)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(os.Stderr, "error", err)
		os.Exit(1)
	}
}

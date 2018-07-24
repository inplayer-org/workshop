package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	write1 := []byte("Test Write N.2 !")
	err := ioutil.WriteFile("Write/File1", write1, 0664)
	check(err)

	f, err := os.Create("Write/File2")
	check(err)
	defer f.Close()

	write2 := []byte{'1', '4', 78, 66, '\n'}
	n2, err := f.Write(write2)
	check(err)
	write3 := []byte("Number 3")
	fmt.Println("Wrote", n2, "bytes")
	n3, err := f.Write(write3)
	check(err)
	fmt.Println("Wrote", n3, "bytes")

	f.Sync()

	bufferedWriter := bufio.NewWriter(f)
	n4, err := bufferedWriter.WriteString("\nbuffered \n")
	fmt.Println("Wrote", n4, "bytes")
	bufferedWriter.Flush()
}

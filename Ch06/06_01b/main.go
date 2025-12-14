package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Files")
	fileName := "./fromString.txt"
	file, err := os.Create(fileName)
	checkError(err)
	defer file.Close()
	length, err := file.WriteString("Hello from Go!")
	checkError(err)
	fmt.Println("Length:", length)
	readFile(fileName)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	file, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println(string(file))
}

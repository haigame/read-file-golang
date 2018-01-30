package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		filename := os.Args[1]
		fmt.Println(filename)
		readFile(filename)
	} else {
		fmt.Println("You need to provide a filename")
		os.Exit(1)
	}
}

func readFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)

	}
	io.Copy(os.Stdout, file)
}

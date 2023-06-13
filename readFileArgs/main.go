package main

import (
	"fmt"
	"io"
	"os"
)

// fmt package => used to print on console
// io package => input/output package used to copy values
// os package => used to catch and retrieve input parameters => go run main.go myFile.txt

func main() {
	// Check if filename is passed as parameter
	if len(os.Args) == 0 {
		fmt.Println("File name is required in input")
		os.Exit(1)
	}

	// get the filename from args
	fn := os.Args[1]
	// use the is Open to return a File type instead of using ReadFile that returns a []byte
	file, err := os.Open(fn)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// since type of File have a Reader we can safely pass the file to io.Copy
	io.Copy(os.Stdout, file)

}

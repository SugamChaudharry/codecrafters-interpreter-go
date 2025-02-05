package main

import (
	"fmt"
	"go/token"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/token"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	// Uncomment this block to pass the first stage

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {
		panic("Scanner not implemented")
	} else {
    for _, char := string(fileContents){
      switch char {
        case token.LEFT_PAREN :
          ftm.Println("LEFT_PAREN ( null")
        case token.RIGHT_PAREN :
          ftm.Println("RIGHT_PAREN ) null")
      }
    }
    ftm.Println("EOF null")

	}
}

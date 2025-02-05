package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/cmd/token"
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

	for _, char := range string(fileContents) {
		switch char {
		case token.LEFT_PAREN:
			fmt.Println("LEFT_PAREN ( null")
		case token.RIGHT_PAREN:
			fmt.Println("RIGHT_PAREN ) null")
		case token.LEFT_BRACE:
			fmt.Println("LEFT_BRACE { null")
		case token.RIGHT_BRACE:
			fmt.Println("RIGHT_BRACE } null")
		case token.COMMA:
			fmt.Println("COMMA , null")
		case token.DOT:
			fmt.Println("DOT . null")
		case token.MINUS:
			fmt.Println("MINUS - null")
		case token.PLUS:
			fmt.Println("PLUS + null")
		case token.STAR:
			fmt.Println("STAR * null")
		case token.SEMICOLON:
			fmt.Println("SEMICOLON ; null")
		}
	}
	fmt.Println("EOF  null")

}

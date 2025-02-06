package token

import (
	"fmt"
	"os"
)

const (
	LEFT_PAREN  rune = '('
	RIGHT_PAREN rune = ')'
	LEFT_BRACE  rune = '{'
	RIGHT_BRACE rune = '}'
	COMMA       rune = ','
	DOT         rune = '.'
	MINUS       rune = '-'
	PLUS        rune = '+'
	STAR        rune = '*'
	SEMICOLON   rune = ';'
)

func tokenError(char rune, line int) {
	// Print error message to stderr
	fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", line, char)

	// Print EOF message before exiting
	fmt.Println("EOF  null")

	// Exit with code 65 after printing the error
	os.Exit(65)
}

func HandleToken(fileContents []byte) {
	line := 1
	var tokens []string

	// Loop over each character in the file contents
	for _, char := range string(fileContents) {
		switch char {
		case '\n':
			line++
		case LEFT_PAREN:
			tokens = append(tokens, "LEFT_PAREN ( null")
		case RIGHT_PAREN:
			tokens = append(tokens, "RIGHT_PAREN ) null")
		case LEFT_BRACE:
			tokens = append(tokens, "LEFT_BRACE { null")
		case RIGHT_BRACE:
			tokens = append(tokens, "RIGHT_BRACE } null")
		case COMMA:
			tokens = append(tokens, "COMMA , null")
		case DOT:
			tokens = append(tokens, "DOT . null")
		case MINUS:
			tokens = append(tokens, "MINUS - null")
		case PLUS:
			tokens = append(tokens, "PLUS + null")
		case STAR:
			tokens = append(tokens, "STAR * null")
		case SEMICOLON:
			tokens = append(tokens, "SEMICOLON ; null")
		default:
			tokenError(char, line) // Error handling now exits
		}
	}

	// Print tokens array one by one
	for _, token := range tokens {
		fmt.Println(token)
	}

	// Print EOF message
	fmt.Println("EOF  null")
}

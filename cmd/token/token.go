package token

import (
	"fmt"
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
	// [line N] Error: Unexpected character: <character>
	fmt.Printf("[line %d] Error: Unexpected character: %c\n", line, char)
}

func HandleToken(fileContents []byte) {
	line := 1
	for _, char := range string(fileContents) {
		switch char {
		case '\n':
			line++
		case LEFT_PAREN:
			fmt.Println("LEFT_PAREN ( null")
		case RIGHT_PAREN:
			fmt.Println("RIGHT_PAREN ) null")
		case LEFT_BRACE:
			fmt.Println("LEFT_BRACE { null")
		case RIGHT_BRACE:
			fmt.Println("RIGHT_BRACE } null")
		case COMMA:
			fmt.Println("COMMA , null")
		case DOT:
			fmt.Println("DOT . null")
		case MINUS:
			fmt.Println("MINUS - null")
		case PLUS:
			fmt.Println("PLUS + null")
		case STAR:
			fmt.Println("STAR * null")
		case SEMICOLON:
			fmt.Println("SEMICOLON ; null")
		default:
			tokenError(char, line)
		}
	}
	fmt.Println("EOF  null")
}

package lexer

import (
	"github.com/chayandatta/Hermes/token"
)

type Lexer struct {
	input            string
	position         int  // current input position
	readingPosition  int  // next reading input position, after current position
	currentCharacter byte // current character under examination
}

func New(input string) *Lexer {
	lex := &Lexer{input: input} // Lexer object creation
	lex.readCharacter()
	return lex
}

func (lex *Lexer) readCharacter() {
	if lex.readingPosition >= len(lex.input) {
		lex.currentCharacter = 0
	} else {
		lex.currentCharacter = lex.input[lex.readingPosition]
	}
	lex.position = lex.readingPosition
	lex.readingPosition += 1
}

func (lex *Lexer) NextToken() token.Token {
	var tok token.Token

	switch lex.currentCharacter {
	case '=':
		tok = newToken(token.ASSIGNMENT, lex.currentCharacter)
	case ';':
		tok = newToken(token.SEMICOLON, lex.currentCharacter)
	case ':':
		tok = newToken(token.COLON, lex.currentCharacter)
	case '(':
		tok = newToken(token.LEFT_PARENTHESIS, lex.currentCharacter)
	case ')':
		tok = newToken(token.RIGHT_PARENTHESIS, lex.currentCharacter)
	case ',':
		tok = newToken(token.COMMA, lex.currentCharacter)
	case '+':
		tok = newToken(token.PLUS, lex.currentCharacter)
	case '-':
		tok = newToken(token.MINUS, lex.currentCharacter)
	case '!':
		tok = newToken(token.BANG, lex.currentCharacter)
	case '*':
		tok = newToken(token.ASTERISK, lex.currentCharacter)
	case '/':
		tok = newToken(token.SLASH, lex.currentCharacter)
	case '{':
		tok = newToken(token.LEFT_BRACES, lex.currentCharacter)
	case '}':
		tok = newToken(token.RIGHT_BRACES, lex.currentCharacter)
	case '[':
		tok = newToken(token.LEFT_BRACKET, lex.currentCharacter)
	case ']':
		tok = newToken(token.RIGHT_BRACKET, lex.currentCharacter)
	case '<':
		tok = newToken(token.LESS_THAN, lex.currentCharacter)
	case '>':
		tok = newToken(token.GREATER_THAN, lex.currentCharacter)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	lex.readCharacter()
	return tok
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

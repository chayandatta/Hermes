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

	lex.skipWhiteSpace()

	switch lex.currentCharacter {
	case '=':
		if lex.peekChar() == '=' {
			ch := lex.currentCharacter
			lex.readCharacter()
			tok = token.Token{Type: token.EQUAL, Literal: string(ch) + string(lex.currentCharacter)}
		} else {
			tok = newToken(token.ASSIGNMENT, lex.currentCharacter)
		}
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
	default:
		if isLetter(lex.currentCharacter) {
			tok.Literal = lex.readIdentifier()
			tok.Type = token.LookupIdentifierType(tok.Literal)
			return tok
		} else if isDigit(lex.currentCharacter) {
			tok.Type = token.INT
			tok.Literal = lex.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lex.currentCharacter)
		}
	}

	lex.readCharacter()
	return tok
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.currentCharacter) {
		l.readCharacter()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhiteSpace() {
	for l.currentCharacter == ' ' || l.currentCharacter == '\t' || l.currentCharacter == '\n' || l.currentCharacter == '\r' {
		l.readCharacter()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.currentCharacter) {
		l.readCharacter()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) peekChar() byte {
	if l.readingPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readingPosition]
	}
}

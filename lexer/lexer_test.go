package lexer

import (
	"testing"

	"github.com/chayandatta/Hermes/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},1;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGNMENT, "="},
		{token.PLUS, "+"},
		{token.LEFT_PARENTHESIS, "("},
		{token.RIGHT_PARENTHESIS, ")"},
		{token.LEFT_BRACES, "{"},
		{token.RIGHT_BRACES, "}"},
		{token.LEFT_BRACKET, "["},
		{token.RIGHT_BRACKET, "]"},
		{token.COMMA, ","},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.EOF, "EOF"},
	}

	l := New(input)

	for idx, tokentext := range tests {
		token := l.NextToken()

		if token.Type != tokentext.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				idx, tokentext.expectedType, token.Type)
		}

		if token.Literal != tokentext.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				idx, tokentext.expectedLiteral, token.Literal)

		}
	}
}

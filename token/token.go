package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "Illegal"
	EOF     = "EOF"

	// Identifiers & literals
	IDENTIFIER = "IDENTIFIER" // add, foobar, x,y
	INT        = "Int"        // 1234567890
	STRING     = "String"     // "x", "y"

	// Operators
	ASSIGNMENT = "="
	PLUS       = "+"
	MINUS      = "-"
	BANG       = "!"
	ASTERISK   = "*"
	SLASH      = "/"
	EQUAL      = "=="
	NOT_EQUAL  = "!="

	LESS_THAN    = "<"
	GREATER_THAN = ">"

	// Delimeters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

	LEFT_PARENTHESIS  = "("
	RIGHT_PARENTHESIS = ")"
	LEFT_BRACES       = "{"
	RIGHT_BRACES      = "}"
	LEFT_BRACKET      = "["
	RIGHT_BRACKET     = "]"

	// Keywords
	FUNCTION = "Function"
	LET      = "Let"
	CONSTANT = "Const"
	TRUE     = "True"
	FALSE    = "False"
	IF       = "If"
	ELSE     = "Else"
	RETURN   = "Return"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdentifierType(identifier string) TokenType {
	if token, ok := keywords[identifier]; ok {
		return token
	}
	return IDENTIFIER
}

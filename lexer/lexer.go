package lexer

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

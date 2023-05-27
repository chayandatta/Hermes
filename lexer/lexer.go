package lexer

type Lexer struct {
    input string
    position int // current input position
    readingPosition int // current reading input position, after current position    
    currentCharacter byte // current character under examination
}

func New(input string) *Lexer {
    lex := &Lexer{input: input} // Lexer object creation
    return lex
}

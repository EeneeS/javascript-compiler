package lexer

import "unicode"

type tokenType int

const (
	EOF tokenType = iota
	Illegal
	Int
	String
	Plus
	Minus
	Multiply
	Divide
	Lparen
	Rparen
)

type Token struct {
	Type    tokenType
	Literal string
}

type Lexer struct {
	input         string
	position      int
	readPostition int
	currentChar   rune
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPostition >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = rune(l.input[l.readPostition])
	}
	l.position = l.readPostition
	l.readPostition++
}

func (l *Lexer) skipWhiteSpace() {
	for unicode.IsSpace(l.currentChar) {
		l.readChar()
	}
}

func (l *Lexer) NextToken() Token {
	var t Token
	return t
}

package lexer

import (
	"unicode"
)

type tokenType int

const (
	EOF tokenType = iota
	Illegal
	Int
	Plus
	Minus
	Multiply
	Divide
	Power
	Modulo
	Equals
	Identifier
	Keyword
	Lparen
	Rparen
)

type Token struct {
	Literal string
	Type    tokenType
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

// TODO implements floats
func (l *Lexer) readNumber() string {
	startPostition := l.position
	for unicode.IsDigit(l.currentChar) {
		l.readChar()
	}
	return l.input[startPostition:l.position]
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.position
	for unicode.IsLetter(l.currentChar) || l.currentChar == '_' {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

func (l *Lexer) lookUpIdentifier(identifier string) tokenType {
	keywords := []string{"log"}
	for _, kw := range keywords {
		if identifier == kw {
			return Keyword
		}
	}
	return Identifier
}

func (l *Lexer) NextToken() Token {
	var t Token

	l.skipWhiteSpace()

	switch l.currentChar {
	case '+':
		t = Token{Type: Plus, Literal: string(l.currentChar)}
	case '-':
		t = Token{Type: Minus, Literal: string(l.currentChar)}
	case '*':
		t = Token{Type: Multiply, Literal: string(l.currentChar)}
	case '/':
		t = Token{Type: Divide, Literal: string(l.currentChar)}
	case '^':
		t = Token{Type: Power, Literal: string(l.currentChar)}
	case '%':
		t = Token{Type: Modulo, Literal: string(l.currentChar)}
	case '=':
		t = Token{Type: Equals, Literal: string(l.currentChar)}
	case '(':
		t = Token{Type: Lparen, Literal: string(l.currentChar)}
	case ')':
		t = Token{Type: Rparen, Literal: string(l.currentChar)}
	default:
		if unicode.IsDigit(l.currentChar) {
			t.Literal = l.readNumber()
			t.Type = Int
		} else if unicode.IsLetter(l.currentChar) {
			t.Literal = l.readIdentifier()
			t.Type = l.lookUpIdentifier(t.Literal)
		}
	}

	l.readChar()

	return t
}

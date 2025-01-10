package lexer

import (
	"unicode"
)

type tokenType int

const (
	EOF tokenType = iota
	Illegal
	Int
	Float
	String
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
	LessThan
	GreaterThan
	Lbracket
	Rbracket
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

func (l *Lexer) readNumber() (string, tokenType) {
	startPostition := l.position
	tt := Int
	for unicode.IsDigit(l.currentChar) || l.currentChar == '.' {
		if l.currentChar == '.' {
			tt = Float
		}
		l.readChar()
	}
	return l.input[startPostition:l.position], tt
}

func (l *Lexer) readString() string {
	startPostition := l.position
	for l.currentChar != '"' {
		l.readChar()
	}
	return l.input[startPostition:l.position]
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.position
	for unicode.IsLetter(l.currentChar) || unicode.IsDigit(l.currentChar) || l.currentChar == '_' {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

func (l *Lexer) lookUpIdentifier(identifier string) tokenType {
	keywords := []string{
		"int", "string", "float",
		"if", "else",
		"and", "or",
		"function", "return",
	}
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
	case '<':
		t = Token{Type: LessThan, Literal: string(l.currentChar)}
	case '>':
		t = Token{Type: GreaterThan, Literal: string(l.currentChar)}
	case '{':
		t = Token{Type: Lbracket, Literal: string(l.currentChar)}
	case '}':
		t = Token{Type: Rbracket, Literal: string(l.currentChar)}
	case '"':
		t.Type = String
		t.Literal = l.readString()
		return t

	default:
		if unicode.IsDigit(l.currentChar) {
			t.Literal, t.Type = l.readNumber()
		} else if unicode.IsLetter(l.currentChar) {
			t.Literal = l.readIdentifier()
			t.Type = l.lookUpIdentifier(t.Literal)
		}
	}

	l.readChar()

	return t
}

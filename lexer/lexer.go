package lexer

import (
	"unicode"
)

type tokenType int

const (
	EOF tokenType = iota // 0
	Illegal
	Int
	Float
	String
	Plus // 5
	Minus
	Multiply
	Divide
	Power
	Modulo // 10
	Equals
	Identifier
	And
	Or
	Lparen // 15
	Rparen
	LessThan
	GreaterThan
	Lbracket
	Rbracket // 20
	Let
	Const
	If
	Else
	True // 25
	False
	Null
	Function
	Return
	Comma
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
	l.readChar()
	startPostition := l.position
	for l.currentChar != '"' {
		l.readChar()
	}
	l.readChar()
	return l.input[startPostition : l.position-1]
}

func (l *Lexer) readLogicalOp(op rune) (tt tokenType, tl string) {
	l.readChar()
	if l.currentChar == op {
		if op == '&' {
			tt = And
			tl = "&&"
		} else if op == '|' {
			tt = Or
			tl = "||"
		}
	} else {
		tt = Illegal
		tl = string(op)
	}
	l.readChar()
	return tt, tl
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.position
	for unicode.IsLetter(l.currentChar) || unicode.IsDigit(l.currentChar) || l.currentChar == '_' {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

func (l *Lexer) lookUpIdentifier(identifier string) tokenType {
	keywords := map[string]tokenType{
		"let":      Let,
		"const":    Const,
		"if":       If,
		"else":     Else,
		"true":     True,
		"false":    False,
		"null":     Null,
		"function": Function,
		"return":   Return,
	}
	if tokenType, exists := keywords[identifier]; exists {
		return tokenType
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
	case ',':
		t = Token{Type: Comma, Literal: string(l.currentChar)}
	case '&':
		tt, tl := l.readLogicalOp('&')
		t.Type = tt
		t.Literal = tl
		return t
	case '|':
		tt, tl := l.readLogicalOp('|')
		t.Type = tt
		t.Literal = tl
		return t
	case '"':
		t.Type = String
		t.Literal = l.readString()
		return t
	default:
		if unicode.IsDigit(l.currentChar) {
			t.Literal, t.Type = l.readNumber()
			return t
		} else if unicode.IsLetter(l.currentChar) {
			t.Literal = l.readIdentifier()
			t.Type = l.lookUpIdentifier(t.Literal)
			return t
		}
	}

	l.readChar()

	return t
}

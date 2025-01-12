package parser

import (
	"fmt"
	"github.com/eenees/slow/lexer"
)

type ASTNode interface{}

type IdentifierNode struct {
	Name string
}

type LiteralNode struct {
	Value string
}

type AssignmentNode struct {
	Identifier  IdentifierNode
	LiteralNode LiteralNode
}

type Parser struct {
	tokens  []lexer.Token
	current int
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{tokens: tokens, current: 0}
}

func (p *Parser) currentToken() lexer.Token {
	if p.current > len(p.tokens) {
		return lexer.Token{Type: lexer.EOF}
	}
	return p.tokens[p.current]
}

// Returns the current token, and moves up a token.
func (p *Parser) consume() lexer.Token {
	token := p.currentToken()
	p.current++
	return token
}

// Takes a look at the next token WIHOUT moving the token
func (p *Parser) peak() lexer.Token {
	if p.current >= len(p.tokens) {
		return lexer.Token{Type: lexer.EOF}
	}
	return p.tokens[p.current]
}

// Entry point
func (p *Parser) Parse() ASTNode {
	token := p.currentToken()
	switch token.Type {
	case lexer.Identifier:
		return p.parseIdentifier()
	case lexer.Keyword:
		return p.parseKeyword()
	default:
		panic(fmt.Sprintf("Unexpected token: %v", token.Type))
	}
}

func (p *Parser) parseIdentifier() ASTNode {
	token := p.consume()

	identifier := IdentifierNode{Name: token.Literal}

	token = p.peak()

	if token.Type == lexer.Equals {
		p.consume()
		return p.parseAssignment(identifier)
	}

	return identifier
}

func (p *Parser) parseAssignment(identifier IdentifierNode) ASTNode {
	token := p.consume()
	switch token.Type {
	case lexer.Int, lexer.Float, lexer.String:
		return AssignmentNode{
			Identifier:  identifier,
			LiteralNode: LiteralNode{Value: token.Literal},
		}
	}
	return nil
}

func (p *Parser) parseKeyword() ASTNode {
	// parse keywords like "function", "if", "else", "const", "let", ...
	return nil
}

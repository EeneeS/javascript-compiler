package parser

import (
	"fmt"

	"github.com/eenees/slow/lexer"
)

type Program struct {
	nodes []ASTNode
}

type ASTNode interface{}

type LiteralNode struct {
	value interface{} // int, float, string
}

type varType string

const (
	Const varType = "const"
	Let   varType = "let"
)

type VariableNode struct {
	name    string
	varType varType
	value   ASTNode
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

func (p *Parser) peak() lexer.Token {
	if p.current+1 > len(p.tokens) {
		return lexer.Token{Type: lexer.EOF}
	}
	return p.tokens[p.current+1]
}

func (p *Parser) consume() lexer.Token {
	token := p.currentToken()
	p.current++
	return token
}

func (p *Parser) Parse() Program {
	ast := Program{}
	for p.current < len(p.tokens) {
		token := p.currentToken()
		switch token.Type {
		case lexer.Identifier:
			node := p.parseIdentifier()
			ast.nodes = append(ast.nodes, node)
		default:
			p.consume()
			// fmt.Println("Unidentified token.")
		}
	}
	return ast
}

func (p *Parser) parseIdentifier() ASTNode {
	nextToken := p.peak()
	if nextToken.Type == lexer.Equals {
		return p.parseVariableNode(false)
	} else if nextToken.Type == lexer.Lparen {
		fmt.Println("parse function")
		p.consume()
		return p.parseFunctionCall()
	} else {
		p.consume()
	}
	return nil
}

func (p *Parser) parseVariableNode(isConst bool) ASTNode {
	currentToken := p.currentToken()
	p.consume() // this consumes the left side
	p.consume() // this consumes the '='
	value := p.currentToken()
	varType := Let
	if isConst {
		varType = Const
	}
	switch value.Type {
	case lexer.Int, lexer.Float, lexer.String:
		return VariableNode{
			name:    currentToken.Literal,
			varType: varType,
			value:   LiteralNode{value: value.Literal},
		}
	}
	return nil
}

func (p *Parser) parseFunctionCall() ASTNode {
	return nil
}

func (p *Parser) parseExpression() ASTNode {
	nextToken := p.currentToken()
	switch nextToken.Type {
	case lexer.Int, lexer.Float, lexer.String:
		return LiteralNode{value: nextToken.Literal}
	}
	return nil
}

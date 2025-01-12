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
	value interface{}
}

type VariableNode struct {
	name  string
	value ASTNode
}

type Parser struct {
	tokens  []lexer.Token
	current int
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{tokens: tokens, current: 0}
}

func (p *Parser) nextToken() lexer.Token {
	if p.current > len(p.tokens) {
		return lexer.Token{Type: lexer.EOF}
	}
	return p.tokens[p.current]
}

func (p *Parser) consume() lexer.Token {
	token := p.nextToken()
	p.current++
	return token
}

func (p *Parser) Parse() Program {
	ast := Program{}
	for p.current < len(p.tokens) {
		token := p.consume()
		switch token.Type {
		case lexer.Identifier:
			node := p.parseVariableNode(token.Literal)
			ast.nodes = append(ast.nodes, node)
		case lexer.Keyword:
			// fmt.Println("Keyword")
		default:
			// fmt.Println("Unidentified token.")
		}
	}
	return ast
}

func (p *Parser) parseVariableNode(name string) ASTNode {
	nextToken := p.nextToken()
	switch nextToken.Type {
	case lexer.Equals:
		p.consume()
		value := p.parseExpression()
		test := VariableNode{
			name:  name,
			value: value,
		}
		return test
	case lexer.Lparen:
		// check if next is also parent
	}
	return nil
}

func (p *Parser) parseExpression() ASTNode {
	nextToken := p.nextToken()
	switch nextToken.Type {
	case lexer.Int, lexer.Float, lexer.String:
		return LiteralNode{value: nextToken.Literal}
	}
	return nil
}

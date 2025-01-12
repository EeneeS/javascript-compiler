package parser

import (
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

func (p *Parser) consume() lexer.Token {
	token := p.currentToken()
	p.current++
	return token
}

func (p *Parser) Parse() Program {
	ast := Program{}
	for p.current < len(p.tokens) {
		token := p.consume()
		switch token.Type {
		case lexer.Identifier:
			node := p.parseIdentifier(token.Literal)
			ast.nodes = append(ast.nodes, node)
		case lexer.Keyword:
			node := p.parseKeyword()
			ast.nodes = append(ast.nodes, node)
		default:
			// fmt.Println("Unidentified token.")
		}
	}
	return ast
}

func (p *Parser) parseIdentifier(name string) ASTNode {
	currentToken := p.consume()
	if currentToken.Type == lexer.Equals {
		return p.parseVariableNode(name, false)
	} else if currentToken.Type == lexer.Lparen {
		return p.parseFunctionCall()
	}
	return nil
}

func (p *Parser) parseVariableNode(name string, isConst bool) ASTNode {
	currentToken := p.consume()
	varType := Let
	if isConst {
		varType = Const
	}
	switch currentToken.Type {
	case lexer.Int, lexer.Float, lexer.String:
		return VariableNode{
			name:    name,
			varType: varType,
			value:   LiteralNode{value: currentToken.Literal},
		}
	}
	return nil
}

func (p *Parser) parseKeyword() ASTNode {
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

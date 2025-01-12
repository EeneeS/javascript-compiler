package parser

import (
	"github.com/eenees/slow/lexer"
)

type Program struct {
	nodes []ASTNode
}

type ASTNode interface{}

type LiteralNode struct {
	value interface{}
}

type IdentifierNode struct {
	name string
}

type AssignmentNode struct {
	identifier IdentifierNode
	value      ASTNode
}

type Parser struct {
	tokens  []lexer.Token
	current int
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{tokens: tokens, current: 0}
}

func (p *Parser) Parse() Program {
	ast := Program{}
	return ast
}

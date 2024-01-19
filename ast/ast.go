package ast

import "evasion/token"

// Root node of AST
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// All nodes in the AST must implement the Node interface
type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// Satisfies the Statement interface
func (ls *LetStatement) statementNode() {}

// Satisfies the Node interface
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

// Satisfies the Expression interface
func (i *Identifier) expressionNode() {}

// Satisfies the Node interface
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

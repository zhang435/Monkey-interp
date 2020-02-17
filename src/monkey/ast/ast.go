package ast

import "monkey/token"

// Node interface
type Node interface {
	TokenLiteral() string
}

//Statement :
type Statement interface {
	Node
	statementNode()
}

// Expression :
type Expression interface {
	Node
	expressionNode()
}

//Program : the actual struct
type Program struct {
	Statements []Statement
}

// LetStatement :
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// Identifier :
type Identifier struct {
	Token token.Token
	Value string
}

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// TokenLiteral :  TokenLiteral
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *Identifier) expressionNode() {}
func (ls *Identifier) TokenLiteral() string {
	return ls.Token.Literal
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

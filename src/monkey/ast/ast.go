package ast

import (
	"bytes"
	"monkey/token"
)

// this define the node of the syntax tree.

// Node interface
type Node interface {
	TokenLiteral() string
	String() string
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

// TokenLiteral :  TokenLiteral
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

///////////////////////////////////////////////////////////////

// LetStatement :
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral  : get the token string
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	} else {
		out.WriteString("nil")
	}
	out.WriteString(";")
	return out.String()
}

///////////////////////////////////////////////////////////////

// Identifier :
type Identifier struct {
	Token token.Token
	Value string
}

func (l *Identifier) expressionNode() {}

// TokenLiteral  : get the token string
func (l *Identifier) TokenLiteral() string {
	return l.Token.Literal
}

func (l *Identifier) String() string {
	return l.Value
}

///////////////////////////////////////////////////////////////

// ReturnStatement :
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral  : get the token string
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	} else {
		out.WriteString("nil")
	}
	out.WriteString(";")
	return out.String()
}

///////////////////////////////////////////////////////////////

// ExpressionStatement :
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral  : get the token string
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

//////////////////////////////////////////////////////
// define int node
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (ie *IntegerLiteral) expressionNode() {}

//TokenLiteral get the literal token
func (ie *IntegerLiteral) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IntegerLiteral) String() string {
	return ie.Token.Literal
}

//PrefixExpression this struct is part of experssion to represents !4 -5 ...
type PrefixExpression struct {
	Token    token.Token // the prefix token !
	Operator string
	Right    Expression
}

//ExpressionNode prefix expression is a expression
func (pe *PrefixExpression) expressionNode() {}

//TokenLiteral get the token of the experssion
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

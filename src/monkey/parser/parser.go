package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"strconv"
)

const (
	_           int = iota
	LOWEST          //??
	EQUALS          // ==
	LESSGREATER     // > <
	SUM
	PRODUCT // *
	PREFIX  //-x +x !x
	CALL
)

// parsing function
type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// parser: parser type
type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  token.Token
	peekToken token.Token

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// constructor
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.IDENT, p.parserIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExperssion)
	p.registerPrefix(token.PLUS, p.parsePrefixExperssion)
	p.registerPrefix(token.MINUS, p.parsePrefixExperssion)

	p.nextToken()
	p.nextToken()

	return p
}

// LetStatement : get two pointers in  lexer
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// parseProgram : walk through the program
func (p *Parser) parseProgram() *ast.Program {
	program := &ast.Program{Statements: []ast.Statement{}}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{p.curToken, p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// todo semi column , we are skipping semi column for now
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}
	return stmt

}
func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekErrors(t)
	return false
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekErrors(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, but got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}
	stmt.Expression = p.parserExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parserExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		return nil
	}

	leftExp := prefix()
	return leftExp
}

func (p *Parser) parserIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	i := &ast.IntegerLiteral{Token: p.curToken}

	value, error := strconv.ParseInt(p.curToken.Literal, 0, 64)

	if error != nil {
		msg := fmt.Sprintf("try to parse a int error, but got %q  ", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	i.Value = value
	return i

}

func (p *Parser) noPredixFunctionFound(t token.TokenType) {
	msg := fmt.Sprintf("Unable to find bind parser for prefix %s", t)
	p.errors = append(p.errors, msg)

}

func (p *Parser) parseExperssion(precedence int) ast.Expression {
	fn := p.prefixParseFns[p.curToken.Type]
	if fn == nil {
		p.noPredixFunctionFound(p.curToken.Type)
	}

	return fn()
}

func (p *Parser) parsePrefixExperssion() ast.Expression {
	expr := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expr.Right = p.parseExperssion(PREFIX)

	return expr
}

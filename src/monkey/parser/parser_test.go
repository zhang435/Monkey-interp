package parser

import (
	"fmt"
	"monkey/ast"
	"monkey/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
		let x = 5;
		let y = 10;
		let foobar = 7890;
	`
	l := lexer.New(input)

	p := New(l)

	program := p.parseProgram()
	checkPraseErrors(t, p)

	if nil == program {
		t.Fatalf("ParseProgram() return nil")
	}

	StatementsCount := 3
	if len(program.Statements) != StatementsCount {
		t.Fatalf("ParseProgram() should have %d statements, got = %d", StatementsCount, len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("let statement need to start with let, got = %s", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not type of Letstatement got = %T", s)
		return false
	}
	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value expected %s, but got %s", letStmt.Name.Value, name)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() expected %s, but got %s", letStmt.Name.TokenLiteral(), name)
		return false
	}

	return true
}

func TestReturnStatements(t *testing.T) {
	input := `
		return 5;
		return 10;
		return 1212121;
	`
	l := lexer.New(input)

	p := New(l)

	program := p.parseProgram()
	checkPraseErrors(t, p)

	statementCount := 3
	if len(program.Statements) != statementCount {
		t.Fatalf("ParseProgram() should have %d statements, got = %d", statementCount, len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("stmt not *ast.ReturnStatement got=%T", returnStmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral() is not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func checkPraseErrors(t *testing.T, p *Parser) {
	if len(p.Errors()) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(p.errors))
	for _, msg := range p.Errors() {
		t.Errorf("parser error: %s", msg)
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"
	l := lexer.New(input)

	p := New(l)

	program := p.parseProgram()

	checkPraseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statments, expected %d, got %d", 1, len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program 0 is not a experssion, got %T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("exp not *Ast.ident, got %T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value is not foo bar, got %s", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenIteral is not foobar, got %s", ident.TokenLiteral())
	}
}

func TestIntegerLiteralExperssion(t *testing.T) {
	input := "5;"

	l := lexer.New(input)

	p := New(l)

	program := p.parseProgram()

	checkPraseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program parse didn't return enough statements expected 1, got %d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("result is not in type ExpressionStatement, got %T", stmt.Expression)
	}

	il, ok := stmt.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("result is not in type IntegerLiteral, got %T", stmt.Expression)
		print(il)
	}

	if il.Value != 5 {
		t.Fatalf("Value expected 5, got %d", il.Value)
	}

	if il.TokenLiteral() != "5" {
		t.Fatalf("TokenLiteral expected '5', got %s", il.TokenLiteral())
	}
}

func TestParsingPrefixExpression(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.parseProgram()
		checkPraseErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program parse didn't return enough statements expected 1, got %d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("result is not in type ExpressionStatement, got %T", stmt.Expression)
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("result is not in type prefixExperssion, got %T", exp)
			print(exp)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("Exp.Operator is not %s, got %s", tt.operator, exp.Operator)
		}

		if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
			return
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integer, ok := il.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf("il is not IntegerLiteral, got %T", il)
		return false
	}

	if integer.Value != value {
		t.Errorf("Expected %d, but got %d", value, integer.Value)
		return false
	}

	if integer.Token.Literal != fmt.Sprintf("%d", value) {
		t.Errorf("integer.TokenLiteral is not %d, got %s", value, integer.Token.Literal)
		return false
	}
	return true
}

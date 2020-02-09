package parser

import (
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

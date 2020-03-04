package ast

import (
	"monkey/token"
	"testing"
)

func TestString(t *testing.T) {
	input := "let myVar = anotherVar;"
	Program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if Program.String() != input {
		t.Errorf("deseralized is not expected, got %s", Program.String())
	}
}

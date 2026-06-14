package ast_test

import (
	"testing"

	"github.com/chadsmith12/monkey-interpreter/ast"
	"github.com/chadsmith12/monkey-interpreter/token"
)

func TestString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "var"},
					Value: "var",
				},
				Value: &ast.Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "let var = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}

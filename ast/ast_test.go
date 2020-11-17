package ast

import (
	"testing"

	"github.com/slier81/monkey/token"
)

func TestString(t *testing.T) {
	input := `let myVar = anotherVar;`
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{
					Type:    token.LET,
					Literal: "let",
				},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENTIFIER, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{
						Type:    token.IDENTIFIER,
						Literal: "anotherVar",
					},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != input {
		t.Errorf("program.String() is wrong..Expecting %q, got=%q", input, program.String())
	}
}

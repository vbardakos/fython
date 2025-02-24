package parser

import (
	"github.com/vbardakos/fython/ast"
	"github.com/vbardakos/fython/lexer"
	"testing"
)

func TestStatements(t *testing.T) {
	input := `
five = 5
ten = 10
seven = 7
`
	lxr := lexer.New(input)
	parser := New(lxr)

	program := parser.ParseProgram()
	if program == nil {
		t.Fatalf("ParseProgram returns nil")
	}

	lines := 3
	if len(program.Statements) != lines {
		t.Fatalf("ParseProgram.Stmt Len Error. exp=%d, got=%d",
			lines, len(program.Statements))
	}

	tests := []struct {
		expIdentifier string
	}{
		{"five"},
		{"ten"},
		{"seven"},
	}

	for idx, tt := range tests {
		stmt := program.Statements[idx]
		assertAssign(t, stmt, tt.expIdentifier)
	}
}

func assertAssign(t *testing.T, s ast.Stmt, name string) {
	if s.TokenLiteral() != "ASSIGN" {
		t.Fatalf("Stmt is not Assign. got=%T", s)
	}

	stmt, ok := s.(*ast.AssignStmt)

	if !ok {
		t.Fatalf("Stmt is not Assign. got=%T", s)
	}

	if stmt.Name.Value != name {
		t.Fatalf("Assign.Name.Value error. exp=%s, got=%s",
			name, stmt.Name.Value)
	}

	if stmt.Name.TokenLiteral() != name {
		t.Fatalf("Assign.Name.TokenLiteral error. exp=%s, got=%s",
			name, stmt.Name.TokenLiteral())
	}
}

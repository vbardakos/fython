package parser

import (
	"testing"

	"github.com/vbardakos/fython/ast"
	"github.com/vbardakos/fython/lexer"
	"github.com/vbardakos/fython/token"
)

func TestAssignStatements(t *testing.T) {
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
	checkParserErrors(t, parser)
}

func TestReturnStatements(t *testing.T) {
	input := `
return five
return 10
return False
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

	tests := []token.Token{
		{Token: token.RETURN, Literal: "return"},
		{Token: token.RETURN, Literal: "return"},
		{Token: token.RETURN, Literal: "return"},
	}

	for idx, tkn := range tests {
		stmt := program.Statements[idx]
		assertReturn(t, stmt, tkn)
	}
	checkParserErrors(t, parser)
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

func assertReturn(t *testing.T, s ast.Stmt, tkn token.Token) {
	if s.TokenLiteral() != tkn.Literal {
		t.Fatalf("Stmt.TokenLiteral err. exp=%s. got=%s",
			tkn.Literal, s.TokenLiteral())
	}

	stmt, ok := s.(*ast.ReturnStmt)

	if !ok {
		t.Fatalf("Stmt is not RETURN. got=%T", s)
	}

	if stmt.Token.Token != tkn.Token {
		t.Fatalf("Return Token error. exp=%s, got=%s",
			stmt.Token.Token, tkn.Token)
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errs := p.Errors()

	if len(errs) == 0 {
		return
	}

	for _, err := range errs {
		t.Errorf("%s", err.Show())
	}
	t.Fail()
}

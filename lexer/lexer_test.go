package lexer

import (
	"github.com/vbardakos/fython/token"
	"testing"
)

func assertToken(t *testing.T, idx int, exp token.Token, act token.Token) {
	if exp.Token != act.Token {
		t.Fatalf("tests[%d] - Token Error. expected=%q, got=%q",
			idx, exp.Token, act.Token)
	}

	if exp.Literal != act.Literal {
		t.Fatalf("tests[%d] - Literal Error. expected=%q, got=%q",
			idx, exp.Literal, act.Literal)
	}
}

func TestBaseExpressions(t *testing.T) {
	input := `five = 5
ten = 10

def add(x, y):
    return x + y

result = add(five, ten)
`
	tests := []token.Token{
		{Token: token.IDENT, Literal: "five"},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.INT, Literal: "5"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.IDENT, Literal: "ten"},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.INT, Literal: "10"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.FUNCTION, Literal: "def"},
		{Token: token.IDENT, Literal: "add"},
		{Token: token.LPAREN, Literal: "("},
		{Token: token.IDENT, Literal: "x"},
		{Token: token.COMMA, Literal: ","},
		{Token: token.IDENT, Literal: "y"},
		{Token: token.RPAREN, Literal: ")"},
		{Token: token.COLON, Literal: ":"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.INDENT, Literal: "    "},
		{Token: token.RETURN, Literal: "return"},
		{Token: token.IDENT, Literal: "x"},
		{Token: token.ADD, Literal: "+"},
		{Token: token.IDENT, Literal: "y"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.IDENT, Literal: "result"},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.IDENT, Literal: "add"},
		{Token: token.LPAREN, Literal: "("},
		{Token: token.IDENT, Literal: "five"},
		{Token: token.COMMA, Literal: ","},
		{Token: token.IDENT, Literal: "ten"},
		{Token: token.RPAREN, Literal: ")"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.EOF, Literal: ""},
	}

	lxr := New(input)
	for i, tt := range tests {
		tkn := lxr.NextToken()
		assertToken(t, i, tt, tkn)
	}
}

func TestOperators(t *testing.T) {
	input := `five = 5
ten = 10

x = ten - five
y = ten * five
z = ten / five
`

	tests := []token.Token{
		{Token: token.IDENT, Literal: "five"},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.INT, Literal: "5"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.IDENT, Literal: "ten"},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.INT, Literal: "10"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.IDENT, Literal: "x"},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.IDENT, Literal: "ten"},
		{Token: token.SUB, Literal: "-"},
		{Token: token.IDENT, Literal: "five"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.IDENT, Literal: "y"},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.IDENT, Literal: "ten"},
		{Token: token.MUL, Literal: "*"},
		{Token: token.IDENT, Literal: "five"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.IDENT, Literal: "z"},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.IDENT, Literal: "ten"},
		{Token: token.DIV, Literal: "/"},
		{Token: token.IDENT, Literal: "five"},
		{Token: token.EOL, Literal: "\n"},
	}

	lxr := New(input)
	for idx, tt := range tests {
		tkn := lxr.NextToken()
		assertToken(t, idx, tt, tkn)
	}
}

func TestCondintions(t *testing.T) {
	input := `x = 5
y = 10
z = 20


if x <= y >= z:
    x = x * 2
elif x < y > z:
    y = y / 2
else:
    z = 0
`

	tests := []token.Token{
		{Token: token.IDENT, Literal: "x"},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.INT, Literal: "5"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.IDENT, Literal: "y"},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.INT, Literal: "10"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.IDENT, Literal: "z"},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.INT, Literal: "20"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.IF, Literal: "if"},
		{Token: token.IDENT, Literal: "x"},
		{Token: token.LE, Literal: "<="},
		{Token: token.IDENT, Literal: "y"},
		{Token: token.GE, Literal: ">="},
		{Token: token.IDENT, Literal: "z"},
		{Token: token.COLON, Literal: ":"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.INDENT, Literal: "    "},
		{Token: token.IDENT, Literal: "x"},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.IDENT, Literal: "x"},
		{Token: token.MUL, Literal: "*"},
		{Token: token.INT, Literal: "2"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.ELIF, Literal: "elif"},
		{Token: token.IDENT, Literal: "x"},
		{Token: token.LT, Literal: "<"},
		{Token: token.IDENT, Literal: "y"},
		{Token: token.GT, Literal: ">"},
		{Token: token.IDENT, Literal: "z"},
		{Token: token.COLON, Literal: ":"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.INDENT, Literal: "    "},
		{Token: token.IDENT, Literal: "y"},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.IDENT, Literal: "y"},
		{Token: token.DIV, Literal: "/"},
		{Token: token.INT, Literal: "2"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.ELSE, Literal: "else"},
		{Token: token.COLON, Literal: ":"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.INDENT, Literal: "    "},
		{Token: token.IDENT, Literal: "z"},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.INT, Literal: "0"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.EOF, Literal: ""},
	}

	lxr := New(input)
	for idx, tt := range tests {
		tkn := lxr.NextToken()
		assertToken(t, idx, tt, tkn)
	}
}

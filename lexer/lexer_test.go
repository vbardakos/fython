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

func assertCount(t *testing.T, idx int, tkn token.Token, exp *Counter, act Counter) {
	if tkn.Token == token.EOL {
		exp.column = 0
		exp.line += 1
	} else {
		exp.column += len(tkn.Literal)
	}

	if act != *exp {
		t.Fatalf("tests[%d] - Counter Error. expected=%d, got=%d @ Token %q",
			idx, exp, act, tkn.Token)
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
		{Token: token.SPC, Literal: " "},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.SPC, Literal: " "},
		{Token: token.INT, Literal: "5"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.IDENT, Literal: "ten"},
		{Token: token.SPC, Literal: " "},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.SPC, Literal: " "},
		{Token: token.INT, Literal: "10"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.FUNCTION, Literal: "def"},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "add"},
		{Token: token.LPAREN, Literal: "("},
		{Token: token.IDENT, Literal: "x"},
		{Token: token.COMMA, Literal: ","},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "y"},
		{Token: token.RPAREN, Literal: ")"},
		{Token: token.COLON, Literal: ":"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.SPC, Literal: " "},
		{Token: token.SPC, Literal: " "},
		{Token: token.SPC, Literal: " "},
		{Token: token.SPC, Literal: " "},
		{Token: token.RETURN, Literal: "return"},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "x"},
		{Token: token.SPC, Literal: " "},
		{Token: token.ADD, Literal: "+"},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "y"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.IDENT, Literal: "result"},
		{Token: token.SPC, Literal: " "},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "add"},
		{Token: token.LPAREN, Literal: "("},
		{Token: token.IDENT, Literal: "five"},
		{Token: token.COMMA, Literal: ","},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "ten"},
		{Token: token.RPAREN, Literal: ")"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.EOF, Literal: ""},
	}

	lxr := New(input)
	counter := new(Counter)

	for i, tt := range tests {
		tkn := lxr.NextToken()
		assertToken(t, i, tt, tkn)
		assertCount(t, i, tkn, counter, lxr.counter)
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
		{Token: token.SPC, Literal: " "},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.SPC, Literal: " "},
		{Token: token.INT, Literal: "5"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.IDENT, Literal: "ten"},
		{Token: token.SPC, Literal: " "},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.SPC, Literal: " "},
		{Token: token.INT, Literal: "10"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.IDENT, Literal: "x"},
		{Token: token.SPC, Literal: " "},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "ten"},
		{Token: token.SPC, Literal: " "},
		{Token: token.SUB, Literal: "-"},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "five"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.IDENT, Literal: "y"},
		{Token: token.SPC, Literal: " "},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "ten"},
		{Token: token.SPC, Literal: " "},
		{Token: token.MUL, Literal: "*"},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "five"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.IDENT, Literal: "z"},
		{Token: token.SPC, Literal: " "},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "ten"},
		{Token: token.SPC, Literal: " "},
		{Token: token.DIV, Literal: "/"},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "five"},
		{Token: token.EOL, Literal: "\n"},
	}

	lxr := New(input)
	counter := new(Counter)


	for idx, tt := range tests {
		tkn := lxr.NextToken()
		assertToken(t, idx, tt, tkn)
		assertCount(t, idx, tt, counter, lxr.counter)
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
		{Token: token.SPC, Literal: " "},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.SPC, Literal: " "},
		{Token: token.INT, Literal: "5"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.IDENT, Literal: "y"},
		{Token: token.SPC, Literal: " "},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.SPC, Literal: " "},
		{Token: token.INT, Literal: "10"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.IDENT, Literal: "z"},
		{Token: token.SPC, Literal: " "},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.SPC, Literal: " "},
		{Token: token.INT, Literal: "20"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.IF, Literal: "if"},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "x"},
		{Token: token.SPC, Literal: " "},
		{Token: token.LE, Literal: "<="},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "y"},
		{Token: token.SPC, Literal: " "},
		{Token: token.GE, Literal: ">="},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "z"},
		{Token: token.COLON, Literal: ":"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.SPC, Literal: " "},
		{Token: token.SPC, Literal: " "},
		{Token: token.SPC, Literal: " "},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "x"},
		{Token: token.SPC, Literal: " "},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "x"},
		{Token: token.SPC, Literal: " "},
		{Token: token.MUL, Literal: "*"},
		{Token: token.SPC, Literal: " "},
		{Token: token.INT, Literal: "2"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.ELIF, Literal: "elif"},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "x"},
		{Token: token.SPC, Literal: " "},
		{Token: token.LT, Literal: "<"},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "y"},
		{Token: token.SPC, Literal: " "},
		{Token: token.GT, Literal: ">"},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "z"},
		{Token: token.COLON, Literal: ":"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.SPC, Literal: " "},
		{Token: token.SPC, Literal: " "},
		{Token: token.SPC, Literal: " "},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "y"},
		{Token: token.SPC, Literal: " "},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "y"},
		{Token: token.SPC, Literal: " "},
		{Token: token.DIV, Literal: "/"},
		{Token: token.SPC, Literal: " "},
		{Token: token.INT, Literal: "2"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.ELSE, Literal: "else"},
		{Token: token.COLON, Literal: ":"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.SPC, Literal: " "},
		{Token: token.SPC, Literal: " "},
		{Token: token.SPC, Literal: " "},
		{Token: token.SPC, Literal: " "},
		{Token: token.IDENT, Literal: "z"},
		{Token: token.SPC, Literal: " "},
		{Token: token.ASSIGN, Literal: "="},
		{Token: token.SPC, Literal: " "},
		{Token: token.INT, Literal: "0"},
		{Token: token.EOL, Literal: "\n"},
		{Token: token.EOF, Literal: ""},
	}

	lxr := New(input)
	counter := new(Counter)


	for idx, tt := range tests {
		tkn := lxr.NextToken()
		assertToken(t, idx, tt, tkn)
		assertCount(t, idx, tt, counter, lxr.counter)
	}
}

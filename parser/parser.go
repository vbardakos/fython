package parser

import (
	"fmt"

	"github.com/vbardakos/fython/ast"
	"github.com/vbardakos/fython/lexer"
	"github.com/vbardakos/fython/token"
)

type Parser struct {
	lxr  *lexer.Lexer
	errs []Error

	curr token.Token
	peek token.Token
}

type Error struct {
	expect token.TokenType
	actual token.Token
	line   int
	column int
}

func (e *Error) Print() string {
	msg := fmt.Sprintf("[%d:%d] exp=%q, got=%q",
		e.line, e.column, e.expect, e.actual)
	return msg
}

func New(lxr *lexer.Lexer) *Parser {
	parser := &Parser{
		lxr:  lxr,
		errs: []Error{},
	}

	parser.nextToken()
	parser.nextToken()
	return parser
}

func (p *Parser) nextToken() {
	p.curr = p.peek
	p.peek = p.lxr.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Stmt{}

	for p.curr.Token != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Stmt {
	switch p.curr.Token {
	case token.IDENT:
		return p.parseAssignStmt()
	default:
		return nil
	}
}

func (p *Parser) parseAssignStmt() ast.Stmt {
	ident := &ast.Identifier{Token: p.curr, Value: p.curr.Literal}
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	stmt := &ast.AssignStmt{Token: p.curr, Name: ident}

	p.nextToken()
	return stmt
}

func (p *Parser) currIs(tt token.TokenType) bool {
	return p.curr.Token == tt
}

func (p *Parser) peekIs(tt token.TokenType) bool {
	return p.peek.Token == tt
}

func (p *Parser) expectPeek(tt token.TokenType) bool {
	if p.peekIs(tt) {
		p.nextToken()
		return true
	}
	return false
}

func (p *Parser) Errors() []Error {
	return p.errs
}

func (p *Parser) peekError(tt token.TokenType) {
	line, col := p.lxr.GetPosition()
	err := Error{
		expect: tt,
		actual: p.peek,
		line:   line,
		column: col,
	}
	p.errs = append(p.errs, err)
}

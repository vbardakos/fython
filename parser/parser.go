package parser

import (
	"github.com/vbardakos/fython/ast"
	"github.com/vbardakos/fython/lexer"
	"github.com/vbardakos/fython/token"
)

type Parser struct {
	lxr *lexer.Lexer

	curr token.Token
	peek token.Token
}

func New(lxr *lexer.Lexer) *Parser {
	parser := &Parser{lxr: lxr}

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
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	stmt := &ast.AssignStmt{Token: p.peek}

	stmt.Name = &ast.Identifier{Token: p.curr, Value: p.curr.Literal}

	p.nextToken()
	return stmt
}

func (p *Parser) expectPeek(tkn token.TokenType) bool {
	return p.peek.Token == tkn
}

func (p *Parser) expectCurr(tkn token.TokenType) bool {
	return p.curr.Token == tkn
}

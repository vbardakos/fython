package parser

import (
	// "github.com/vbardakos/fython/ast"
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

package ast

import "github.com/vbardakos/fython/token"

type Node interface {
	TokenLiteral() string
}

type Stmt interface {
	Node
	stmtNode()
}

type Expr interface {
	Node
	exprNode()
}

type Program struct {
	Statements []Stmt
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) exprNode()
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type AssignStmt struct {
	Token token.Token
	Name  *Identifier
	Value Expr
}

func (s *AssignStmt) stmtNode()            {}
func (s *AssignStmt) TokenLiteral() string { return "ASSIGN" }

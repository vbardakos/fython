package lexer

import (
	"github.com/vbardakos/fython/token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         byte
	counter      Counter
}

type Counter struct {
	line   int
	column int
}

func (c *Counter) addLine() {
	c.line += 1
	c.column = -1
}

func (c *Counter) addColumn() {
	c.column += 1
}

func New(input string) *Lexer {
	c := Counter{line: 0, column: -1}
	lxr := &Lexer{input: input, counter: c}
	lxr.readChar()
	return lxr
}

func (lxr *Lexer) readChar() {
	lxr.counter.addColumn()
	lxr.position = lxr.readPosition

	if lxr.readPosition >= len(lxr.input) {
		lxr.char = 0
		return
	}

	lxr.char = lxr.input[lxr.readPosition]
	lxr.readPosition += 1
}

func (lxr *Lexer) peekChar(idx int) byte {
	peek := lxr.readPosition + idx
	if peek >= len(lxr.input) {
		return 0
	}
	return lxr.input[peek]
}

func (lxr *Lexer) NextToken() token.Token {
	var tkn token.Token

	switch lxr.char {
	case ' ':
		tkn = newToken(token.SPC, lxr.char)
	case '=':
		if lxr.peekChar(0) == '=' {
			curr := lxr.char
			lxr.readChar()
			literal := string(curr) + string(lxr.char)
			tkn = token.Token{Token: token.EQ, Literal: literal}
		} else {
			tkn = newToken(token.ASSIGN, lxr.char)
		}
	case '+':
		tkn = newToken(token.ADD, lxr.char)
	case '-':
		tkn = newToken(token.SUB, lxr.char)
	case '*':
		tkn = newToken(token.MUL, lxr.char)
	case '/':
		tkn = newToken(token.DIV, lxr.char)
	case '!':
		if lxr.peekChar(0) == '=' {
			curr := lxr.char
			lxr.readChar()
			literal := string(curr) + string(lxr.char)
			tkn = token.Token{Token: token.NE, Literal: literal}
		} else {
			tkn = newToken(token.ILLEGAL, lxr.char)
		}
	case '(':
		tkn = newToken(token.LPAREN, lxr.char)
	case ')':
		tkn = newToken(token.RPAREN, lxr.char)
	case '>':
		if lxr.peekChar(0) == '=' {
			curr := lxr.char
			lxr.readChar()
			literal := string(curr) + string(lxr.char)
			tkn = token.Token{Token: token.GE, Literal: literal}
		} else {
			tkn = newToken(token.GT, lxr.char)
		}
	case '<':
		if lxr.peekChar(0) == '=' {
			curr := lxr.char
			lxr.readChar()
			literal := string(curr) + string(lxr.char)
			tkn = token.Token{Token: token.LE, Literal: literal}
		} else {
			tkn = newToken(token.LT, lxr.char)
		}
	case '{':
		tkn = newToken(token.LBRACE, lxr.char)
	case '}':
		tkn = newToken(token.RBRACE, lxr.char)
	case ',':
		tkn = newToken(token.COMMA, lxr.char)
	case ';':
		tkn = newToken(token.SEMICOLON, lxr.char)
	case ':':
		tkn = newToken(token.COLON, lxr.char)
	case '\n':
		lxr.counter.addLine()
		tkn = newToken(token.EOL, lxr.char)
	case 0:
		tkn.Literal = ""
		tkn.Token = token.EOF
		return tkn
	default:
		if isLetter(lxr.char) {
			tkn.Literal = lxr.readIdentifier()
			tkn.Token = token.LookupKeyword(tkn.Literal)
			return tkn
		}
		if isDigit(lxr.char) {
			tkn.Literal = lxr.readNumber()
			tkn.Token = token.INT
			return tkn
		}
		tkn = newToken(token.ILLEGAL, lxr.char)
	}

	lxr.readChar()
	return tkn
}

func (lxr *Lexer) readIdentifier() string {
	position := lxr.position
	for isLetter(lxr.char) {
		lxr.readChar()
	}
	return lxr.input[position:lxr.position]
}

func (lxr *Lexer) readNumber() string {
	position := lxr.position
	for isDigit(lxr.char) {
		lxr.readChar()
	}
	return lxr.input[position:lxr.position]
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Token: tokenType, Literal: string(char)}
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}

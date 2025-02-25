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

func (c *Counter) moveLine() {
	c.line += 1
	c.column = -1
}

func (c *Counter) moveCol() {
	c.column += 1
}

func New(input string) *Lexer {
	c := Counter{line: 0, column: -1}
	lxr := &Lexer{input: input, counter: c}
	lxr.readChar()
	return lxr
}

func (lxr *Lexer) readChar() {
	lxr.counter.moveCol()
	lxr.position = lxr.readPosition

	if lxr.readPosition >= len(lxr.input) {
		lxr.char = 0
		return
	}

	lxr.char = lxr.input[lxr.readPosition]
	lxr.readPosition += 1
}

func (lxr *Lexer) peekChar(idx int) byte {
	peek := lxr.position + idx
	if peek >= len(lxr.input) {
		return 0
	}
	return lxr.input[peek]
}

func (lxr *Lexer) NextToken() token.Token {
	var tkn token.Token

	switch lxr.char {
	case ' ':
		if isIndent(lxr) {
			lit := string(lxr.char)
			for range 3 {
				lxr.readChar()
				lit += string(lxr.char)
			}
			tkn = token.Token{Token: token.INDENT, Literal: lit}
		} else {
			lxr.readChar()
			return lxr.NextToken()
		}
	case '=':
		if lxr.peekChar(1) == '=' {
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
		if lxr.peekChar(1) == '=' {
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
		if lxr.peekChar(1) == '=' {
			curr := lxr.char
			lxr.readChar()
			literal := string(curr) + string(lxr.char)
			tkn = token.Token{Token: token.GE, Literal: literal}
		} else {
			tkn = newToken(token.GT, lxr.char)
		}
	case '<':
		if lxr.peekChar(1) == '=' {
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
		lxr.counter.moveLine()
		tkn = newToken(token.EOL, lxr.char)
	case 0:
		tkn.Literal = ""
		tkn.Token = token.EOF
		return tkn
	default:
		if lxr.isBytes() {
			tkn.Literal = lxr.readBytes()
			tkn.Token = token.BYTES
			return tkn
		}
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
		if isQuote(lxr.char) {
			tkn.Literal = lxr.readString()
			tkn.Token = token.STR
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
	// todo: add exponents, octals, bins
	// fixme: fix illegal cases
	position := lxr.position
	for isDigit(lxr.char) || lxr.char == '.' {
		lxr.readChar()
	}
	return lxr.input[position:lxr.position]
}

func (lxr *Lexer) readString() string {
	quote := lxr.char
	if quote == lxr.peekChar(1) {
		return ""
	}
	lxr.readChar()       // read through quote
	defer lxr.readChar() // read closing quote

	position := lxr.position
	for lxr.char != quote || lxr.peekChar(-1) == '\\' {
		lxr.readChar()
	}
	return lxr.input[position:lxr.position]
}

func (lxr *Lexer) readBytes() string {
	lxr.readChar() // move to quote
	return lxr.readString()
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

func isQuote(char byte) bool {
	return char == '\'' || char == '"'
}

func (lxr *Lexer) isBytes() bool {
	return lxr.char == 'b' && isQuote(lxr.peekChar(1))
}

func isIndent(lxr *Lexer) bool {
	col := lxr.counter.column

	if col%4 != 0 {
		return false
	}

	// peek line start until 3 infront
	for i := -col; i < 3; i++ {
		if lxr.peekChar(i) != ' ' {
			return false
		}
	}
	return true
}

func (lxr *Lexer) GetPosition() (int, int) {
	return lxr.counter.line, lxr.counter.column
}

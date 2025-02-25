package token

type TokenType string

type Token struct {
	Token   TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	EOL     = "EOL"
	INDENT  = "INDENT"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	BOOL  = "BOOL"  // True/False
	INT   = "INT"   // 1343456
	BYTES = "BYTES" // b'\x00'
	FLOAT = "FLOAT" // 1.21345
	STR   = "STR"   // "abcde"

	// Operators
	ASSIGN = "="
	ADD    = "+"
	SUB    = "-"
	MUL    = "*"
	DIV    = "/"
	BANG   = "!"

	LT = "<"
	GT = ">"
	EQ = "=="
	NE = "!="
	GE = "<="
	LE = ">="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	SQUOTE    = "'"
	DQUOTE    = "\""

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	RETURN   = "RETURN"
	NOT      = "NOT"
	AND      = "AND"
	OR       = "OR"
	IF       = "IF"
	ELIF     = "ELIF"
	ELSE     = "ELSE"
)

var keywords = map[string]TokenType{
	"def":    FUNCTION,
	"return": RETURN,
	"not":    NOT,
	"and":    AND,
	"or":     OR,
	"if":     IF,
	"elif":   ELIF,
	"else":   ELSE,
	"True":   BOOL,
	"False":  BOOL,
}

func LookupKeyword(ident string) TokenType {
	if tkn, ok := keywords[ident]; ok {
		return tkn
	}
	return IDENT
}

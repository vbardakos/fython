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

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

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
	SPC       = " "
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
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
	"True":   TRUE,
	"False":  FALSE,
}

func LookupKeyword(ident string) TokenType {
	if tkn, ok := keywords[ident]; ok {
		return tkn
	}
	return IDENT
}

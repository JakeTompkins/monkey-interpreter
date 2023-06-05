package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER"
	INTEGER    = "INTEGER"

	// Operators
	ASSIGN    = "="
	EQUAL     = "=="
	NOT_EQUAL = "!="
	PLUS      = "+"
	MINUS     = "-"
	DIVIDE    = "/"
	MULTIPLY  = "*"
	BANG      = "!"

	// Delimiters
	COMMA        = ","
	SEMICOLON    = ";"
	DOUBLE_QUOTE = "\""
	SINGLE_QUOTE = "'"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LARROW = "<"
	RARROW = ">"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENTIFIER
}

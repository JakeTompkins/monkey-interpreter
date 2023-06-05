package lexer

import (
	"monkey-interpreter/token"
)

type Lexer struct {
	input        string
	ch           byte // current char under examination
	position     int  // index of char
	readPosition int  // index after position (need to peek ahead)
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace() // NOTE: Whitespace can be ignored,so we want to just skip over it to the next important character

	switch l.ch {
	case '!':
		if l.peekChar() == '=' {
			bang := l.ch
			l.readChar()
			equal := l.ch
			tok = newToken(token.NOT_EQUAL, bang, equal)
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '=':
		if l.peekChar() == '=' {
			left := l.ch
			l.readChar()
			right := l.ch
			tok = newToken(token.EQUAL, left, right)
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '"':
		tok = newToken(token.DOUBLE_QUOTE, l.ch)
	case '\'':
		tok = newToken(token.SINGLE_QUOTE, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '/':
		tok = newToken(token.DIVIDE, l.ch)
	case '*':
		tok = newToken(token.MULTIPLY, l.ch)
	case '<':
		tok = newToken(token.LARROW, l.ch)
	case '>':
		tok = newToken(token.RARROW, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INTEGER
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()

	return tok
}

// NOTE: DEVIATION: Refactored this slightly to take multiple byte literals
func newToken(tokenType token.TokenType, chs ...byte) token.Token {
	literal := ""

	for _, ch := range chs {
		literal += string(ch)
	}

	return token.Token{Type: tokenType, Literal: literal}
}

// TODO: readIdentifier and readNumber can be condensed into a single func
func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// NOTE: This function is important to determining the kinds of variable names that can be used. Currently just lower/upper case letters and _. If I want different kinds of identifiers, this is the place to add them
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// TODO: Extend this to be able to read in floats
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

package lexer

import (
	"testing"

	token "monkey-interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5+(){},;fn-/*<>!if true else false return!= == " '`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INTEGER, "5"},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.FUNCTION, "fn"},
		{token.MINUS, "-"},
		{token.DIVIDE, "/"},
		{token.MULTIPLY, "*"},
		{token.LARROW, "<"},
		{token.RARROW, ">"},
		{token.BANG, "!"},
		{token.IF, "if"},
		{token.TRUE, "true"},
		{token.ELSE, "else"},
		{token.FALSE, "false"},
		{token.RETURN, "return"},
		{token.NOT_EQUAL, "!="},
		{token.EQUAL, "=="},
		{token.DOUBLE_QUOTE, "\""},
		{token.SINGLE_QUOTE, "'"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf(
				"tests[%d] - literal wrong, expected %q, got %q",
				i,
				tt.expectedLiteral,
				tok.Literal,
			)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf(
				"tests[%d] - literal wrong, expected %q, got %q",
				i,
				tt.expectedLiteral,
				tok.Literal,
			)
		}
	}
}

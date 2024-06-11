package lexer

import (
	"fmt"
	"testing"

	"github.com/RafaelRCL/interpreter/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5
	let ten = 10

	let add = fn(x,y) {
		x + y;
	};

	int result = add(five, ten);
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "int"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := NewLexer(input)

	for i, tt := range tests {

		t.Run(fmt.Sprintf("Test #%v", i), func(t *testing.T) {
			tok := l.NextToken()
			t.Log(tok)

			if tok.Type != tt.expectedType {
				t.Fatalf("tokentype wrong. expected='%v' - got='%v'", tt.expectedType, tok.Type)
			}

			if tok.Literal != tt.expectedLiteral {
				t.Fatalf("token literal wrong. expected='%v' - got='%v'", tt.expectedLiteral, tok.Literal)
			}

		})
	}
}

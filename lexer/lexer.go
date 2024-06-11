package lexer

import (
	"fmt"

	"github.com/RafaelRCL/interpreter/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipSpaces()

	switch l.ch {
	case '=':
		tok = token.Token{Type: token.ASSIGN, Literal: string(l.ch)}
	case ';':
		tok = token.Token{Type: token.SEMICOLON, Literal: string(l.ch)}
	case '(':
		tok = token.Token{Type: token.LPAREN, Literal: string(l.ch)}
	case ')':
		tok = token.Token{Type: token.RPAREN, Literal: string(l.ch)}
	case ',':
		tok = token.Token{Type: token.COMMA, Literal: string(l.ch)}
	case '+':
		tok = token.Token{Type: token.PLUS, Literal: string(l.ch)}
	case '{':
		tok = token.Token{Type: token.LBRACE, Literal: string(l.ch)}
	case '}':
		tok = token.Token{Type: token.RBRACE, Literal: string(l.ch)}
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:

		switch {
		case isLetter(l.ch):
			word := l.readIdentifier()
			fmt.Println(word)
			return token.Token{Type: token.LookupIdent(word), Literal: word}

		case isDigit(l.ch):
			num := l.readNumber()
			return token.Token{Type: token.INT, Literal: num}

		default:
			tok = token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}

		}

	}

	l.readChar()

	return tok
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar() // l.readChar updates the position in the struct l
	}

	return l.input[position:l.position]
}

func (l *Lexer) readChar() {

	l.ch = 0

	if l.readPosition < len(l.input) {
		l.ch = l.input[l.readPosition]
	} // if l.readPosition is bigger than input len current char stays 0

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipSpaces() {
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

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}

	l.readChar()

	return l
}

package lexer

import "session1/token"

type Lexer struct {
	raw_input string
	cursor    int
}

func New(s string) *Lexer {
	return &Lexer{
		raw_input: s,
		cursor:    0,
	}
}

func (l *Lexer) NextToken() *token.Token {
	// Q: do we use accessor method instead of []?
	if l.cursor >= len(l.raw_input) {
		return &token.Token{Type: token.EOF, Literal: token.EOF}
	}
	next_char := l.raw_input[l.cursor]
	var next_token token.Token
	switch {
	case next_char == '=':
		next_token = token.Token{Type: token.ASSIGN, Literal: token.ASSIGN}
	case next_char == '+':
		next_token = token.Token{Type: token.PLUS, Literal: token.PLUS}
	case next_char == '(':
		next_token = token.Token{Type: token.LPAREN, Literal: token.LPAREN}
	case next_char == ')':
		next_token = token.Token{Type: token.RPAREN, Literal: token.RPAREN}
	case next_char == '{':
		next_token = token.Token{Type: token.LBRACE, Literal: token.LBRACE}
	case next_char == '}':
		next_token = token.Token{Type: token.RBRACE, Literal: token.RBRACE}
	case next_char == ',':
		next_token = token.Token{Type: token.COMMA, Literal: token.COMMA}
	case next_char == ';':
		next_token = token.Token{Type: token.SEMICOLON, Literal: token.SEMICOLON}
	default:
		next_token = token.Token{Type: token.ILLEGAL, Literal: token.ILLEGAL}

	}
	l.cursor += 1

	return &next_token
}

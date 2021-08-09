package lexer

import (
	"monkeyinterpreter/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
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
	var t string
	var literal string

	l.readChar()

	switch string(l.ch) {
	case "ILLEGAL":
		t = token.ILLEGAL
	case "IDENT":
		t = token.IDENT
	case "INT":
		t = token.INT
	case "=":
		t = token.ASSIGN
	case "+":
		t = token.PLUS
	case ",":
		t = token.COMMA
	case ";":
		t = token.SEMICOLON
	case "(":
		t = token.LPAREN
	case ")":
		t = token.RPAREN
	case "{":
		t = token.LBRACE
	case "}":
		t = token.RBRACE
	case "FUNCTION":
		t = token.FUNCTION
	case "LET":
		t = token.LET
	default:
		t = token.EOF
	}

	if t != token.EOF {
		literal = string(t)
	} else {
		literal = ""
	}

	return token.Token{
		Type:    token.TokenType(t),
		Literal: literal,
	}
}

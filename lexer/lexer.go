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
	l.readChar()
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
	var tok token.Token

	skipWhiteSpaceAndNewline(l)

	if isSpecialCharacter(l.ch) {
		tok = lexSpecialCharacter(l)
	} else if isNumeric(l.ch) {
		tok = lexNumber(l)
	} else {
		tok = lexIdentifier(l)
	}

	l.readChar()

	return tok
}

func skipWhiteSpaceAndNewline(l *Lexer) {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' {
		l.readChar()
	}
}

func lexNumber(l *Lexer) token.Token {
	num := string(l.ch)
	for isNumeric(l.peek()) {
		l.readChar()
		num += string(l.ch)
	}

	return token.Token{
		Type:    token.INT,
		Literal: num,
	}
}

func lexIdentifier(l *Lexer) token.Token {
	for isLetter(l.peek()) {
		l.readPosition++
	}
	literal := l.input[l.position:l.readPosition]

	return token.Token{
		Type:    token.LookupIdent(literal),
		Literal: literal,
	}
}

func lexSpecialCharacter(l *Lexer) token.Token {
	var tok token.Token
	switch l.ch {
	case ',':
		tok = newToken(token.COMMA, l.ch)
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
	case '=':
		nextChar := l.peek()
		if nextChar == '=' {
			l.readChar()
			return token.Token{
				Type:    token.EQ,
				Literal: "==",
			}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		nextChar := l.peek()
		if nextChar == '=' {
			l.readChar()
			return token.Token{
				Type:    token.NOT_EQ,
				Literal: "!=",
			}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	}
	return tok
}

func isLetter(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

func isNumeric(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isSpecialCharacter(ch byte) bool {
	return contains(token.SPECIALCHARLIST, ch)
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (l *Lexer) peek() byte {
	if l.readPosition >= len(l.input){
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func contains(s []byte, e byte) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

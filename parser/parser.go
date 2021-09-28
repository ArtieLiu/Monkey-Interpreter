package parser

import (
	"fmt"
	"monkeyinterpreter/ast"
	"monkeyinterpreter/lexer"
	"monkeyinterpreter/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l: l,
		errors: []string{},
	}
	// Read two tokens, so curToken and peekToken are both set
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseProgram()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseProgram() ast.Statement { // Todo understand return type
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement() // ??? type of return value?
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	// Let x = 42
	// ^

	stmt := &ast.LetStatement{
		Token: p.curToken, // current token is "LET"
	}

	// expect next token is an identifier
	if !p.expectPeek(token.IDENT) {
		return nil
	} else{
		p.nextToken()

		stmt.Name = &ast.Identifier{
			Token: p.curToken, //
			Value: p.curToken.Literal,
		}
	}

	if !p.peekTokenIs(token.ASSIGN) {
		return nil
	} else {
		// TODO: We're skipping the expressions until we encounter a semicolon
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		return true
	} else{
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t token.TokenType)  {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) Errors() []string {
	return p.errors
}
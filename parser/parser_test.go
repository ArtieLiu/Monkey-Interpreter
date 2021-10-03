package parser

import (
	"monkeyinterpreter/ast"
	"monkeyinterpreter/lexer"
	"testing"
)

func TestShouldParseLetStatements(t *testing.T) {
	input := `
   let x = 5;
   let y = 5;
   `
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t,p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

// should fail
func TestShouldParseLetStatementsWhichIsMissingAssignToken(t *testing.T) {
	input := `
   let x 5;
   `
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t,p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	tests := []struct {
		expectedIdentifier string
	}{
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

// should fail
func TestShouldParseLetStatementsWhichContainsOnlyOneValidStatement(t *testing.T) {
	input := `
   let x = 5;
   let y 5;
   `
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t,p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
	}
	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
		}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}
	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

package parser

import (
	"testing"
	"uakari/ast"
	"uakari/lexer"
)

func TestIdentifierExpression(t *testing.T) {
	input := `foobar;`

	lx := lexer.New(input)
	parser := New(lx)

	program := parser.Parse()
	checkParseError(t, parser)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement, got=%T", program.Statements[0])
	}

	ident, ok := stmt.Expression.(*ast.Identifier)

	if !ok {
		t.Errorf("exp not *ast.Identifier, got=%T", stmt.Expression)
	}

	if ident.Value != "foobar" {
		t.Errorf("ident.Value not %s, got=%s", "foobar", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Errorf("ident.TokenLiteral() not %s, got=%s", "foobar", ident.TokenLiteral())
	}
}

func TestIntegerLiteral(t *testing.T) {
	input := `5;`

	lx := lexer.New(input)
	parser := New(lx)

	program := parser.Parse()
	checkParseError(t, parser)

	if len(program.Statements) != 1 {
		t.Fatalf("program has not enough statements. got=%d", len(program.Statements))
	}

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

	if !ok {
		t.Fatalf("program.Statements[0] is not *ast.ExpressionStatement, got=%T", program.Statements[0])
	}

	literal, ok := stmt.Expression.(*ast.IntegerLiteral)

	if !ok {
		t.Errorf("exp not *ast.IntegerLiteral, got=%T", stmt.Expression)
	}

	if literal.Value != 5 {
		t.Errorf("literal.Value not %d, got=%d", 5, literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("literal.TokenLiteral() not %s, got=%s", "5", literal.TokenLiteral())
	}
}

func TestLetStatement(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foo = 86400;
`
	lx := lexer.New(input)
	p := New(lx)

	program := p.Parse()
	checkParseError(t, p)

	if program == nil {
		t.Fatalf("Parse() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain r statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		needIdentifier string
	}{
		{"x"},
		{"y"},
		{"foo"},
	}

	for i, test := range tests {
		stmt := program.Statements[i]

		if !testLetStatement(t, stmt, test.needIdentifier) {
			return
		}
	}
}

func TestReturnStatement(t *testing.T) {
	input := `
	return 10;
	return 5;
	return 10 11;
`

	lx := lexer.New(input)
	parser := New(lx)
	program := parser.Parse()

	checkParseError(t, parser)

	if program == nil {
		t.Fatalf("Parse() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain r statements. got=%d", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)

		if !ok {
			t.Errorf("stmt is not *ast.ReturnStatement, got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral() not 'return', got=%q", returnStmt.TokenLiteral())
		}
	}

}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("excpected let got=%s", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s', got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s', got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParseError(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))

	for _, err := range errors {
		t.Errorf("parser error: %q", err)
	}

	t.FailNow()
}

package parser

import (
	"fmt"
	"strconv"
	"uakari/ast"
)

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: p.currentToken,
		Value: p.currentToken.Literal,
	}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	liter := &ast.IntegerLiteral{Token: p.currentToken}
	value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)

	if err != nil {
		msg := fmt.Sprintf("clout not parse %s to integer", p.currentToken.Literal)
		p.pushError(msg)
		return nil
	}

	liter.Value = value

	return liter
}

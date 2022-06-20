package lexer

import (
	"uakari/token"
	"unicode"
)

func newToken(p token.Type, c rune) token.Token {
	return token.Token{
		Type:    p,
		Literal: string(c),
	}
}

func isLetter(c rune) bool {
	return unicode.IsLetter(c)
}

func isDigit(c rune) bool {
	return unicode.IsDigit(c)
}

func isDecimal(c rune) bool {
	return '.' == c
}

func isNegative(c rune) bool {
	return '-' == c
}

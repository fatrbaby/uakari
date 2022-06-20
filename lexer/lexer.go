package lexer

import (
	"uakari/token"
	"unicode"
)

type Lexer struct {
	input    []rune
	position int
	readPos  int
	ch       rune
}

func New(input string) *Lexer {
	lexer := &Lexer{
		input: []rune(input),
	}

	lexer.readNext()

	return lexer
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhiteSpace()

	switch l.ch {
	case '=':
		if l.nextIs('=') {
			ch := l.ch
			l.readNext()
			literal := string(ch) + string(l.ch)
			t = token.Token{Type: token.EQ, Literal: literal}
		} else {
			t = newToken(token.ASSIGN, l.ch)
		}
	case '!':
		if l.nextIs('=') {
			ch := l.ch
			l.readNext()
			literal := string(ch) + string(l.ch)
			t = token.Token{Type: token.NEQ, Literal: literal}
		} else {
			t = newToken(token.BANG, l.ch)
		}
	case '+':
		t = newToken(token.PLUS, l.ch)
	case '-':
		if isDigit(l.peekNext()) {
			return l.readNumber()
		}

		t = newToken(token.MINUS, l.ch)
	case '*':
		t = newToken(token.ASTERISK, l.ch)
	case '/':
		t = newToken(token.SLASH, l.ch)
	case '>':
		t = newToken(token.GT, l.ch)
	case '<':
		t = newToken(token.LT, l.ch)
	case ';':
		t = newToken(token.SEMICOLON, l.ch)
	case ',':
		t = newToken(token.COMMA, l.ch)
	case '(':
		t = newToken(token.LPAREN, l.ch)
	case ')':
		t = newToken(token.RPAREN, l.ch)
	case '{':
		t = newToken(token.LBRACE, l.ch)
	case '}':
		t = newToken(token.RBRACE, l.ch)
	case 0:
		t.Type = token.EOF
		t.Literal = ""
	default:
		if isLetter(l.ch) {
			t.Literal = string(l.readIdentifier())
			t.Type = token.LookupIdent(t.Literal)

			// call lexer.readNext in readIdentifier, so just return
			return t
		} else if isDigit(l.ch) {
			return l.readNumber()
		} else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readNext()

	return t
}

func (l *Lexer) readNext() {
	if l.readPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPos]
	}

	l.position = l.readPos
	l.readPos += 1
}

func (l *Lexer) readIdentifier() []rune {
	pos := l.position

	for isLetter(l.ch) {
		l.readNext()
	}

	return l.input[pos:l.position]

}

func (l *Lexer) readNumber() token.Token {
	var tp token.Type
	pos := l.position

	for isDigit(l.ch) {
		l.readNext()
	}

	tp = token.INT

	// read next char
	if isDecimal(l.ch) {
		// next char is digit
		if isDigit(l.peekNext()) {
			tp = token.FLOAT
			l.readNext()

			for isDigit(l.ch) {
				l.readNext()
			}
		} else {
			tp = token.ILLEGAL
		}
	}

	return token.Token{Type: tp, Literal: string(l.input[pos:l.position])}
}

func (l *Lexer) peekNext() rune {
	if l.readPos >= len(l.input) {
		return 0
	}

	return l.input[l.readPos]
}

func (l *Lexer) nextIs(c rune) bool {
	return l.peekNext() == c
}

func (l *Lexer) skipWhiteSpace() {
	for unicode.IsSpace(l.ch) {
		l.readNext()
	}
}

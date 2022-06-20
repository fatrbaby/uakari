package token

const (
	ILLEGAL = Type("ILLEGAL")
	EOF     = Type("EOF")

	IDENT = Type("IDENT")
	INT   = Type("INT")
	FLOAT = Type("FLOAT")

	ASSIGN   = Type("=")
	PLUS     = Type("+")
	MINUS    = Type("-")
	ASTERISK = Type("*")
	SLASH    = Type("/")
	BANG     = Type("!")

	EQ  = Type("==")
	NEQ = Type("!=")
	LT  = Type("<")
	GT  = Type(">")

	COMMA     = Type(",")
	SEMICOLON = Type(";")

	LPAREN = Type("(")
	RPAREN = Type(")")
	LBRACE = Type("{")
	RBRACE = Type("}")

	FUNCTION = Type("FUNCTION")
	LET      = Type("LET")
	TRUE     = Type("TRUE")
	FALSE    = Type("FALSE")
	IF       = Type("IF")
	ELSE     = Type("ELSE")
	RETURN   = Type("RETURN")
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) Type {
	if t, ok := keywords[ident]; ok {
		return t
	}

	return IDENT
}

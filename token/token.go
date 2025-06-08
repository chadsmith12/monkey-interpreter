package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IDENT = "IDENT"
	INT = "INT"

	ASSIGN = "="	
	PLUS = "+"
	
	COMMA = ","
	SEMICOLAN = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET = "LET"
)

var keywords = map[string]TokenType {
	"fn": FUNCTION,
	"let": LET,
}

// looks up the ident read and returns the correct token type.
// If it's a keyword returns the token type for that keyword.
// if not then returns just IDENT
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return  IDENT
}

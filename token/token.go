package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func New(ttype TokenType, literal string) Token {
	return Token{ttype, literal}
}

// define all the tkens
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// operators
	PLUS   = "+"
	ASSIGN = "="

	// delimeters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// identifiers
	LET      = "LET"
	FUNCTION = "FUNCTION"
)

// keyword map
var keywords = map[string]TokenType{
	"let": LET,
	"fn":  FUNCTION,
}

func LookUpIdent(ident string) TokenType {
	if ttype, ok := keywords[ident]; ok {
		return ttype
	}

	return IDENT
}

package token

type Type string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	LBRACE   = "LBRACE"   // {
	RBRACE   = "RBRACE"   // }
	LBRACKET = "LBRACKET" // [
	RBRACKET = "RBRACKET" // ]
	COLON    = "COLON"    // :
	COMMA    = "COMMA"    // ,

	NUMBER = "NUMBER"
	STRING = "STRING"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	NULL   = "NULL"
)

type Token struct {
	Type    Type
	Literal string
}

func New(token Type, literal byte) Token {
	return Token{
		Type:    token,
		Literal: string(literal),
	}
}

var keywords = map[string]Type{
	"true":  TRUE,
	"false": FALSE,
	"null":  NULL,
}

func LookupIdentifier(indentifier string) Type {
	if tokenType, ok := keywords[indentifier]; ok {
		return tokenType
	}
	return ILLEGAL
}

package token

type Type int

// Token types
const (
	ADD Type = iota
	SUBTRACT

	ASSIGN
	EQUATE
	SQUARE_BRACKET_OPEN
	SQUARE_BRACKET_CLOSE

	IDENTIFIER

	// Literals
	NUMBER
	TRUE
	FALSE

	// Keywords
	VARIABLE
	FUNCTION
	IF
	ELSE
)

var KeywordsToTokenType = map[string]Type{
	// Keywords
	"fun":  FUNCTION,
	"if":   IF,
	"else": ELSE,
	"var":  VARIABLE,

	// Boolean literals
	"true":  TRUE,
	"false": FALSE,
}

type Token struct {
	Type    Type
	Literal string
}

func New(tokenType Type, literal string) *Token {
	return &Token{
		Type:    tokenType,
		Literal: literal,
	}
}

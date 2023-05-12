package token

type Type int

// Token types
const (
	ADD Type = iota
	SUBTRACT

	ASSIGN
	EQUATE

	IDENTIFIER

	// Literals
	STRING
	NUMBER

	// Keywords
	FUNCTION
	IF
	ELSE
)

type Token struct {
	Type    Type
	Literal string
}

func New(tokenType Type, literal byte) *Token {
	return &Token{
		Type:    tokenType,
		Literal: string(literal),
	}
}

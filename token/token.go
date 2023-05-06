package token

type Type int

// Token types
const (
	ADD Type = iota
	SUBTRACT
	ASSIGN
	EQUATE
	IDENTIFIER
	NUMBER
)

type Token struct {
	Type    Type
	Literal string
}

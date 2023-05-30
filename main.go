package main

import (
	"interpreter/lexer"
	"interpreter/parser"
)

func main() {
	l := lexer.New("x = 4")
	_, _ = l.Lex()

	_ = parser.NewVariableDeclarationNode()
}

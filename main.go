package main

import (
	"interpreter/lexer"
)

func main() {
	l := lexer.New("x = 4")
	_ = l.Lex()
}

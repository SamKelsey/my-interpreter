package lexer

import "interpreter/token"

type Lexer interface {
	Lex() []token.Token
}

type lexer struct {
	input string
}

// TODO
func (l *lexer) Lex() []token.Token {
	return []token.Token{}
}

func New(input string) Lexer {
	return &lexer{input: input}
}

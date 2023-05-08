package lexer

import (
	"fmt"
	"interpreter/token"
	"strings"
)

type Lexer interface {
	Lex() []token.Token
}

type lexer struct {
	input string
}

// TODO
func (l *lexer) Lex() []token.Token {
	r := newReader(l.input)

	currentPosition := 0
	tokens := make([]token.Token, 0)

	for currentPosition < len(l.input) {
		b, err := r.ReadByte()
		if err != nil {
			fmt.Printf("%v", err)
		}

		tokens = append(tokens, token.Token{Literal: string(b)})
		currentPosition += 1
	}

	fmt.Println(tokens)
	return []token.Token{}
}

func New(input string) Lexer {
	return &lexer{input: input}
}

// ======= Reader ======== //

type reader struct {
	*strings.Reader
}

func (r *reader) PeekByte() (byte, error) {
	readByte, err := r.ReadByte()
	if err != nil {
		return byte(0), err
	}

	err = r.UnreadByte()
	if err != nil {
		return byte(0), err
	}

	return readByte, nil
}

func newReader(s string) *reader {
	return &reader{strings.NewReader(s)}
}

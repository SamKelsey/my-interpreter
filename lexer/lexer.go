package lexer

import (
	"fmt"
	"interpreter/token"
	"io"
	"strings"
)

type Lexer interface {
	Lex() ([]token.Token, error)
}

type lexer struct {
	input string
}

// TODO
func (l *lexer) Lex() ([]token.Token, error) {
	r := newReader(l.input)

	tokens := make([]token.Token, 0)
	var err error
	var b byte
	for err != io.EOF {
		b, err = r.ReadByte()

		// TODO: ignore whitespace here

		var newToken *token.Token
		switch string(b) {
		case "+":
			newToken = token.New(token.ADD, b)
		case "-":
			newToken = token.New(token.SUBTRACT, b)
		case "=":
			// Check for equate
			nextByte, err := r.PeekByte()
			if err != nil {
				return tokens, err
			}

			if string(nextByte) == "=" {
				newToken = token.New(token.EQUATE, b)
			} else {
				newToken = token.New(token.ASSIGN, b)
			}
		default:
			// Check for keyword
			// If not literal token
		}

		if newToken == nil {
			continue
		}

		tokens = append(tokens, *newToken)
	}

	fmt.Println(tokens)
	return []token.Token{}, nil
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

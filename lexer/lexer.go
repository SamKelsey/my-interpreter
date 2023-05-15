package lexer

import (
	log "github.com/sirupsen/logrus"
	"interpreter/token"
	"io"
	"regexp"
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
		err = l.skipWhitespace(r)
		if err != nil {
			if err == io.EOF {
				return tokens, nil
			}

			return tokens, err
		}

		b, err = r.ReadByte()

		var newToken *token.Token
		switch string(b) {
		case "+":
			newToken = token.New(token.ADD, string(b))
		case "-":
			newToken = token.New(token.SUBTRACT, string(b))
		case "[":
			newToken = token.New(token.SQUARE_BRACKET_OPEN, string(b))
		case "]":
			newToken = token.New(token.SQUARE_BRACKET_CLOSE, string(b))
		case "=":
			nextByte, err := r.PeekByte()
			if err != nil && err != io.EOF {
				return tokens, err
			}

			if string(nextByte) == "=" {
				newToken = token.New(token.EQUATE, string(b))
			} else {
				newToken = token.New(token.ASSIGN, string(b))
			}
		default:
			// If digit, keep looping until not a digit and return single token
			if isDigit(string(b)) {
				bytes := make([]byte, 0)

				for isDigit(string(b)) {
					bytes = append(bytes, b)

					b, err = r.ReadByte()
					if err != nil && err != io.EOF {
						return tokens, err
					}
				}

				newToken = token.New(token.NUMBER, string(bytes))
			} else if isLetter(string(b)) {
				bytes := make([]byte, 0)
				for isLetter(string(b)) || isDigit(string(b)) || string(b) == " " {
					// -> Check if it's a keyword or boolean literal
					tokenType, isKeyword := token.KeywordsToTokenType[string(bytes)]
					if isKeyword {
						newToken = token.New(tokenType, string(bytes))
					} else {
						// -> If not keyword, it must be literal. Keep looping until not a letter or digit **or whitespace**. (eg. parenthesis)

					}
				}

			}
		}

		if newToken == nil {
			continue
		}

		tokens = append(tokens, *newToken)
	}

	return tokens, nil
}

// Moves reader such that next time readByte() is called it will return the start of the non-whitespace.
func (l *lexer) skipWhitespace(r *reader) error {
	var b byte
	var err error
	for err != io.EOF && (string(b) == " " || b == 0) {
		b, err = r.ReadByte()
		if err != nil && err != io.EOF {
			return err
		}
	}

	if err == io.EOF {
		return err
	}

	// Unread byte so next time readByte()   is called it's the start of the non-whitespace
	err = r.UnreadByte()
	if err != nil {
		return err
	}

	return nil
}

func isDigit(c string) bool {
	match, err := regexp.MatchString("^[0-9]$", c)
	if err != nil {
		log.Warnf("error matching string: %v", err)
		return false
	}

	return match
}

func isLetter(c string) bool {
	match, err := regexp.MatchString("^[A-Za-z]$", c)
	if err != nil {
		log.Warnf("error matching string: %v", err)
		return false
	}

	return match
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
		return 0, err
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

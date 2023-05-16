package lexer

import (
	"github.com/stretchr/testify/assert"
	"interpreter/token"
	"testing"
)

func Test_lexer(t *testing.T) {

	testCases := []struct {
		name     string
		input    string
		expected []token.Token
	}{
		{
			name:  "lex number token",
			input: "1 23 45 678",
			expected: []token.Token{
				{token.NUMBER, "1"},
				{token.NUMBER, "23"},
				{token.NUMBER, "45"},
				{token.NUMBER, "678"},
			},
		},
		{
			name:  "lex identifier token",
			input: "false1",
			expected: []token.Token{
				{token.IDENTIFIER, "false1"},
			},
		},
		{
			name:  "lex keyword with identifier",
			input: "fun hello",
			expected: []token.Token{
				{token.FUNCTION, "fun"},
				{token.IDENTIFIER, "hello"},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {

			l := New(tc.input)
			got, err := l.Lex()
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, got)
		})
	}
}

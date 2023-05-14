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
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			l := New(tc.input)
			got, err := l.Lex()
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, got)
		})
	}
}

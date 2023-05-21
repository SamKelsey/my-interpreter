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
		{
			name:  "lex identifier with keyword in it",
			input: "funny fox",
			expected: []token.Token{
				{token.IDENTIFIER, "funny"},
				{token.IDENTIFIER, "fox"},
			},
		},
		{
			name:  "lex var assign statement",
			input: "var x = animal",
			expected: []token.Token{
				{token.VARIABLE, "var"},
				{token.IDENTIFIER, "x"},
				{token.ASSIGN, "="},
				{token.IDENTIFIER, "animal"},
			},
		},
		{
			name:  "lex if condition",
			input: "if weight == 1 [\n    weight = weight + 1\n]",
			expected: []token.Token{
				{token.IF, "if"},
				{token.IDENTIFIER, "weight"},
				{token.EQUATE, "=="},
				{token.NUMBER, "1"},
				{token.SQUARE_BRACKET_OPEN, "["},
				{token.IDENTIFIER, "weight"},
				{token.ASSIGN, "="},
				{token.IDENTIFIER, "weight"},
				{token.ADD, "+"},
				{token.NUMBER, "1"},
				{token.SQUARE_BRACKET_CLOSE, "]"},
			},
		},
		{
			name:  "lex func name",
			input: "fun PrintAnimal(animal) void [\n\n]",
			expected: []token.Token{
				{token.FUNCTION, "fun"},
				{token.IDENTIFIER, "PrintAnimal"},
				{token.PARENTHESIS_OPEN, "("},
				{token.IDENTIFIER, "animal"},
				{token.PARENTHESIS_CLOSE, ")"},
				{token.VOID, "void"},
				{token.SQUARE_BRACKET_OPEN, "["},
				{token.SQUARE_BRACKET_CLOSE, "]"},
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

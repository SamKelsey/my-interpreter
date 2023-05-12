package lexer

import "testing"

func Test_lexer(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		input := "var animal = \"snake\""

		lexer := New(input)
		_, _ = lexer.Lex()

	})
}

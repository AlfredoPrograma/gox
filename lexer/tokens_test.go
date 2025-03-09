package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokens(t *testing.T) {
	t.Run("should create new token from token kind", func(t *testing.T) {
		for kind, lexeme := range TokenKindToLexemeMap {
			expected := Token{kind, lexeme, 1}
			got := MustCreateTokenFromKind(kind, 1)

			assert.Equal(t, expected, got)
		}
	})

	t.Run("should panic on non matching lexeme for provided token kind", func(t *testing.T) {
		nonLexemeKinds := []TokenKind{String, Number, Identifier, Eof}

		for _, kind := range nonLexemeKinds {
			assert.Panics(t, func() { MustCreateTokenFromKind(kind, 1) })
		}
	})
}

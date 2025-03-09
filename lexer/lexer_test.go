package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLexer(t *testing.T) {
	t.Run("should tokenize single char lexemes", func(t *testing.T) {
		source := "(){},.-+;/*"
		lexer := New(source)
		expected := []Token{
			MustCreateTokenFromKind(LeftParen, 1),
			MustCreateTokenFromKind(RightParen, 1),
			MustCreateTokenFromKind(LeftBrace, 1),
			MustCreateTokenFromKind(RightBrace, 1),
			MustCreateTokenFromKind(Comma, 1),
			MustCreateTokenFromKind(Dot, 1),
			MustCreateTokenFromKind(Minus, 1),
			MustCreateTokenFromKind(Plus, 1),
			MustCreateTokenFromKind(Semicolon, 1),
			MustCreateTokenFromKind(Slash, 1),
			MustCreateTokenFromKind(Star, 1),
			MustCreateTokenFromKind(Eof, 1),
		}
		got, _ := lexer.Tokenize()

		assert.Equal(t, expected, got)
	})

	t.Run("should throw unexpected character error when there invalid characters at source", func(t *testing.T) {
		source := "()$.#," // "$" is an invalid character
		lexer := New(source)
		expected := []error{
			newUnexpectedCharacterError('$', 1, 2),
			newUnexpectedCharacterError('#', 1, 4),
		}
		_, got := lexer.Tokenize()

		assert.Equal(t, expected, got)
	})
}

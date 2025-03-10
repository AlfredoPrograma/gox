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

	t.Run("should tokenize pairable char lexemes", func(t *testing.T) {
		source := "!!====>>=<<="
		lexer := New(source)
		expected := []Token{
			MustCreateTokenFromKind(Bang, 1),
			MustCreateTokenFromKind(BangEqual, 1),
			MustCreateTokenFromKind(DoubleEqual, 1),
			MustCreateTokenFromKind(Equal, 1),
			MustCreateTokenFromKind(Greater, 1),
			MustCreateTokenFromKind(GreaterEqual, 1),
			MustCreateTokenFromKind(Less, 1),
			MustCreateTokenFromKind(LessEqual, 1),
			MustCreateTokenFromKind(Eof, 1),
		}
		got, _ := lexer.Tokenize()

		assert.Equal(t, expected, got)
	})

	t.Run("should skip comment", func(t *testing.T) {
		source := "()// This is a comment"
		lexer := New(source)
		expected := []Token{
			MustCreateTokenFromKind(LeftParen, 1),
			MustCreateTokenFromKind(RightParen, 1),
			MustCreateTokenFromKind(Eof, 1),
		}
		got, _ := lexer.Tokenize()

		assert.Equal(t, expected, got)
	})

	t.Run("should tokenize strings", func(t *testing.T) {
		source := "\"Hello world\nMy name is Gox\""
		lexer := New(source)
		expected := []Token{
			CreateToken(String, "Hello world\nMy name is Gox", 2),
			MustCreateTokenFromKind(Eof, 2),
		}
		got, _ := lexer.Tokenize()

		assert.Equal(t, expected, got)
	})

	t.Run("should throw unexpected character error when there invalid characters at source", func(t *testing.T) {
		source := "()$.#," // "$" and "#" are invalid characters
		lexer := New(source)
		expected := []error{
			newUnexpectedCharacterError('$', 1, 2),
			newUnexpectedCharacterError('#', 1, 4),
		}
		_, got := lexer.Tokenize()

		assert.Equal(t, expected, got)
	})

	t.Run("should throw unterminated string error", func(t *testing.T) {
		source := "\"Unterminated string"
		lexer := New(source)
		expected := []error{
			newUnterminatedStringError("Unterminated string", 1, 0),
		}
		_, got := lexer.Tokenize()

		assert.Equal(t, expected, got)
	})
}

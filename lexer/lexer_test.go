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

	t.Run("should tokenize numbers", func(t *testing.T) {
		source := "10.25 0.50 30 12"
		lexer := New(source)
		expected := []Token{
			CreateToken(Number, "10.25", 1),
			CreateToken(Number, "0.50", 1),
			CreateToken(Number, "30", 1),
			CreateToken(Number, "12", 1),
			MustCreateTokenFromKind(Eof, 1),
		}
		got, _ := lexer.Tokenize()

		assert.Equal(t, expected, got)
	})

	t.Run("should tokenize identifiers", func(t *testing.T) {
		source := "myVar MyVar my_var my_var1"
		lexer := New(source)
		expected := []Token{
			CreateToken(Identifier, "myVar", 1),
			CreateToken(Identifier, "MyVar", 1),
			CreateToken(Identifier, "my_var", 1),
			CreateToken(Identifier, "my_var1", 1),
			MustCreateTokenFromKind(Eof, 1),
		}
		got, _ := lexer.Tokenize()

		assert.Equal(t, expected, got)
	})

	t.Run("should tokenize keywords", func(t *testing.T) {
		source := "and class else false function for if null or print return super this true var while"
		lexer := New(source)
		expected := []Token{
			MustCreateTokenFromKind(And, 1),
			MustCreateTokenFromKind(Class, 1),
			MustCreateTokenFromKind(Else, 1),
			MustCreateTokenFromKind(False, 1),
			MustCreateTokenFromKind(Function, 1),
			MustCreateTokenFromKind(For, 1),
			MustCreateTokenFromKind(If, 1),
			MustCreateTokenFromKind(Null, 1),
			MustCreateTokenFromKind(Or, 1),
			MustCreateTokenFromKind(Print, 1),
			MustCreateTokenFromKind(Return, 1),
			MustCreateTokenFromKind(Super, 1),
			MustCreateTokenFromKind(This, 1),
			MustCreateTokenFromKind(True, 1),
			MustCreateTokenFromKind(Var, 1),
			MustCreateTokenFromKind(While, 1),
			MustCreateTokenFromKind(Eof, 1),
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

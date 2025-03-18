package ast

import (
	"testing"

	"github.com/alfredoprograma/gox/lexer"
)

func TestAST(t *testing.T) {
	type testCase struct {
		tokens   []lexer.Token
		expected Expr
	}

	testCases := []testCase{
		{
			tokens: []lexer.Token{
				lexer.CreateToken(lexer.Number, "10", 1),
				lexer.MustCreateTokenFromKind(lexer.Eof, 1),
			},
			expected: NewLiteral("10", lexer.Number),
		},
		{
			tokens: []lexer.Token{
				lexer.CreateToken(lexer.True, "true", 1),
				lexer.MustCreateTokenFromKind(lexer.Eof, 1),
			},
			expected: NewLiteral("true", lexer.True),
		},
		{
			tokens: []lexer.Token{
				lexer.CreateToken(lexer.False, "false", 1),
				lexer.MustCreateTokenFromKind(lexer.Eof, 1),
			},
			expected: NewLiteral("false", lexer.False),
		},
		{
			tokens: []lexer.Token{
				lexer.CreateToken(lexer.Null, "null", 1),
				lexer.MustCreateTokenFromKind(lexer.Eof, 1),
			},
			expected: NewLiteral("null", lexer.Null),
		},
		{
			tokens: []lexer.Token{
				lexer.CreateToken(lexer.String, "Hello world", 1),
				lexer.MustCreateTokenFromKind(lexer.Eof, 1),
			},
			expected: NewLiteral("Hello world", lexer.String),
		},
		{
			tokens: []lexer.Token{
				lexer.MustCreateTokenFromKind(lexer.LeftParen, 1),
				lexer.CreateToken(lexer.String, "Grouping expr", 1),
				lexer.MustCreateTokenFromKind(lexer.RightParen, 1),
				lexer.MustCreateTokenFromKind(lexer.Eof, 1),
			},
			expected: NewGroup(NewLiteral("Grouping expr", lexer.String)),
		},
		{
			tokens: []lexer.Token{
				lexer.MustCreateTokenFromKind(lexer.Minus, 1),
				lexer.CreateToken(lexer.Number, "12", 1),
				lexer.MustCreateTokenFromKind(lexer.Eof, 1),
			},
			expected: NewUnary(lexer.Minus, NewLiteral("12", lexer.Number)),
		},
		{
			tokens: []lexer.Token{
				lexer.CreateToken(lexer.Number, "5", 1),
				lexer.MustCreateTokenFromKind(lexer.Star, 1),
				lexer.CreateToken(lexer.Number, "5", 1),
				lexer.MustCreateTokenFromKind(lexer.Eof, 1),
			},
			expected: NewBinary(NewLiteral("5", lexer.Number), lexer.Star, NewLiteral("5", lexer.Number)),
		},
		{
			tokens: []lexer.Token{
				lexer.CreateToken(lexer.Number, "5", 1),
				lexer.MustCreateTokenFromKind(lexer.Slash, 1),
				lexer.CreateToken(lexer.Number, "5", 1),
				lexer.MustCreateTokenFromKind(lexer.Eof, 1),
			},
			expected: NewBinary(NewLiteral("5", lexer.Number), lexer.Slash, NewLiteral("5", lexer.Number)),
		},
		{
			tokens: []lexer.Token{
				lexer.CreateToken(lexer.Number, "10", 1),
				lexer.MustCreateTokenFromKind(lexer.Plus, 1),
				lexer.CreateToken(lexer.Number, "10", 1),
				lexer.MustCreateTokenFromKind(lexer.Eof, 1),
			},
			expected: NewBinary(NewLiteral("10", lexer.Number), lexer.Plus, NewLiteral("10", lexer.Number)),
		},
		{
			tokens: []lexer.Token{
				lexer.CreateToken(lexer.Number, "10", 1),
				lexer.MustCreateTokenFromKind(lexer.Minus, 1),
				lexer.CreateToken(lexer.Number, "10", 1),
				lexer.MustCreateTokenFromKind(lexer.Eof, 1),
			},
			expected: NewBinary(NewLiteral("10", lexer.Number), lexer.Minus, NewLiteral("10", lexer.Number)),
		},
		{
			tokens: []lexer.Token{
				lexer.CreateToken(lexer.Number, "14", 1),
				lexer.MustCreateTokenFromKind(lexer.Greater, 1),
				lexer.CreateToken(lexer.Number, "10", 1),
				lexer.MustCreateTokenFromKind(lexer.Eof, 1),
			},
			expected: NewBinary(NewLiteral("14", lexer.Number), lexer.Greater, NewLiteral("10", lexer.Number)),
		},
		{
			tokens: []lexer.Token{
				lexer.CreateToken(lexer.Number, "14", 1),
				lexer.MustCreateTokenFromKind(lexer.GreaterEqual, 1),
				lexer.CreateToken(lexer.Number, "10", 1),
				lexer.MustCreateTokenFromKind(lexer.Eof, 1),
			},
			expected: NewBinary(NewLiteral("14", lexer.Number), lexer.GreaterEqual, NewLiteral("10", lexer.Number)),
		}, {
			tokens: []lexer.Token{
				lexer.CreateToken(lexer.Number, "9", 1),
				lexer.MustCreateTokenFromKind(lexer.Less, 1),
				lexer.CreateToken(lexer.Number, "10", 1),
				lexer.MustCreateTokenFromKind(lexer.Eof, 1),
			},
			expected: NewBinary(NewLiteral("9", lexer.Number), lexer.Less, NewLiteral("10", lexer.Number)),
		}, {
			tokens: []lexer.Token{
				lexer.CreateToken(lexer.Number, "9", 1),
				lexer.MustCreateTokenFromKind(lexer.LessEqual, 1),
				lexer.CreateToken(lexer.Number, "10", 1),
				lexer.MustCreateTokenFromKind(lexer.Eof, 1),
			},
			expected: NewBinary(NewLiteral("9", lexer.Number), lexer.LessEqual, NewLiteral("10", lexer.Number)),
		},
		{
			tokens: []lexer.Token{
				lexer.CreateToken(lexer.Number, "7", 1),
				lexer.MustCreateTokenFromKind(lexer.DoubleEqual, 1),
				lexer.CreateToken(lexer.Number, "7", 1),
				lexer.MustCreateTokenFromKind(lexer.Eof, 1),
			},
			expected: NewBinary(NewLiteral("7", lexer.Number), lexer.DoubleEqual, NewLiteral("7", lexer.Number)),
		},
		{
			tokens: []lexer.Token{
				lexer.CreateToken(lexer.Number, "7", 1),
				lexer.MustCreateTokenFromKind(lexer.BangEqual, 1),
				lexer.CreateToken(lexer.Number, "7", 1),
				lexer.MustCreateTokenFromKind(lexer.Eof, 1),
			},
			expected: NewBinary(NewLiteral("7", lexer.Number), lexer.BangEqual, NewLiteral("7", lexer.Number)),
		},
	}

	for _, tc := range testCases {
		ast := New(tc.tokens)
		got := ast.expr()

		if tc.expected != got {
			t.Errorf("expected %v but got %v", tc.expected, got)
		}
	}
}

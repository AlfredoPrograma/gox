package ast

import (
	"testing"

	"github.com/alfredoprograma/gox/lexer"
)

func TestExpressionsStrings(t *testing.T) {
	type testCase struct {
		expr     Expr
		expected string
	}

	tcs := []testCase{
		{
			expr:     NewBinary(NewLiteral(10), lexer.MustCreateTokenFromKind(lexer.Plus, 1), NewLiteral(12)),
			expected: "(10 + 12)",
		},
		{
			expr:     NewUnary(lexer.MustCreateTokenFromKind(lexer.Minus, 1), NewLiteral(5)),
			expected: "(-5)",
		},
		{
			expr:     NewGroup(NewLiteral("Hello world")),
			expected: "(Hello world)",
		},
		{
			expr:     NewLiteral(20),
			expected: "20",
		},
		{
			expr: NewGroup(
				NewBinary(
					NewLiteral(10),
					lexer.MustCreateTokenFromKind(lexer.Star, 1),
					NewBinary(
						NewUnary(
							lexer.MustCreateTokenFromKind(lexer.Minus, 1),
							NewLiteral(20),
						),
						lexer.MustCreateTokenFromKind(lexer.Slash, 1),
						NewLiteral(8),
					),
				),
			),
			expected: "((10 * ((-20) / 8)))",
		},
	}

	for _, tc := range tcs {
		got := tc.expr.String()

		if tc.expected != got {
			t.Errorf("expected %s, but got %s", tc.expected, got)
		}
	}
}

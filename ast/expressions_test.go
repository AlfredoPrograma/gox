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
			expr:     NewBinary(NewLiteral(10), lexer.Plus, NewLiteral(12)),
			expected: "(10 + 12)",
		},
		{
			expr:     NewUnary(lexer.Minus, NewLiteral(5)),
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
					lexer.Star,
					NewBinary(
						NewUnary(
							lexer.Minus,
							NewLiteral(20),
						),
						lexer.Slash,
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

func TestExpressionsComputing(t *testing.T) {
	type testCase struct {
		expr     Expr
		expected any
	}

	testCases := []testCase{
		{
			expr:     NewBinary(NewLiteral(10.0), lexer.Slash, NewLiteral(5.0)),
			expected: 2.0,
		},
		{
			expr:     NewUnary(lexer.Minus, NewLiteral(10.0)),
			expected: -10.0,
		},
		{
			expr:     NewGroup(NewLiteral(20.0)),
			expected: 20.0,
		},
		{
			expr:     NewLiteral(1.0),
			expected: 1.0,
		},
	}

	for _, tc := range testCases {
		got, _ := tc.expr.Compute()

		if tc.expected != got {
			t.Errorf("expected %s, but got %s", tc.expected, got)
		}
	}
}

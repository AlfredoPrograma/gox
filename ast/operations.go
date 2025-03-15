package ast

import (
	"fmt"

	"github.com/alfredoprograma/gox/lexer"
)

// Computes the result of the unary operation corresponding to given operator and number
func computeNumberUnaryOperation(operator lexer.TokenKind, value float64) (float64, error) {
	switch operator {
	case lexer.Minus:
		return -value, nil
	default:
		return 0, createASTError(fmt.Sprintf("invalid operator %s for integer unary operation", lexer.TokenKindToLexemeMap[operator]))
	}
}

// Computes the result of the binary operation corresponding to given operator and numbers
func computeNumberBinaryOperation(left float64, operator lexer.TokenKind, right float64) (float64, error) {
	switch operator {
	case lexer.Plus:
		return left + right, nil
	case lexer.Minus:
		return left - right, nil
	case lexer.Star:
		return left * right, nil
	case lexer.Slash:
		return left / right, nil
	default:
		return 0, createASTError(fmt.Sprintf("invalid operator %s for integers binary operation", lexer.TokenKindToLexemeMap[operator]))
	}
}

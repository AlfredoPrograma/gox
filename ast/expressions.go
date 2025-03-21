package ast

import (
	"fmt"
	"strconv"

	"github.com/alfredoprograma/gox/lexer"
)

// An expression can generate a direct result from it.
type Expr interface {
	String() string // Exposes stringified version of the expression.
	Compute() (any, error)
}

// An expression composed by two nested expressions and an operator.
type Binary struct {
	left     Expr
	operator lexer.TokenKind
	right    Expr
}

func NewBinary(left Expr, operator lexer.TokenKind, right Expr) Expr {
	return Binary{left, operator, right}
}

func (b Binary) String() string {
	return fmt.Sprintf("(%s %s %s)", b.left.String(), lexer.TokenKindToLexemeMap[b.operator], b.right.String())
}

func (b Binary) Compute() (any, error) {
	left, err := b.left.Compute()

	if err != nil {
		return nil, err
	}

	right, err := b.right.Compute()

	if err != nil {
		return nil, err
	}

	switch leftValue := left.(type) {
	case float64:
		switch rightValue := right.(type) {
		case float64:
			return computeNumberBinaryOperation(leftValue, b.operator, rightValue)
		default:
			break
		}
	default:
		break
	}

	return nil, createASTError(fmt.Sprintf("unrecognized value types %t and %t for binary operation", left, right))
}

// An expression composed by an expression and an operator.
type Unary struct {
	operator lexer.TokenKind
	right    Expr
}

func NewUnary(operator lexer.TokenKind, right Expr) Expr {
	return Unary{operator, right}
}

func (u Unary) String() string {
	return fmt.Sprintf("(%s%s)", lexer.TokenKindToLexemeMap[u.operator], u.right.String())
}

func (u Unary) Compute() (any, error) {
	right, err := u.right.Compute()

	if err != nil {
		return nil, err
	}

	switch value := right.(type) {
	case float64:
		return computeNumberUnaryOperation(u.operator, value)
	default:
		return nil, createASTError(fmt.Sprintf("unrecognized value type %t for unary operation", value))
	}

}

// An expression which groups another expression.
type Group struct {
	expr Expr
}

func NewGroup(expr Expr) Expr {
	return Group{expr}
}

func (g Group) String() string {
	return fmt.Sprintf("(%s)", g.expr.String())
}

func (g Group) Compute() (any, error) {
	return g.expr.Compute()
}

// Bottom level expression which wraps a native type.
type Literal struct {
	value any
}

// Builds a new literal expression from lexeme and token kind.
// It parses the raw string into the corresponding native type indicated by token kind.
func NewLiteral(lexeme string, kind lexer.TokenKind) Expr {
	var value any
	var err error

	switch {
	case kind == lexer.Number:
		value, err = strconv.ParseFloat(lexeme, 64)
	case kind == lexer.True || kind == lexer.False:
		value, err = strconv.ParseBool(lexeme)
	case kind == lexer.Null:
		value = nil
	case kind == lexer.String:
		value = lexeme
	}

	// If this branch is reached, it means one of the following rare errors:
	//
	// - Some tokenization rule is broken and non parseable lexeme was generated by the lexer.
	// - Some tokenization rule is broken and provided kind does not match with the generatex lexeme.
	if err != nil {
		panic(err)
	}

	return Literal{value}
}

func (l Literal) String() string {
	return fmt.Sprintf("%v", l.value)
}

func (l Literal) Compute() (any, error) {
	return l.value, nil
}

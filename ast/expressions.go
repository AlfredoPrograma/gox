package ast

import (
	"fmt"

	"github.com/alfredoprograma/gox/lexer"
)

// An expression can generate a direct result from it.
type Expr interface {
	String() string // Exposes stringified version of the expression.
}

// An expression composed by two nested expressions and an operator.
type Binary struct {
	left     Expr
	operator lexer.Token
	right    Expr
}

func NewBinary(left Expr, operator lexer.Token, right Expr) Expr {
	return Binary{left, operator, right}
}

func (b Binary) String() string {
	return fmt.Sprintf("(%s %s %s)", b.left.String(), b.operator.Lexeme(), b.right.String())
}

// An expression composed by an expression and an operator.
type Unary struct {
	operator lexer.Token
	right    Expr
}

func NewUnary(operator lexer.Token, right Expr) Expr {
	return Unary{operator, right}
}

func (u Unary) String() string {
	return fmt.Sprintf("(%s%s)", u.operator.Lexeme(), u.right.String())
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

// Bottom level expression which wraps a native type.
type Literal struct {
	value any
}

func NewLiteral(value any) Expr {
	return Literal{value}
}

func (l Literal) String() string {
	return fmt.Sprintf("%v", l.value)
}

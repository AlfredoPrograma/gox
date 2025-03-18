package ast

import (
	"github.com/alfredoprograma/gox/lexer"
)

type AST struct {
	tokens  []lexer.Token
	errors  []error
	start   uint
	current uint
}

func New(tokens []lexer.Token) AST {
	return AST{
		tokens:  tokens,
		errors:  make([]error, 0),
		start:   0,
		current: 0,
	}
}

func (ast *AST) expr() Expr {
	return ast.equality()
}

// Equality expression is built from left and right operands, and also by an == or != operator.
// It parses the operands as comparison expressions.
// If there is not any operator, just parses a comparison expression.
func (ast *AST) equality() Expr {
	left := ast.comparison()

	if ast.match(lexer.DoubleEqual, lexer.BangEqual) {
		operator := ast.previous()
		right := ast.comparison()

		return NewBinary(left, operator.Kind, right)
	}

	return left
}

// Comparison expression is built from left and right operands, and also by an >, >=, < or <= operator.
// It parses the operands as term expressions.
// If there is not any operator, just parses a term expression.
func (ast *AST) comparison() Expr {
	left := ast.term()

	if ast.match(lexer.Greater, lexer.GreaterEqual, lexer.Less, lexer.LessEqual) {
		operator := ast.previous()
		right := ast.term()

		return NewBinary(left, operator.Kind, right)
	}

	return left
}

// Term expression is built from left and right operands, and also by an + or - operator.
// It parses the operands as factor expressions.
// If there is not any operator, just parses a factor expression.
func (ast *AST) term() Expr {
	left := ast.factor()

	if ast.match(lexer.Plus, lexer.Minus) {
		operator := ast.previous()
		right := ast.factor()

		return NewBinary(left, operator.Kind, right)
	}

	return left
}

// Factor expression is built from left and right operands, and also by an * or / operator.
// It parses the operands as unary expressions.
// If there is not any operator, just parses a unary expression.
func (ast *AST) factor() Expr {
	left := ast.unary()

	if ast.match(lexer.Star, lexer.Slash) {
		operator := ast.previous()
		right := ast.unary()

		return NewBinary(left, operator.Kind, right)
	}

	return left
}

// Unary expression is built from operator and its right operand.
// If there is not any operator, just parses a primary expression.
func (ast *AST) unary() Expr {
	if ast.match(lexer.Minus, lexer.Bang) {
		operator := ast.previous()
		expr := ast.mustPrimary()

		return NewUnary(operator.Kind, expr)
	}

	return ast.mustPrimary()
}

// Primary is the most simpler expression possible. It just holds a value.
// Also, it can be a group expression, which basically holds another nested expression.
func (ast *AST) mustPrimary() Expr {
	if ast.match(lexer.True, lexer.False, lexer.Null, lexer.Number, lexer.String) {
		token := ast.previous()
		return NewLiteral(token.Lexeme, token.Kind)
	}

	if ast.match(lexer.LeftParen) {
		expr := ast.mustPrimary() // Recursive call to itself
		ast.mustConsume(lexer.RightParen)

		return NewGroup(expr)
	}

	panic("uncaught primary expression")
}

// Checks if current token matches with the given target, but not advances.
func (ast *AST) check(kind lexer.TokenKind) bool {
	if ast.isEnd() {
		return false
	}

	return ast.peek().Kind == kind
}

// Checks if current token matches with given target, if matches, advance.
// Else, panics.
func (ast *AST) mustConsume(kind lexer.TokenKind) lexer.Token {
	if ast.check(kind) {
		return ast.advance()
	}

	panic("expected token to consume didn't match with provided at stream")
}

// Consumes the token, returns it and advance.
func (ast *AST) advance() lexer.Token {
	ast.current++
	return ast.previous()
}

// Returns the previous token.
func (ast *AST) previous() lexer.Token {
	return ast.tokens[ast.current-1]
}

// Peeks the current token without consuming it.
func (ast *AST) peek() lexer.Token {
	return ast.tokens[ast.current]
}

// Checks if tokens stream has ended.
func (ast *AST) isEnd() bool {
	return ast.peek().Kind == lexer.Eof
}

// Tries to match current token kind against the given list of kinds.
// If matches, then consumes it.
func (ast *AST) match(targetKinds ...lexer.TokenKind) bool {
	for _, target := range targetKinds {
		if ast.tokens[ast.current].Kind == target {
			ast.advance()
			return true
		}
	}

	return false
}

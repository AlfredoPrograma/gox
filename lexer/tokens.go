package lexer

import "fmt"

type TokenKind int

const (
	// Single character tokens
	LeftParen TokenKind = iota
	RightParen
	LeftBrace
	RightBrace
	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star

	// Pairable character tokens
	Bang
	BangEqual
	Equal
	DoubleEqual
	Greater
	GreaterEqual
	Less
	LessEqual

	// Literals
	Identifier
	String
	Number

	// Keywords
	And
	Class
	Else
	False
	Function
	For
	If
	Null
	Or
	Print
	Return
	Super
	This
	True
	Var
	While

	Eof
)

var TokenKindToLexemeMap = map[TokenKind]string{
	LeftParen:    "(",
	RightParen:   ")",
	LeftBrace:    "{",
	RightBrace:   "}",
	Comma:        ",",
	Dot:          ".",
	Minus:        "-",
	Plus:         "+",
	Semicolon:    ";",
	Slash:        "/",
	Star:         "*",
	Bang:         "!",
	BangEqual:    "!=",
	Equal:        "=",
	DoubleEqual:  "==",
	Greater:      ">",
	GreaterEqual: ">=",
	Less:         "<",
	LessEqual:    "<=",
	Eof:          "",
}

type Token struct {
	kind   TokenKind
	lexeme string
	line   uint
}

// Creates a new token from given args.
func CreateToken(kind TokenKind, lexeme string, line uint) Token {
	return Token{kind, lexeme, line}
}

// Creates a token with its corresponding fixed lexeme based on the provided TokenKind.
//
// If provided TokenKind does not match with any lexeme entry in TokenKindToLexemeMap; it panics.
func MustCreateTokenFromKind(kind TokenKind, line uint) Token {
	lexeme, ok := TokenKindToLexemeMap[kind]

	if !ok {
		panic("unexpected use of NewTokenFromMap. Provided TokenKind doesn't match with any lexeme")
	}

	return Token{kind, lexeme, line}
}

func (t Token) String() string {
	return fmt.Sprintf("Token <%v> (%v) at line %d", t.kind, t.lexeme, t.line)
}

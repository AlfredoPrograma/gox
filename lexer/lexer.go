package lexer

import (
	"strings"
	"unicode"
)

// Lexer reads from source and transform them into
// an intermediate meaningful representation for Gox.
type Lexer struct {
	source  string  // plain string source code
	tokens  []Token // generated tokens
	errors  []error // errors raised during tokenization
	current uint    // current cursor at source
	start   uint    // start point of each scan iteration for source
	line    uint    // current line at source
}

func New(source string) Lexer {
	return Lexer{
		source:  source,
		tokens:  make([]Token, 0),
		errors:  make([]error, 0),
		current: 0,
		start:   0,
		line:    1,
	}
}

// Transforms lexer's source into a slice of Tokens.
func (l *Lexer) Tokenize() ([]Token, []error) {
	for !l.isEnd() {
		l.start = l.current
		if err := l.scan(); err != nil {
			l.registerError(err)
		}
	}

	l.addToken(MustCreateTokenFromKind(Eof, l.line))
	return l.tokens, l.errors
}

func (l *Lexer) scan() error {
	ch := l.advance()

	switch {
	case ch == '\n':
		l.line++
	case unicode.IsSpace(ch):
		break
	case ch == '(':
		l.addToken(MustCreateTokenFromKind(LeftParen, l.line))
	case ch == ')':
		l.addToken(MustCreateTokenFromKind(RightParen, l.line))
	case ch == '{':
		l.addToken(MustCreateTokenFromKind(LeftBrace, l.line))
	case ch == '}':
		l.addToken(MustCreateTokenFromKind(RightBrace, l.line))
	case ch == ',':
		l.addToken(MustCreateTokenFromKind(Comma, l.line))
	case ch == '.':
		l.addToken(MustCreateTokenFromKind(Dot, l.line))
	case ch == '-':
		l.addToken(MustCreateTokenFromKind(Minus, l.line))
	case ch == '+':
		l.addToken(MustCreateTokenFromKind(Plus, l.line))
	case ch == ';':
		l.addToken(MustCreateTokenFromKind(Semicolon, l.line))
	case ch == '*':
		l.addToken(MustCreateTokenFromKind(Star, l.line))
	case ch == '!' && l.match('='):
		l.addToken(MustCreateTokenFromKind(BangEqual, l.line))
	case ch == '!':
		l.addToken(MustCreateTokenFromKind(Bang, l.line))
	case ch == '=' && l.match('='):
		l.addToken(MustCreateTokenFromKind(DoubleEqual, l.line))
	case ch == '=':
		l.addToken(MustCreateTokenFromKind(Equal, l.line))
	case ch == '>' && l.match('='):
		l.addToken(MustCreateTokenFromKind(GreaterEqual, l.line))
	case ch == '>':
		l.addToken(MustCreateTokenFromKind(Greater, l.line))
	case ch == '<' && l.match('='):
		l.addToken(MustCreateTokenFromKind(LessEqual, l.line))
	case ch == '<':
		l.addToken(MustCreateTokenFromKind(Less, l.line))
	case ch == '/' && l.match('/'):
		l.skipComment()
	case ch == '"':
		if err := l.string(); err != nil {
			return err
		}
	case ch == '/':
		l.addToken(MustCreateTokenFromKind(Slash, l.line))
	default:
		return newUnexpectedCharacterError(ch, l.line, l.start)
	}

	return nil
}

func (l *Lexer) registerError(err error) {
	l.errors = append(l.errors, err)
}

// Pushes a new token to the tokens slice.
func (l *Lexer) addToken(token Token) {
	l.tokens = append(l.tokens, token)
}

// Builds string token.
func (l *Lexer) string() error {
	var buf strings.Builder

	for !l.isEnd() && l.peek() != '"' {
		if l.peek() == '\n' {
			l.line++
		}

		ch := l.advance()
		buf.WriteRune(ch)
	}

	content := buf.String()

	// Notice `l.line` points to the last line of the string.
	if l.isEnd() {
		return newUnterminatedStringError(content, l.line, l.start)
	}

	l.addToken(CreateToken(String, content, l.line))
	return nil
}

// Consumes all next characters which are within a comment.
func (l *Lexer) skipComment() {
	for !l.isEnd() && l.peek() != '\n' {
		l.advance()
	}
}

// Takes the character at current source cursor and updates to next index.
func (l *Lexer) advance() rune {
	ch := rune(l.source[l.current])
	l.current++
	return ch
}

// Takes the character at current source cursor, but NOT updates to next index.
func (l *Lexer) peek() rune {
	if l.isEnd() {
		return 0
	}

	return rune(l.source[l.current])
}

// Tries to match the current source cursor character with an arbitrary character.
//
// If matches, advance it and return true; otherwise just return false.
func (l *Lexer) match(target rune) bool {
	if l.isEnd() {
		return false
	}

	next := rune(l.source[l.current])

	if next != target {
		return false
	}

	l.current++
	return true
}

// Checks if current source cursor has reached end.
func (l *Lexer) isEnd() bool {
	return int(l.current) >= len(l.source)
}

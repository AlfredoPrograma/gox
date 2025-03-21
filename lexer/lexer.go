package lexer

import (
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
	case ch == '/':
		l.addToken(MustCreateTokenFromKind(Slash, l.line))
	case ch == '"':
		if err := l.string(); err != nil {
			return err
		}
	case unicode.IsDigit(ch):
		l.number()
	case unicode.IsLetter(ch):
		l.identifierOrKeyword()
	default:
		return newUnexpectedCharacterError(ch, l.line, l.start)
	}

	return nil
}

// Pushes a new token to the tokens slice.
func (l *Lexer) addToken(token Token) {
	l.tokens = append(l.tokens, token)
}

// Builds an identifier or keyword token.
//
// During the tokenization process, lexer is not capable to determine if some stream
// of characters which start with a letter is an identifier or a keyword. So lexer
// needs to build the entire word and only then, verify if the built word is a keyword.
// If it is a keyword; search the corresponding token kind for the lexeme and create the token.
// Else, create an identifier token.
func (l *Lexer) identifierOrKeyword() {
	for !l.isEnd() && l.isValidCharForIdentifier(l.peek()) {
		l.advance()
	}

	lexeme := l.source[l.start:l.current]
	keywordKind, ok := LexemeToTokenKindMap[lexeme]

	if ok {
		l.addToken(MustCreateTokenFromKind(keywordKind, l.line))
	} else {
		l.addToken(CreateToken(Identifier, lexeme, l.line))
	}
}

// Builds number token.
//
// Numbers don't allow leading or trailing decimal point.
func (l *Lexer) number() {
	// Consumes int part of the number
	for unicode.IsDigit(l.peek()) {
		l.advance()
	}

	if l.peek() == '.' && unicode.IsDigit(l.peekNext()) {
		l.advance() // Consumes decimal point

		// Consumes decimal part of the number
		for unicode.IsDigit(l.peek()) {
			l.advance()
		}
	}

	lexeme := l.source[l.start:l.current]

	l.addToken(CreateToken(Number, lexeme, l.line))
}

// Builds string token.
//
// Strings should start with double quote character and be close with it too.
// Also, strings allow multiline by default.
func (l *Lexer) string() error {
	for !l.isEnd() && l.peek() != '"' {
		if l.peek() == '\n' {
			l.line++
		}

		l.advance()
	}

	lexeme := l.source[l.start+1 : l.current]

	// Notice l.line points to the last line of the string.
	if l.isEnd() {
		return newUnterminatedStringError(lexeme, l.line, l.start)
	}

	// Consume closing quote
	l.advance()

	l.addToken(CreateToken(String, lexeme, l.line))
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

// Takes the character at next of the current source cursor and NOT updates its index.
func (l *Lexer) peekNext() rune {
	nextIdx := l.current + 1

	if l.isEnd() || int(nextIdx) >= len(l.source) {
		return 0
	}

	return rune(l.source[nextIdx])
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

// Valid identifiers contain a combination of alphanumeric and underscore characters.
// Also, identifiers should start with a letter character.
//
// myVar, my_var, my_1_var -> Are valid identifiers.
//
// _myVar, _my_var, 1_my_var -> Aren't valid identifiers.
func (l *Lexer) isValidCharForIdentifier(ch rune) bool {
	return unicode.IsDigit(ch) || unicode.IsLetter(ch) || ch == '_'
}

// Checks if current source cursor has reached end.
func (l *Lexer) isEnd() bool {
	return int(l.current) >= len(l.source)
}

// Pushes error into lexer's errors slice.
func (l *Lexer) registerError(err error) {
	l.errors = append(l.errors, err)
}

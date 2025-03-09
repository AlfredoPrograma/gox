package lexer

// Lexer reads from source and transform them into
// an intermediate meaningful representation for Gox.
type Lexer struct {
	source  string  // plain string source code
	tokens  []Token // generated tokens
	current uint    // current cursor at source
	start   uint    // start point of each scan iteration for source
	line    uint    // current line at source
}

func New(source string) Lexer {
	return Lexer{
		source:  source,
		tokens:  make([]Token, 0),
		current: 0,
		start:   0,
		line:    1,
	}
}

// Transforms lexer's source into a slice of Tokens.
func (l *Lexer) Tokenize() []Token {
	for !l.isEnd() {
		l.start = l.current
		l.scan()
	}

	l.addToken(MustCreateTokenFromKind(Eof, l.line))
	return l.tokens
}

func (l *Lexer) scan() {
	ch := l.advance()

	switch ch {
	case '(':
		l.addToken(MustCreateTokenFromKind(LeftParen, l.line))
	case ')':
		l.addToken(MustCreateTokenFromKind(RightParen, l.line))
	case '{':
		l.addToken(MustCreateTokenFromKind(LeftBrace, l.line))
	case '}':
		l.addToken(MustCreateTokenFromKind(RightBrace, l.line))
	case ',':
		l.addToken(MustCreateTokenFromKind(Comma, l.line))
	case '.':
		l.addToken(MustCreateTokenFromKind(Dot, l.line))
	case '-':
		l.addToken(MustCreateTokenFromKind(Minus, l.line))
	case '+':
		l.addToken(MustCreateTokenFromKind(Plus, l.line))
	case ';':
		l.addToken(MustCreateTokenFromKind(Semicolon, l.line))
	case '/':
		l.addToken(MustCreateTokenFromKind(Slash, l.line))
	case '*':
		l.addToken(MustCreateTokenFromKind(Star, l.line))
	}
}

// Pushes a new token to the tokens slice.
func (l *Lexer) addToken(token Token) {
	l.tokens = append(l.tokens, token)
}

// Takes the character at current source cursor, updates to next index and returns the character.
func (l *Lexer) advance() rune {
	ch := l.source[l.current]
	l.current++
	return rune(ch)
}

// Checks if current source cursor has reached end.
func (l *Lexer) isEnd() bool {
	return int(l.current) >= len(l.source)
}

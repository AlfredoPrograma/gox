package lexer

import "fmt"

// Exposes when, during tokenization process, source contains an unexpected and non tokenizable character.
type unexpectedCharacterError struct {
	character rune
	line      uint
	position  uint
}

func newUnexpectedCharacterError(character rune, line uint, position uint) unexpectedCharacterError {
	return unexpectedCharacterError{character, line, position}
}

func (e unexpectedCharacterError) Error() string {
	return fmt.Sprintf("[Lexer]: unexpected character (%s) at line %d and position %d", string(e.character), e.line, e.position)
}

// Exposes when, during tokenization process, sources contains an unterminated string.
//
// Ex: "Hello world <- Unterminated string because does not have closing double quote.
type unterminatedStringError struct {
	content  string
	line     uint
	position uint
}

func newUnterminatedStringError(content string, line uint, position uint) unterminatedStringError {
	return unterminatedStringError{content, line, position}
}

func (e unterminatedStringError) Error() string {
	return fmt.Sprintf("[Lexer]: unterminated string (%s) at line %d and position %d", e.content, e.line, e.position)
}

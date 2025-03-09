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

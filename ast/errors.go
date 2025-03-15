package ast

import (
	"fmt"
)

const AST_PREFIX = "[AST]"

func createASTError(msg string) error {
	return fmt.Errorf("%s: %s", AST_PREFIX, msg)
}

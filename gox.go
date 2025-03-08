package gox

import (
	"fmt"
	"os"
)

// Top level runtime for Gox language
type Gox struct {
	args []string
}

func New(args []string) Gox {
	return Gox{args}
}

// Executes the Gox runtime. If file path is provided as argument, reads source code from it
// else, executes an interactive REPL prompt.
func (g *Gox) Run() {
	if len(g.args) >= 2 {
		g.readFromFile(g.args[1])
		return
	}

	g.readFromRepl()
}

func (g *Gox) readFromFile(path string) {
	source, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	fmt.Println(source)
}

func (g *Gox) readFromRepl() {
	panic("implement read source from repl")
}

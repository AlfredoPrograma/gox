package gox

import "fmt"

type Gox struct{}

func New() Gox {
	return Gox{}
}

func (g *Gox) Run() {
	fmt.Println("Hello Gox")
}

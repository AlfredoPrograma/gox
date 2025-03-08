package main

import (
	"os"

	"github.com/alfredoprograma/gox"
)

func main() {
	runtime := gox.New(os.Args)
	runtime.Run()
}

package main

import (
	"fmt"

	"github.com/mattn/go-colorable"
)

func main() {
	_, _ = fmt.Fprintln(colorable.NewColorableStdout(), "\u001b[31mHello World / Olá Mundo\u001b[0m")
}

package main

import (
	"os"
	"topus/src/topus/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}

package main

import (
	"fmt"
	"os"

	"github.com/radu-matei/monkey/repl"
)

func main() {
	fmt.Printf("REPL\n")
	repl.Start(os.Stdin, os.Stdout)
}

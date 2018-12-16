package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/radu-matei/monkey/evaluator"
	"github.com/radu-matei/monkey/lexer"
	"github.com/radu-matei/monkey/object"
	"github.com/radu-matei/monkey/parser"
	"github.com/radu-matei/monkey/repl"
)

var (
	tokenOutputPtr *bool
)

func init() {
	tokenOutputPtr = flag.Bool("output-tokens", false, "when passed, it will print all tokens from the source code")
	flag.Parse()
}

func main() {
	// no source code input file provided, starting REPL
	if len(os.Args) == 1 {
		fmt.Printf("REPL\n")
		repl.Start(os.Stdin, os.Stdout, true)

		// a source code input file was provided, attempting to interpret
	} else if len(os.Args) == 2 {
		input, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			log.Fatalf("cannot read input file: %v", err)
		}

		env := object.NewEnvironment()
		l := lexer.New(string(input), true)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			repl.PrintParserErrors(os.Stdout, p.Errors())
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(os.Stdout, evaluated.Inspect())
			io.WriteString(os.Stdout, "\n")
		}
	}
}

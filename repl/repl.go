package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/radu-matei/monkey/object"

	"github.com/radu-matei/monkey/evaluator"
	"github.com/radu-matei/monkey/lexer"
	"github.com/radu-matei/monkey/parser"
)

const prompt = ">>"

// Start starts a REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			PrintParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

// PrintParserErrors writes errors on the provided writer
func PrintParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
